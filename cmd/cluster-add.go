/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bensivo.com/kcli/internal/cluster"
	"github.com/spf13/cobra"
)

var addClusterCmd = &cobra.Command{
	Use:   "add <name> ",
	Short: "Add a new kafka cluster",
	Run: func(cmd *cobra.Command, arguments []string) {
		cluster.AddCluster(cluster.ClusterArgs{
			BootstrapServer: "localhost:9092",
			Timeout:         10,
		})
	},
}

func init() {
	clusterCmd.AddCommand(addClusterCmd)
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addClusterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
