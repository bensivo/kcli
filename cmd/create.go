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
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&clusterName, "cluster-name", "c", "", "Cluster name")
	createCmd.Flags().IntVarP(&partitions, "partitions", "p", 1, "Number of partitions")
	createCmd.Flags().IntVarP(&replicationFactor, "replicas", "r", 1, "Replication Factor")
}

var createCmd = &cobra.Command{
	Use:   "create <topic>",
	Short: "Create a topic",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, arguments []string) {
		topic := cmd.Flags().Arg(0)
		client.CreateTopic(client.CreateTopicArgs{
			Topic:             topic,
			Partitions:        partitions,
			ReplicationFactor: replicationFactor,
			ClusterArgs:       cluster.GetClusterArgs(clusterName),
		})
	},
}
