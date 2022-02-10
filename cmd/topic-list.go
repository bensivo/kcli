/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/bensivo/kcli/internal/client"
	"gitlab.com/bensivo/kcli/internal/cluster"
)

func init() {
	topicCmd.AddCommand(topicListCmd)

	topicListCmd.Flags().StringVarP(&clusterName, "cluster-name", "c", "", "Cluster name")
}

var topicListCmd = &cobra.Command{
	Aliases: []string{"ls"},
	Use:     "list",
	Short:   "('ls') List topics, partition counts, and current offsets",
	Run: func(cmd *cobra.Command, arguments []string) {
		topics := client.ListTopics(cluster.GetClusterArgs(clusterName))
		for _, t := range topics {
			fmt.Printf("%s (%d)\n", t.Name, t.NumPartitions)
		}
	},
}
