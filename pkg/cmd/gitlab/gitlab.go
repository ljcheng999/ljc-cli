package gitlab

import (
	"log/slog"

	"github.com/ljcheng999/ljc-deploy/pkg/constant"
	"github.com/ljcheng999/ljc-deploy/pkg/util"
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
	//   echo "YAMLTEMPLATE $YAMLTEMPLATE"
	//   echo "YAMLPATH $YAMLPATH"
	//   echo "CI_COMMIT_REF_SLUG $CI_COMMIT_REF_SLUG" #refslug
	//   echo "CI_PROJECT_ID $CI_PROJECT_ID"
	//   echo "CI_PROJECT_NAME $CI_PROJECT_NAME" #repo name
	//   echo "CI_ENVIRONMENT_SLUG $CI_ENVIRONMENT_SLUG"
	//   echo "CI_ENVIRONMENT_NAME $CI_ENVIRONMENT_NAME"
	//   echo "CI_PROJECT_NAMESPACE $CI_PROJECT_NAMESPACE" # split and grab first
	//   echo "USER PROVIDED RELEASE VALUE (-r): $RELEASE"

	// var_gitlab_ci_environment := strings.ToLower(os.Getenv(CI_ENVIRONMENT_NAME))
	slog.Info("-------------------GITLAB PROJECT INFO-------------------")
	slog.Info("CI_ENVIRONMENT_NAME" + util.GetEnvOrDefault(CI_ENVIRONMENT_NAME, ""))
	slog.Info("CI_COMMIT_REF_SLUG" + util.GetEnvOrDefault(CI_COMMIT_REF_SLUG, ""))
	slog.Info("CI_PROJECT_ID" + util.GetEnvOrDefault(CI_PROJECT_ID, ""))
	slog.Info("CI_PROJECT_NAME" + util.GetEnvOrDefault(CI_PROJECT_NAME, ""))
	slog.Info("CI_ENVIRONMENT_SLUG" + util.GetEnvOrDefault(CI_ENVIRONMENT_SLUG, ""))
	slog.Info("CI_ENVIRONMENT_NAME" + util.GetEnvOrDefault(CI_ENVIRONMENT_NAME, ""))
	slog.Info("CI_PROJECT_NAMESPACE" + util.GetEnvOrDefault(CI_PROJECT_NAMESPACE, ""))
}

func ReformatGitlabVarible() {

	if constant.VAR_GITLAB_REPO_PROJECT_OVERWRITE == constant.DEFAULT_VALUE_EMPTY_STRING {

	}
	//   if [[ -n "$PROJECT_OVERRIDE" ]]; then
	//       echo "Using project override $PROJECT_OVERRIDE"
	//       PROJECT=$PROJECT_OVERRIDE
	//   else
	//       PROJECT=$(echo "$CI_PROJECT_NAMESPACE" | awk '{split($0, a, "/"); print a[1]}')
	//   fi
	//   PROJECT_INCLUDE=$(echo "$PROJECT" | tr " " "-")
	//   NAMESPACE=$PROJECT_INCLUDE-$CI_PROJECT_NAME${SUFFIX}
	//   RELEASE_NAME=$PROJECT_INCLUDE-$CI_PROJECT_ID-$RELEASE
	//   ARGO_CD_MANIFEST_NAME=${PROJECT_INCLUDE}-${CI_PROJECT_NAME}

	//   if [[ -n "$NAMESPACE_OVERRIDE" ]]; then
	//       echo "Using namespace override $NAMESPACE_OVERRIDE"
	//       NAMESPACE=$NAMESPACE_OVERRIDE
	//   fi

	//   NAMESPACE=$(echo "$NAMESPACE" | awk '{print tolower($0)}' | sed 's/_/-/g')
	//   RELEASE_NAME=$(echo "$RELEASE_NAME" | awk '{print tolower($0)}' | sed 's/_/-/g' | cut -c 1-51)

	//   if [[ -n $ARGO_CD_MANIFEST_NAME_OVERWRITE ]]; then
	//     ARGO_CD_MANIFEST_NAME_OVERWRITE=$(echo "$ARGO_CD_MANIFEST_NAME_OVERWRITE" | tr " " "-")
	//     ARGO_CD_MANIFEST_NAME=$ARGO_CD_MANIFEST_NAME_OVERWRITE
	//     RELEASE_NAME=$PROJECT_INCLUDE-$CI_PROJECT_ID-$ARGO_CD_MANIFEST_NAME_OVERWRITE
	//   fi
	//   if [[ $ARGOCD == "true" ]]; then
	//     if [[ ! -n ${ARGO_CD_MANIFEST_BRANCH} ]]; then
	//       echo -e "\n[-] Error: Please setup ARGO_CD_MANIFEST_BRANCH variable. This will let system switch to your ARGO_CD_MANIFEST_BRANCH branch, and do git pull and git push"
	//       exit 1
	//     fi
	//   fi

	//   if [[ "$RELEASE_NAME" =~ -$ ]]; then
	//       echo "Found dash at the end of the release name due to truncation. Appending A"
	//       RELEASE_NAME=${RELEASE_NAME}1
	//   fi

	//   if [[ -n "$RELEASE_NAME_OVERWRITE" ]]; then
	//       RELEASE_NAME=${RELEASE_NAME_OVERWRITE}
	//   fi

	//   if [[ "$CONVERT" == true ]]; then
	//       NO_DEPLOY=true
	//   fi

}
