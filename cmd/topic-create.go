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
	topicCmd.AddCommand(topicCreateCmd)

	topicCreateCmd.Flags().StringVarP(&clusterName, "cluster-name", "c", "", "Cluster name")
	topicCreateCmd.Flags().IntVarP(&partitions, "partitions", "p", 1, "Number of partitions")
	topicCreateCmd.Flags().IntVarP(&replicationFactor, "replicas", "r", 1, "Replication Factor")
}

var topicCreateCmd = &cobra.Command{
	Aliases: []string{"c"},
	Use:     "create <topic>",
	Short:   "('c') Create a topic",
	Args:    cobra.ExactArgs(1),
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
