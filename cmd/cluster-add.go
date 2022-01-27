/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/bensivo/kcli/internal/cluster"
)

var addClusterCmd = &cobra.Command{
	Use:   "add <name> ",
	Short: "Add a new kafka cluster",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, arguments []string) {
		name := cmd.Flags().Arg(0)
		cluster.AddCluster(name, cluster.ClusterArgs{
			BootstrapServer: bootstrapServer,
			Timeout:         int64(connectionTimeout),
			SaslMechanism:   saslMechanism,
			SaslUsername:    saslUsername,
			SaslPassword:    saslPassword,
		})
	},
}

func init() {
	clusterCmd.AddCommand(addClusterCmd)
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addClusterCmd.Flags().StringVarP(&bootstrapServer, "bootstrap-server", "b", "localhost:9092", "Bootstrap Server")
	addClusterCmd.Flags().StringVarP(&saslMechanism, "sasl-mechanism", "m", "", "Sasl Mechanism")
	addClusterCmd.Flags().StringVarP(&saslUsername, "sasl-username", "u", "", "Sasl Username")
	addClusterCmd.Flags().StringVarP(&saslPassword, "sasl-password", "p", "", "Sasl Password")
}
