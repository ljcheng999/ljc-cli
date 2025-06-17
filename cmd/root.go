/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/ljcheng999/ljc-app-deploy/pkg/constant"
	"github.com/ljcheng999/ljc-app-deploy/pkg/util/logger"
	"github.com/spf13/cobra"
)

var (
	LogFormatText bool
	LogFormatJson bool
	LogVerbose    bool
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Version: constant.LCJ_DEPLOY_VERSION,
	Use:     "ljc-deploy",
	Short:   "A tool to deploy your application code to the cloud",
	Long: `ljc-deploy is a cli tool to deploy your application with helm wrapper
to the Kubernetes cluster`,
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true, // hides cmd
		// DisableDefaultCmd: true, // removes cmd
	},
	// PersistentPreRun: initCli,
	Run: runRootCmd,
}

func runRootCmd(c *cobra.Command, _ []string) {
	// display help if no subcommand provided
	c.Help()

	if LogFormatText {
		logger.SetTextLogger()
	}
	if LogFormatJson {
		logger.SetJsonLogger()
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

// func initCli(c *cobra.Command, _ []string) {
// 	v := viper.New()
// 	if debug {
// 		zerolog.SetGlobalLevel(zerolog.DebugLevel)
// 		log.Debug().Str("configFile", configFile).Msg("using config file")
// 		v.Debug()
// 	}
// 	configRepository = config.NewFileRepository(configFile, v)
// }

func init() {

	// Global flag can be used from root command to any subcommand
	RootCmd.PersistentFlags().BoolVar(&LogFormatText, "log-text", false, "Display text output format in the console. (default: false)")
	RootCmd.PersistentFlags().BoolVar(&LogFormatJson, "log-json", false, "Display json output format in the console. (default: false)")
}
