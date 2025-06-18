package helmc

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/ljcheng999/ljc-app-deploy/pkg/constant"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/storage/driver"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/tools/clientcmd"
)

func HelmDeployLogic(helmChartFullRegistryUrl string, helmChartVersion string, helmChartUsername string, helmChartPassword string, kubeConfigPath string) {
	slog.Info("helmChartRegistryUrl - " + helmChartFullRegistryUrl)
	slog.Info("helmChartVersion - " + helmChartVersion)
	slog.Info("kubeConfigPath - " + kubeConfigPath)
	slog.Info("helmChartUsername - " + helmChartUsername)
	slog.Info("helmChartPassword - " + helmChartPassword)

	settings := cli.New()
	actionConfig, err := initActionConfig(settings, "cj-ns", "docker-desktop")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	chartPathOptions := &action.ChartPathOptions{
		RepoURL:  helmChartFullRegistryUrl,
		Version:  helmChartVersion,
		Username: helmChartUsername,
		Password: helmChartPassword,
	}

	chart, err := action.NewPull().LocateChart(chartPathOptions.RepoURL, settings)
	if err != nil {
		slog.Error("NewPull - " + err.Error())
		os.Exit(1)
	}
	loaderChart, err := loader.Load(chart)
	// chart, err := getChart(chartPathOptions, chartName, settings)
	if err != nil {
		slog.Error("loader.Load - " + err.Error())
		os.Exit(1)
	}

	// var rel *release.Release
	// var requestError error

	histClient := action.NewHistory(actionConfig)
	histClient.Max = 3

	if _, err := histClient.Run(constant.VAR_HELM_RELEASE_NAME); err == driver.ErrReleaseNotFound {
		clientInstall := action.NewInstall(actionConfig)
		clientInstall.ReleaseName = constant.VAR_HELM_RELEASE_NAME
		clientInstall.Namespace = constant.VAR_APP_DEPLOY_NAMESPACE
		clientInstall.ChartPathOptions = *chartPathOptions

		fmt.Println(clientInstall)

		clientInstall.Run(loaderChart, nil)
		// rel, requestError = clientInstall.Run(loaderChart, nil)
	} else {
		clientUpgrade := action.NewUpgrade(actionConfig)
		clientUpgrade.Namespace = constant.VAR_APP_DEPLOY_NAMESPACE
		clientUpgrade.ChartPathOptions = *chartPathOptions

		fmt.Println(clientUpgrade)

		clientUpgrade.Run(constant.VAR_HELM_RELEASE_NAME, loaderChart, nil)
		// rel, requestError = clientUpgrade.Run(constant.VAR_HELM_RELEASE_NAME, loaderChart, nil)
	}

	// config, err := clientcmd.LoadFromFile(kubeConfigPath)
	// if err != nil {
	// 	slog.Error(err.Error())
	// 	os.Exit(1)
	// }

	// contextName := "docker-desktop"
	// configFlags := genericclioptions.NewConfigFlags(true)
	// configFlags.KubeConfig = &kubeConfigPath
	// configFlags.Context = &contextName
	// settings := cli.New()

	// actionConfig := new(action.Configuration)

	// if err := actionConfig.Init(
	// 	//settings.RESTClientGetter(),
	// 	restClientGetter,
	// 	namespace,
	// 	helmDriver,
	// 	log.Printf); err != nil {
	// 	return nil, err
	// }

	// helmChartOptions := &action.ChartPathOptions{
	// 	RepoURL:  helmChartFullRegistryUrl,
	// 	Version:  helmChartVersion,
	// 	Username: helmChartUsername,
	// 	Password: helmChartPassword,
	// }
	// // // helmChartOptions := &action.ChartPathOptions{
	// // // 	RepoURL: "https://argoproj.github.io/argo-helm",
	// // // 	Version: "6.9.1",
	// // // }
	// // //var outputBuffer bytes.Buffer

	// getChart, err := action.NewPull().LocateChart(helmChartOptions.RepoURL, settings)
	// if err != nil {
	// 	slog.Error(err.Error())
	// 	os.Exit(1)
	// }
	// loaderChart, err := loader.Load(getChart)
	// if err != nil {
	// 	slog.Error(err.Error())
	// 	os.Exit(1)
	// }

}

func initActionConfig(settings *cli.EnvSettings, namespace, contextName string) (*action.Configuration, error) {
	restClientGetter, err := getKubeConfigWithContext(contextName)
	if err != nil {
		if err != nil {
			slog.Error(err.Error())
		}
		return nil, fmt.Errorf("failed to load kubeconfig")
	}
	actionConfig := new(action.Configuration)

	if namespace == "" {
		namespace = settings.Namespace()
	}

	HELM_DRIVER := "secrets"
	if err := actionConfig.Init(
		//settings.RESTClientGetter(),
		restClientGetter,
		namespace,
		HELM_DRIVER,
		log.Printf); err != nil {
		return nil, err
	}

	return actionConfig, nil
}

// creating our custom RESTClientGetter or use settings.RESTClientGetter()
func getKubeConfigWithContext(contextName string) (genericclioptions.RESTClientGetter, error) {
	kubeconfig := "/tmp/kubeconfig"
	// kubeconfig := os.Getenv("KUBECONFIG")
	// if kubeconfig == "" {
	// 	if home := os.UserHomeDir(); home != "" {
	// 		kubeconfig = fmt.Sprintf("%s/.kube/config", home)
	// 	}
	// }

	if constant.VAR_KUBECOFNIG_LOCATION == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		if home != "" {
			kubeconfig = "/tmp/kubeconfig"
			// kubeconfig = fmt.Sprintf("%s/.kube/config", home)
		}
	}

	config, err := clientcmd.LoadFromFile(kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("failed to load kubeconfig: %v", err)
	}

	// Ensure context exists
	if _, exists := config.Contexts[contextName]; !exists {
		return nil, fmt.Errorf("context %s not found in kubeconfig", contextName)
	}

	configFlags := genericclioptions.NewConfigFlags(true)
	configFlags.KubeConfig = &kubeconfig
	configFlags.Context = &contextName

	return configFlags, nil
}
