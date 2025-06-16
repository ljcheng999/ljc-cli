/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/ljcheng999/ljc-app-deploy/cmd"
	"github.com/spf13/cobra"
)

// var
var (
	CLOUD_PROVIDER string
	CLOUD_REGION   string
	// CLOUD_
)

// gitlabCmd represents the gitlab command
var gitlabCmd = &cobra.Command{
	Use:   "gitlab",
	Short: "A subcommand where the ljc-deploy is to deploy the version control platform",
	Long:  `A subcommand where the git version control is targeting to deploy your application with Helm.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitlab called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(gitlabCmd)

	gitlabCmd.Flags().StringVar(&CLOUD_PROVIDER, "cloud-provider", "aws", "Public Cloud Provider")
	gitlabCmd.Flags().StringVar(&CLOUD_REGION, "region", "us-east-1", "AWS region")

	// RootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "Display debugging output in the console. (default: false)")
}
