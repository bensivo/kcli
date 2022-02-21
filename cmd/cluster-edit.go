/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gitlab.com/bensivo/kcli/internal/cluster"
)

var editClusterCmd = &cobra.Command{
	Aliases: []string{"e"},
	Use:     "edit <name>",
	Short:   "('e') Edit a cluster configuration",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flags().Arg(0)

		currentArgs := cluster.GetClusterArgs(name)
		cmd.Flags().Visit(func(f *pflag.Flag) {
			if f.Name == "bootstrap-server" {
				currentArgs.BootstrapServer = bootstrapServer
			}
			if f.Name == "sasl-mechanism" {
				currentArgs.SaslMechanism = saslMechanism
			}
			if f.Name == "sasl-username" {
				currentArgs.SaslUsername = saslUsername
			}
			if f.Name == "sasl-password" {
				currentArgs.SaslPassword = saslPassword
			}
			if f.Name == "ssl" {
				currentArgs.SSLEnabled = sslEnabled
			}
			if f.Name == "ssl-ca" {
				currentArgs.SSLCaCertificatePath = sslCaCertificatePath
			}
			if f.Name == "ssl-skip-verification" {
				currentArgs.SSLSkipVerification = sslSkipVerification
			}
		})
		cluster.WriteCluster(name, currentArgs)
	},
}

func init() {
	clusterCmd.AddCommand(editClusterCmd)

	editClusterCmd.Flags().StringVarP(&bootstrapServer, "bootstrap-server", "b", "", "Bootstrap Server")
	editClusterCmd.Flags().StringVarP(&saslMechanism, "sasl-mechanism", "m", "", "Sasl Mechanism")
	editClusterCmd.Flags().StringVarP(&saslUsername, "sasl-username", "u", "", "Sasl Username")
	editClusterCmd.Flags().StringVarP(&saslPassword, "sasl-password", "p", "", "Sasl Password")
	editClusterCmd.Flags().BoolVarP(&sslEnabled, "ssl", "", false, "SSL Enabled")
	editClusterCmd.Flags().StringVarP(&sslCaCertificatePath, "ssl-ca", "", "", "CA")
	editClusterCmd.Flags().BoolVarP(&sslSkipVerification, "ssl-skip-verification", "", false, "Skip Verification")
}
