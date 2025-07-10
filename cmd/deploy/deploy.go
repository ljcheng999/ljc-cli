/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log/slog"
	"os"
	"strings"

	"github.com/ljcheng999/ljc-cli/cmd"
	"github.com/ljcheng999/ljc-cli/pkg/constant"
	git "github.com/ljcheng999/ljc-cli/pkg/git"
	awscloud "github.com/ljcheng999/ljc-cli/pkg/util/awscloud/assume-role"
	"github.com/ljcheng999/ljc-cli/pkg/util/helmc"
	"github.com/ljcheng999/ljc-cli/pkg/util/logger"
	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "A subcommand where the ljc-cli is to deploy the application",
	Long:  `A subcommand where the public cloud provider is targeting to deploy your application with Helm.`,
	Run:   runDeployCommand,
}

func runDeployCommand(g *cobra.Command, args []string) {

	if constant.VAR_LOG_FORMAT_TEXT {
		logger.SetTextLogger()
	}
	if constant.VAR_LOG_FORMAT_JSON {
		logger.SetJsonLogger()
	}

	if constant.VAR_CLOUD_PROVIDER == constant.DEFAULT_VALUE_CLOUD_PROVIDER_NAME {
		if constant.VAR_KUBECOFNIG_LOCATION == "" {
			cfg := awscloud.AssumeRole(constant.VAR_CLOUD_PROVIDER, constant.VAR_CLOUD_REGION, constant.VAR_AWS_CLOUD_ROLE_ARN, constant.DEFAULT_VALUE_AWS_ASSUME_ROLE_DURATION)
			awscloud.GetEksKubeConfig(cfg, constant.VAR_APP_DEPLOY_CLUSTER, constant.VAR_APP_DEPLOY_NAMESPACE, constant.VAR_CLOUD_REGION)
		} else {
			// Find home directory.
			// home, err := os.UserHomeDir()
			// cobra.CheckErr(err)
			// viper.AddConfigPath(home + constant.VAR_KUBECOFNIG_LOCATION)
			// viper.SetConfigFile(constant.VAR_KUBECOFNIG_LOCATION)

		}
	}

	if constant.VAR_GIT_VERSION_CONTROL_PLATFORM_NAME == constant.GITLAB_PLATFORM_NAME {
		git.GetGitlabProjectInformation()
	}
	if constant.VAR_GIT_VERSION_CONTROL_PLATFORM_NAME == constant.GITHUB_PLATFORM_NAME {
		// ToDo
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
		slog.Error("Something wrong in HelmDeployLogic() - " + err.Error())
		os.Exit(1)
	}
}

