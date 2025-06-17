package awscloud

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func AssumeRole(awsProvider string, awsRegion string, awsAssumeRoleString string, assumeRoleDuration int64) aws.Config {
	slog.Info("Public cloud provider - " + awsProvider)
	slog.Info("AWS region - " + awsRegion)
	slog.Info("AWS role arn - " + awsAssumeRoleString)

	assumeRoleSessionName := awsProvider + "-deployment-short-session"

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(awsRegion))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	provider := stscreds.NewAssumeRoleProvider(sts.NewFromConfig(cfg), awsAssumeRoleString, func(o *stscreds.AssumeRoleOptions) {
		o.RoleSessionName = aws.ToString(&assumeRoleSessionName)
		o.Duration = time.Duration(15) * time.Minute
	})

	cfg.Credentials = aws.NewCredentialsCache(provider)

	return cfg

}
