package gitlab

import (
	"os"
	"strings"
)

const (
	CI_ENVIRONMENT_NAME  = "CI_ENVIRONMENT_NAME"
	CI_COMMIT_REF_SLUG   = "CI_COMMIT_REF_SLUG"
	CI_PROJECT_ID        = "CI_PROJECT_ID"
	CI_PROJECT_NAME      = "CI_PROJECT_NAME"
	CI_ENVIRONMENT_SLUG  = "CI_ENVIRONMENT_SLUG"
	CI_PROJECT_NAMESPACE = "CI_PROJECT_NAMESPACE"
)

func GetGitlabProjectInformation() {

	// CI_ENVIRONMENT_NAME=${CI_ENVIRONMENT_NAME,,}
	//   echo "_________________GITLAB PROJECT INFO________________"
	//   YAMLPATH="/tmp/chart-values.yml"
	//   echo "Using $CHARTNAME on cluster $CLUSTER with id $CLUSTERID"
	//   echo "YAMLTEMPLATE $YAMLTEMPLATE"
	//   echo "YAMLPATH $YAMLPATH"
	//   echo "CI_COMMIT_REF_SLUG $CI_COMMIT_REF_SLUG" #refslug
	//   echo "CI_PROJECT_ID $CI_PROJECT_ID"
	//   echo "CI_PROJECT_NAME $CI_PROJECT_NAME" #repo name
	//   echo "CI_ENVIRONMENT_SLUG $CI_ENVIRONMENT_SLUG"
	//   echo "CI_ENVIRONMENT_NAME $CI_ENVIRONMENT_NAME"
	//   echo "CI_PROJECT_NAMESPACE $CI_PROJECT_NAMESPACE" # split and grab first
	//   echo "USER PROVIDED RELEASE VALUE (-r): $RELEASE"

	var_gitlab_ci_environment := strings.ToLower(os.Getenv(CI_ENVIRONMENT_NAME))
}
