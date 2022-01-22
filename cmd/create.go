/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bensivo.com/kcli/internal/args"
	"bensivo.com/kcli/internal/client"
	"bensivo.com/kcli/internal/cluster"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().IntVarP(&partitions, "partitions", "p", 1, "Number of partitions")
	createCmd.Flags().IntVarP(&replicationFactor, "replicas", "r", 1, "Replication Factor")
}

var createCmd = &cobra.Command{
	Use:   "create <topic>",
	Short: "Create a topic",
	Run: func(cmd *cobra.Command, arguments []string) {
		topic := cmd.Flags().Arg(0)
		client.CreateTopic(args.CreateTopicArgs{
			Topic:             topic,
			Partitions:        partitions,
			ReplicationFactor: replicationFactor,
			ClusterArgs: cluster.ClusterArgs{
				BootstrapServer: bootstrapServer,
				Timeout:         timeoutSec,
			},
		})
	},
}
