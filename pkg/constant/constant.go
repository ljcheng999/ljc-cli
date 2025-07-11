package constant

var (
	VAR_LOG_FORMAT_TEXT bool
	VAR_LOG_FORMAT_JSON bool
	VAR_LOG_VERBOSE     bool
	// LogFormatText

	VAR_GIT_VERSION_CONTROL_PLATFORM_NAME string
	VAR_CLOUD_PROVIDER                    string
	VAR_CLOUD_REGION                      string
	VAR_AWS_CLOUD_ROLE_ARN                string
	VAR_KUBECOFNIG_LOCATION               string
	VAR_APP_DEPLOY_CLUSTER                string
	VAR_APP_DEPLOY_NAMESPACE              string
	VAR_HELM_FILE_VALUES_LOCATION         string
	VAR_HELM_RELEASE_NAME                 string
	VAR_HELM_CHART_NAME                   string
	VAR_HELM_CHART_VERSION                string
	VAR_HELM_CHART_FULL_REGISTRY_URL      string
	VAR_HELM_CHART_USERNAME               string
	VAR_HELM_CHART_PASSWORD               string

	// Not sure if I need the following for now
	VAR_GITLAB_REPO_PROJECT_OVERWRITE string
	VAR_APP_DEPLOYMENT_ENVIRONMENT    string
)

const (
	// Base
	LJC_GO_CLI_VERSION             = "1.0.0"
	GITLAB_PLATFORM_NAME           = "gitlab"
	GITHUB_PLATFORM_NAME           = "github"
	DEFAULT_HELM_CACHE_FOLDER_PATH = "/tmp/helm_cache_folder"

	// Git
	DEFAULT_VALUE_GIT_VERSION_CONTROL_PLATFORM_NAME = "git"

	// Cloud
	DEFAULT_VALUE_CLOUD_PROVIDER_NAME        = "aws"
	DEFAULT_VALUE_AWS_REGION                 = "us-east-1"
	DEFAULT_VALUE_APP_DEPLOYMENT_ENVIRONMENT = "develop"
	DEFAULT_VALUE_AWS_ASSUME_ROLE_DURATION   = 15 // minutes
	DEFAULT_VALUE_APP_DEPLOY_NAMESPACE       = "default"
	DEFAULT_VALUE_EMPTY_STRING               = ""
	DEFAULT_VALUE_TEMP_KUBE_COFNIG_LOCATION  = "/tmp/kubeconfig"

	// Default values for flags - base
	DEFAULT_FLAG_NAME_LOG_FORMAT_TEXT_NAME = "log-text"
	DEFAULT_FLAG_NAME_LOG_FORMAT_JSON_NAME = "log-json"
	DEFAULT_FLAG_NAME_LOG_VERBOSE_NAME     = "verbose"
	// Default values for flags
	DEFAULT_FLAG_NAME_GIT_VERSION_CONTROL_NAME                    = "git"
	DEFAULT_FLAG_NAME_CLOUD_PROVIDER                              = "cloud-provider"
	DEFAULT_FLAG_NAME_AWS_REGION                                  = "region"
	DEFAULT_FLAG_NAME_AWS_ROLE_ARN_NAME                           = "role-arn"
	DEFAULT_FLAG_NAME_KUBECOFNIG_LOCATION                         = "kubeconfig"
	DEFAULT_FLAG_NAME_APP_DEPLOY_CLUSTER                          = "cluster"
	DEFAULT_FLAG_NAME_APP_DEPLOY_CLUSTER_SINGLE_LETETR            = "c"
	DEFAULT_FLAG_NAME_APP_DEPLOY_CLUSTER_NAMESPACE                = "namespace"
	DEFAULT_FLAG_NAME_APP_DEPLOY_CLUSTER_NAMESPACE_SINGLE_LETETR  = "n"
	DEFAULT_FLAG_NAME_HELM_FILE_VALUES_LOCATION                   = "file"
	DEFAULT_FLAG_NAME_HELM_FILE_VALUES_LOCATION_SINGLE_LETETR     = "f"
	DEFAULT_FLAG_NAME_HELM_RELEASE_NAME                           = "release"
	DEFAULT_FLAG_NAME_HELM_RELEASE_NAME_SINGLE_LETETR             = "r"
	DEFAULT_FLAG_NAME_HELM_CHART_NAME                             = "chart"
	DEFAULT_FLAG_NAME_HELM_CHART_VERSION                          = "chart-version"
	DEFAULT_FLAG_NAME_HELM_CHART_FULL_REGISTRY_URL                = "chart-registry-url"
	DEFAULT_FLAG_NAME_HELM_CHART_USERNAME                         = "username"
	DEFAULT_FLAG_NAME_HELM_CHART_PASSWORD                         = "password"
	DEFAULT_FLAG_NAME_GITLAB_REPO_PROJECT_OVERWRITE               = "project"
	DEFAULT_FLAG_NAME_GITLAB_REPO_PROJECT_OVERWRITE_SINGLE_LETETR = "p"

	// Not sure if I need the following for now
	DEFAULT_FLAG_NAME_APP_DEPLOYMENT_ENVIRONMENT = "env"
)
