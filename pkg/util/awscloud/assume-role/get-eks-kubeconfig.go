package awscloud

import (
	"context"
	"encoding/base64"
	"errors"
	"log/slog"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/ljcheng999/ljc-app-deploy/pkg/constant"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func GetEksKubeConfig(cfg aws.Config, clusterName string, namespace string, cloudRegion string) {

	creds, _ := cfg.Credentials.Retrieve(context.Background())
	slog.Info("AccessKeyID - " + creds.AccessKeyID)

	clusterInfo, err := eks.NewFromConfig(cfg).DescribeCluster(context.Background(), &eks.DescribeClusterInput{Name: &clusterName})
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	err = CreateKubeConfigFileForRestConfig(clusterInfo, namespace, cloudRegion)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func CreateKubeConfigFileForRestConfig(clusterInfo *eks.DescribeClusterOutput, namespace string, cloudRegion string) error {

	clusterCA, err := base64.StdEncoding.DecodeString(aws.ToString(clusterInfo.Cluster.CertificateAuthority.Data))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	clusters := make(map[string]*clientcmdapi.Cluster)
	clusters[*clusterInfo.Cluster.Arn] = &clientcmdapi.Cluster{
		Server:                   *clusterInfo.Cluster.Endpoint,
		CertificateAuthorityData: clusterCA,
	}

	contexts := make(map[string]*clientcmdapi.Context)
	contexts[*clusterInfo.Cluster.Arn] = &clientcmdapi.Context{
		Cluster:   *clusterInfo.Cluster.Arn,
		AuthInfo:  *clusterInfo.Cluster.Arn,
		Namespace: namespace,
	}

	cmdArgs := []string{"--region", cloudRegion, "eks", "get-token", "--cluster-name", *clusterInfo.Cluster.Name, "--output", "json"}
	authinfos := make(map[string]*clientcmdapi.AuthInfo)
	authinfos[*clusterInfo.Cluster.Arn] = &clientcmdapi.AuthInfo{
		Exec: &clientcmdapi.ExecConfig{
			APIVersion:         "client.authentication.k8s.io/v1beta1",
			Command:            "aws",
			Args:               cmdArgs,
			Env:                nil,
			InteractiveMode:    "IfAvailable",
			ProvideClusterInfo: false,
		},
	}
	clientConfig := clientcmdapi.Config{
		Kind:           "Config",
		APIVersion:     "v1",
		Clusters:       clusters,
		Contexts:       contexts,
		CurrentContext: *clusterInfo.Cluster.Arn,
		AuthInfos:      authinfos,
	}

	if _, err := os.Stat(constant.DEFAULT_VALUE_TEMP_KUBE_COFNIG_LOCATION); errors.Is(err, os.ErrExist) {
		err = os.Remove(constant.DEFAULT_VALUE_TEMP_KUBE_COFNIG_LOCATION)
		return err
	}
	err = clientcmd.WriteToFile(clientConfig, constant.DEFAULT_VALUE_TEMP_KUBE_COFNIG_LOCATION)
	if err != nil {
		return err
	}

	return nil
}
