/*
	Package cmd is the CLI handler
*/
package cmd

import (
	goflag "flag"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd is the root command
var RootCmd = &cobra.Command{ // nolint:gochecknoglobals
	Use:   "logtester",
	Short: "Sample Go files to test loggers",
	Long:  `Sample Go files to test loggers`,
}

// Execute is the main function
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Printf("Runtime error: %s\n", err)
		os.Exit(1)
	}
}

func getEnvReplacer() *strings.Replacer {
	return strings.NewReplacer("-", "_", ".", "_")
}

func init() { // nolint:gochecknoinits
	cobra.OnInitialize(initConfig)

	cobra.OnInitialize()

	goflag.CommandLine.Usage = func() {
		RootCmd.Usage() // nolint:gosec,errcheck
	}
	goflag.Parse()
}

func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvKeyReplacer(getEnvReplacer())
}