func init() {

	// BindFlag from root command (passing flags from rootcommand to subcommand)
	// viper.BindPFlag(constant.DEFAULT_FLAG_NAME_LOG_FORMAT_TEXT_NAME, cmd.RootCmd.PersistentFlags().Lookup(constant.DEFAULT_FLAG_NAME_LOG_FORMAT_TEXT_NAME))
	// viper.BindPFlag(constant.DEFAULT_FLAG_NAME_LOG_FORMAT_JSON_NAME, cmd.RootCmd.PersistentFlags().Lookup(constant.DEFAULT_FLAG_NAME_LOG_FORMAT_TEXT_NAME))

	// Add subcommand from rootcommand
	cmd.RootCmd.AddCommand(deployCmd)

	// Git version control platform flags
	deployCmd.PersistentFlags().StringVar(&constant.VAR_GIT_VERSION_CONTROL_PLATFORM_NAME, constant.DEFAULT_VALUE_GIT_VERSION_CONTROL_PLATFORM_NAME, constant.DEFAULT_VALUE_EMPTY_STRING, "git version control platform (default '')")
	deployCmd.MarkFlagRequired(constant.VAR_GIT_VERSION_CONTROL_PLATFORM_NAME)

	// Git version control platform flags - extended (gitlab)
	// deployCmd.PersistentFlags().StringVar(&constant.VAR_APP_DEPLOYMENT_ENVIRONMENT, constant.DEFAULT_FLAG_NAME_APP_DEPLOYMENT_ENVIRONMENT, constant.DEFAULT_VALUE_APP_DEPLOYMENT_ENVIRONMENT, "Application deployment environment.")

	// Cloud Provider Platform flags
	deployCmd.PersistentFlags().StringVar(&constant.VAR_CLOUD_PROVIDER, constant.DEFAULT_FLAG_NAME_CLOUD_PROVIDER, constant.DEFAULT_VALUE_CLOUD_PROVIDER_NAME, "Public Cloud Provider")
	// if constant.VAR_CLOUD_PROVIDER == constant.DEFAULT_VALUE_CLOUD_PROVIDER_NAME {
	deployCmd.PersistentFlags().StringVar(&constant.VAR_CLOUD_REGION, constant.DEFAULT_FLAG_NAME_AWS_REGION, constant.DEFAULT_VALUE_AWS_REGION, "AWS region")
	deployCmd.PersistentFlags().StringVar(&constant.VAR_AWS_CLOUD_ROLE_ARN, constant.DEFAULT_FLAG_NAME_AWS_ROLE_ARN_NAME, constant.DEFAULT_VALUE_EMPTY_STRING, "AWS role arn to be used")
	deployCmd.MarkFlagRequired(constant.DEFAULT_FLAG_NAME_AWS_ROLE_ARN_NAME)
	// }

	// Kubernetes flags
	deployCmd.PersistentFlags().StringVar(&constant.VAR_KUBECOFNIG_LOCATION, constant.DEFAULT_FLAG_NAME_KUBECOFNIG_LOCATION, constant.DEFAULT_VALUE_EMPTY_STRING, "kubeconfig file (default is $HOME/kubeconfig)")
	deployCmd.PersistentFlags().StringVarP(&constant.VAR_APP_DEPLOY_CLUSTER, constant.DEFAULT_FLAG_NAME_APP_DEPLOY_CLUSTER, constant.DEFAULT_FLAG_NAME_APP_DEPLOY_CLUSTER_SINGLE_LETETR, constant.DEFAULT_VALUE_EMPTY_STRING, "Deployment cluster name. (default: '')")
	deployCmd.PersistentFlags().StringVarP(&constant.VAR_APP_DEPLOY_NAMESPACE, constant.DEFAULT_FLAG_NAME_APP_DEPLOY_CLUSTER_NAMESPACE, constant.DEFAULT_FLAG_NAME_APP_DEPLOY_CLUSTER_NAMESPACE_SINGLE_LETETR, constant.DEFAULT_VALUE_APP_DEPLOY_NAMESPACE, "Deployment namespace.")

	// Kubernetes flags - extended Helm
	deployCmd.PersistentFlags().StringVarP(&constant.VAR_HELM_FILE_VALUES_LOCATION, constant.DEFAULT_FLAG_NAME_HELM_FILE_VALUES_LOCATION, constant.DEFAULT_FLAG_NAME_HELM_FILE_VALUES_LOCATION_SINGLE_LETETR, constant.DEFAULT_VALUE_EMPTY_STRING, "Helm deployment values file. (default: '')")
	deployCmd.PersistentFlags().StringVarP(&constant.VAR_HELM_RELEASE_NAME, constant.DEFAULT_FLAG_NAME_HELM_RELEASE_NAME, constant.DEFAULT_FLAG_NAME_HELM_RELEASE_NAME_SINGLE_LETETR, constant.DEFAULT_VALUE_EMPTY_STRING, "Helm release name. (default: '')")
	deployCmd.PersistentFlags().StringVar(&constant.VAR_HELM_CHART_NAME, constant.DEFAULT_FLAG_NAME_HELM_CHART_NAME, constant.DEFAULT_VALUE_EMPTY_STRING, "Helm chart name. (default: '')")
	deployCmd.PersistentFlags().StringVar(&constant.VAR_HELM_CHART_VERSION, constant.DEFAULT_FLAG_NAME_HELM_CHART_VERSION, constant.DEFAULT_VALUE_EMPTY_STRING, "Helm chart version. (default: '')")
	deployCmd.PersistentFlags().StringVar(&constant.VAR_HELM_CHART_FULL_REGISTRY_URL, constant.DEFAULT_FLAG_NAME_HELM_CHART_FULL_REGISTRY_URL, constant.DEFAULT_VALUE_EMPTY_STRING, "Helm chart registry url.")
	deployCmd.PersistentFlags().StringVar(&constant.VAR_HELM_CHART_USERNAME, constant.DEFAULT_FLAG_NAME_HELM_CHART_USERNAME, constant.DEFAULT_VALUE_EMPTY_STRING, "Helm chart username.")
	deployCmd.PersistentFlags().StringVar(&constant.VAR_HELM_CHART_PASSWORD, constant.DEFAULT_FLAG_NAME_HELM_CHART_PASSWORD, constant.DEFAULT_VALUE_EMPTY_STRING, "Helm chart password.")
}
