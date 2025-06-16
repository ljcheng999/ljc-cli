package awscloud

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func AssumeRole(awsRegion string, awsAssumeRoleString string, appDeployEnvironment string, assumeRoleDuration int64) {
	slog.Info(awsRegion)
	slog.Info(awsAssumeRoleString)
	slog.Info(appDeployEnvironment)

	assumeRoleSessionName := appDeployEnvironment + "-deployment-short-session"

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(awsRegion))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	provider := stscreds.NewAssumeRoleProvider(sts.NewFromConfig(cfg), awsAssumeRoleString, func(o *stscreds.AssumeRoleOptions) {
		o.RoleSessionName = aws.ToString(&assumeRoleSessionName)
		o.Duration = time.Duration(15) * time.Minute
	})

	cfg.Credentials = aws.NewCredentialsCache(provider)
	creds, err := cfg.Credentials.Retrieve(context.Background())
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	if !creds.HasKeys() {
		slog.Error("No credential keys returned")
		os.Exit(1)
	}

	// fmt.Printf("%+v\n", cfg)
	fmt.Printf("here - %+v\n", creds)

}
