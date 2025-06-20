/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log/slog"
	"os"
	"strings"

	"github.com/ljcheng999/ljc-app-deploy/cmd"
	"github.com/ljcheng999/ljc-app-deploy/pkg/constant"
	awscloud "github.com/ljcheng999/ljc-app-deploy/pkg/util/awscloud/assume-role"
	helmc "github.com/ljcheng999/ljc-app-deploy/pkg/util/helmc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// gitlabCmd represents the gitlab command
var gitlabCmd = &cobra.Command{
	Use:   "gitlab",
	Short: "A subcommand where the ljc-deploy is to deploy the version control platform",
	Long:  `A subcommand where the git version control is targeting to deploy your application with Helm.`,
	Run:   runGitlabCommand,
}

func runGitlabCommand(g *cobra.Command, args []string) {

	slog.Info("gitlab subcommand got called!")

	if constant.VAR_CLOUD_PROVIDER == constant.DEFAULT_VALUE_CLOUD_PROVIDER_NAME {
		if constant.VAR_KUBECOFNIG_LOCATION == "" {
			cfg := awscloud.AssumeRole(constant.VAR_CLOUD_PROVIDER, constant.VAR_CLOUD_REGION, constant.VAR_AWS_CLOUD_ROLE_ARN, constant.DEFAULT_VALUE_AWS_ASSUME_ROLE_DURATION)
			awscloud.GetEksKubeConfig(cfg, constant.VAR_APP_DEPLOY_CLUSTER, constant.VAR_APP_DEPLOY_NAMESPACE, constant.VAR_CLOUD_REGION)
		} else {
			// Find home directory.
			home, err := os.UserHomeDir()
			cobra.CheckErr(err)
			viper.AddConfigPath(home + constant.VAR_KUBECOFNIG_LOCATION)
			viper.SetConfigFile(constant.VAR_KUBECOFNIG_LOCATION)
		}
	}

	_, err := helmc.HelmDeployLogic(
		strings.ToLower(constant.VAR_HELM_CHART_NAME),
		strings.ToLower(constant.VAR_HELM_CHART_FULL_REGISTRY_URL),
		strings.ToLower(constant.VAR_HELM_CHART_VERSION),
		strings.ToLower(constant.VAR_HELM_CHART_USERNAME),
		strings.ToLower(constant.VAR_HELM_CHART_PASSWORD),
		strings.ToLower(constant.VAR_HELM_RELEASE_NAME),
		strings.ToLower(constant.VAR_KUBECOFNIG_LOCATION),
		strings.ToLower(constant.VAR_APP_DEPLOY_NAMESPACE),
	)
	if err != nil {
		slog.Error("loader.Load - " + err.Error())
		os.Exit(1)
	}
}

