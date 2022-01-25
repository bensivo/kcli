/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/bensivo/kcli/internal/client"
	"gitlab.com/bensivo/kcli/internal/cluster"
)

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&clusterName, "cluster-name", "c", "", "Cluster name")
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List topics, partition counts, and current offsets",
	Run: func(cmd *cobra.Command, arguments []string) {
		client.ListTopics(cluster.GetClusterArgs(clusterName))
	},
}
