package helmc

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/ljcheng999/ljc-app-deploy/pkg/util"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/kube"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/storage/driver"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func HelmDeployLogic(
	helmChartName string,
	helmChartFullRegistryUrl string,
	helmChartVersion string,
	helmChartUsername string,
	helmChartPassword string,
	helmChartReleaseName string,
	kubeConfigPath string,
	targetNamespace string) (*release.Release, error) {

	slog.Info("helmChartName - " + helmChartName)
	slog.Info("helmChartRegistryUrl - " + helmChartFullRegistryUrl)
	slog.Info("helmChartVersion - " + helmChartVersion)
	slog.Info("kubeConfigPath - " + kubeConfigPath)
	slog.Info("helmChartUsername - " + helmChartUsername)
	slog.Info("helmChartPassword - " + helmChartPassword)
	slog.Info("helmChartReleaseName - " + helmChartReleaseName)
	slog.Info("targetNamespace - " + targetNamespace)

	actionConfig, settings, err := initActionConfig(targetNamespace, kubeConfigPath)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	var chartPathOptions action.ChartPathOptions = action.ChartPathOptions{
		RepoURL:  helmChartFullRegistryUrl,
		Version:  helmChartVersion,
		Username: helmChartUsername,
		Password: helmChartPassword,
	}

	loaderChart, err := getChart(chartPathOptions, helmChartName, settings)
	if err != nil {
		slog.Error("getChart - " + err.Error())
		os.Exit(1)
	}

	var rel *release.Release
	var requestError error

	histClient := action.NewHistory(actionConfig)
	histClient.Max = 3

	if _, err := histClient.Run(helmChartReleaseName); err == driver.ErrReleaseNotFound {
		clientInstall := action.NewInstall(actionConfig)
		clientInstall.ReleaseName = helmChartReleaseName
		clientInstall.Namespace = targetNamespace
		clientInstall.ChartPathOptions = chartPathOptions

		slog.Info("Performing the install release name " + helmChartReleaseName + " with " + helmChartName)
		fmt.Println(clientInstall)

		rel, requestError = clientInstall.Run(loaderChart, nil)
	} else {

		clientUpgrade := action.NewUpgrade(actionConfig)
		clientUpgrade.Namespace = targetNamespace
		clientUpgrade.ChartPathOptions = chartPathOptions

		slog.Info("Performing the upgrade release name " + helmChartReleaseName + " with " + helmChartName)
		fmt.Println(clientUpgrade)

		rel, requestError = clientUpgrade.Run(helmChartReleaseName, loaderChart, nil)
	}

	return rel, requestError
}

func initActionConfig(targetNamespace, kubeConfigPath string) (*action.Configuration, *cli.EnvSettings, error) {
	settings := cli.New()
	actionConfig := new(action.Configuration)

	kubeConfigFlags, err := getKubeConfig(kubeConfigPath, targetNamespace)
	if err != nil {
		slog.Error("getKubeConfig - " + err.Error())
		return nil, nil, err
	}

	err = actionConfig.Init(kubeConfigFlags, targetNamespace, util.GetEnvOrDefault("HELM_DRIVER", "secrets"), log.Printf)
	if err != nil {
		slog.Error(err.Error())
	}
	return actionConfig, settings, err
}

func getChart(chartPathOption action.ChartPathOptions, chartName string, settings *cli.EnvSettings) (*chart.Chart, error) {

	chartPath, err := chartPathOption.LocateChart(chartName, settings)
	slog.Info("chartPath - " + chartPath)
	if err != nil {
		return nil, err
	}

	chart, err := loader.Load(chartPath)
	if err != nil {
		slog.Error("loader.Load - " + err.Error())
		return nil, err
	}

	return chart, nil
}

// creating our custom RESTClientGetter or use settings.RESTClientGetter()
func getKubeConfig(kubeConfigPath string, targetNamespace string) (*genericclioptions.ConfigFlags, error) {

	kubeContext := ""
	kubeConfig := ""
	if kubeConfigPath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			slog.Error(err.Error())
			return nil, err
		}
		kubeConfig = home + "/.kube/config"

	} else {
		kubeConfig = kubeConfigPath
		// kubeconfig = "/tmp/kubeconfig"
		// kubeconfig = fmt.Sprintf("%s/.kube/config", home)
	}
	slog.Info("kubeconfig location - " + kubeConfig)

	kubeConfigFlags := kube.GetConfig(kubeConfig, kubeContext, targetNamespace)
	return kubeConfigFlags, nil
}

//////////////////////////////////////////////////////////////////////
// // Get a list in action
// listClient := action.NewList(actionConfig)
// slog.Info("New NewList")
// // Only list deployed
// listClient.Deployed = true
// releases, err := listClient.Run()
// if err != nil {
// 	slog.Error(err.Error())
// 	os.Exit(1)
// }
// slog.Info("For loop")
// for _, release := range releases {
// 	// slog.Info(release)
// 	log.Printf("here - %+v", release)
// }
//////////////////////////////////////////////////////////////////////