func init() {
	cmd.RootCmd.AddCommand(gitlabCmd)

	gitlabCmd.PersistentFlags().StringVar(&constant.VAR_CLOUD_PROVIDER, constant.DEFAULT_FLAG_NAME_CLOUD_PROVIDER, constant.DEFAULT_VALUE_CLOUD_PROVIDER_NAME, "Public Cloud Provider")
	if constant.VAR_CLOUD_PROVIDER == constant.DEFAULT_VALUE_CLOUD_PROVIDER_NAME {
		gitlabCmd.PersistentFlags().StringVar(&constant.VAR_CLOUD_REGION, constant.DEFAULT_FLAG_NAME_AWS_REGION, constant.DEFAULT_VALUE_AWS_REGION, "AWS region")
		gitlabCmd.PersistentFlags().StringVar(&constant.VAR_AWS_CLOUD_ROLE_ARN, constant.DEFAULT_FLAG_NAME_AWS_ROLE_ARN_NAME, constant.DEFAULT_VALUE_EMPTY_STRING, "AWS role arn to be used")
		gitlabCmd.MarkFlagRequired(constant.DEFAULT_FLAG_NAME_AWS_ROLE_ARN_NAME)
	}

	gitlabCmd.PersistentFlags().StringVar(&constant.VAR_KUBECOFNIG_LOCATION, constant.DEFAULT_FLAG_NAME_KUBECOFNIG_LOCATION, constant.DEFAULT_VALUE_EMPTY_STRING, "kubeconfig file (default is $HOME/kubeconfig)")
	gitlabCmd.PersistentFlags().StringVarP(&constant.VAR_APP_DEPLOY_CLUSTER, constant.DEFAULT_FLAG_NAME_APP_DEPLOY_CLUSTER, constant.DEFAULT_FLAG_NAME_APP_DEPLOY_CLUSTER_SINGLE_LETETR, constant.DEFAULT_VALUE_EMPTY_STRING, "Deployment cluster name. (default: '')")
	gitlabCmd.PersistentFlags().StringVarP(&constant.VAR_APP_DEPLOY_NAMESPACE, constant.DEFAULT_FLAG_NAME_APP_DEPLOY_CLUSTER_NAMESPACE, constant.DEFAULT_FLAG_NAME_APP_DEPLOY_CLUSTER_NAMESPACE_SINGLE_LETETR, constant.DEFAULT_VALUE_APP_DEPLOY_NAMESPACE, "Deployment namespace.")
	gitlabCmd.PersistentFlags().StringVarP(&constant.VAR_HELM_FILE_VALUES_LOCATION, constant.DEFAULT_FLAG_NAME_HELM_FILE_VALUES_LOCATION, constant.DEFAULT_FLAG_NAME_HELM_FILE_VALUES_LOCATION_SINGLE_LETETR, constant.DEFAULT_VALUE_EMPTY_STRING, "Helm deployment values file. (default: '')")

	gitlabCmd.PersistentFlags().StringVarP(&constant.VAR_HELM_RELEASE_NAME, constant.DEFAULT_FLAG_NAME_HELM_RELEASE_NAME, constant.DEFAULT_FLAG_NAME_HELM_RELEASE_NAME_SINGLE_LETETR, constant.DEFAULT_VALUE_EMPTY_STRING, "Helm release name. (default: '')")

	gitlabCmd.PersistentFlags().StringVar(&constant.VAR_HELM_CHART_NAME, constant.DEFAULT_FLAG_NAME_HELM_CHART_NAME, constant.DEFAULT_VALUE_EMPTY_STRING, "Helm chart name. (default: '')")
	gitlabCmd.PersistentFlags().StringVar(&constant.VAR_HELM_CHART_VERSION, constant.DEFAULT_FLAG_NAME_HELM_CHART_VERSION, constant.DEFAULT_VALUE_EMPTY_STRING, "Helm chart version. (default: '')")
	gitlabCmd.PersistentFlags().StringVar(&constant.VAR_HELM_CHART_FULL_REGISTRY_URL, constant.DEFAULT_FLAG_NAME_HELM_CHART_FULL_REGISTRY_URL, constant.DEFAULT_VALUE_EMPTY_STRING, "Helm chart registry url.")
	gitlabCmd.PersistentFlags().StringVar(&constant.VAR_HELM_CHART_USERNAME, constant.DEFAULT_FLAG_NAME_HELM_CHART_USERNAME, constant.DEFAULT_VALUE_EMPTY_STRING, "Helm chart username.")
	gitlabCmd.PersistentFlags().StringVar(&constant.VAR_HELM_CHART_PASSWORD, constant.DEFAULT_FLAG_NAME_HELM_CHART_PASSWORD, constant.DEFAULT_VALUE_EMPTY_STRING, "Helm chart password.")
	gitlabCmd.PersistentFlags().StringVarP(&constant.VAR_GITLAB_REPO_PROJECT_OVERWRITE, constant.DEFAULT_FLAG_NAME_GITLAB_REPO_PROJECT_OVERWRITE, constant.DEFAULT_FLAG_NAME_GITLAB_REPO_PROJECT_OVERWRITE_SINGLE_LETETR, constant.DEFAULT_VALUE_EMPTY_STRING, "Helm deployment project name.")
	gitlabCmd.PersistentFlags().StringVar(&constant.VAR_GITLAB_APP_DEPLOY_ENVIRONMENT, constant.DEFAULT_FLAG_NAME_GITLAB_APP_DEPLOY_ENVIRONMENT, constant.DEFAULT_VALUE_APP_DEPLOYMENT_ENVIRONMENT, "Gitlab deployment environment.")

	// RootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "Display debugging output in the console. (default: false)")

}
