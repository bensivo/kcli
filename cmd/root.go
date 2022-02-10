package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Variables for all the options that can be specified in all subcommands
var partition int = 0
var partitions int = 1
var replicationFactor int = 0
var offset int = 0
var exit bool = false
var bootstrapServer string = ""
var connectionTimeout int = 10
var saslMechanism string = ""
var saslUsername string = ""
var saslPassword string = ""
var clusterName string = ""
var sslEnabled bool = false
var sslSkipVerification bool = false
var sslCaCertificatePath string = ""

var rootCmd = &cobra.Command{
	Use:   "kcli",
	Short: "A simple command line client for kafka",
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
}
