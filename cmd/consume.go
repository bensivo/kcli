package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/bensivo/kcli/internal/client"
	"gitlab.com/bensivo/kcli/internal/cluster"
)

func init() {
	rootCmd.AddCommand(consumeCmd)

	consumeCmd.Flags().StringVarP(&clusterName, "cluster-name", "c", "", "Cluster name")
	consumeCmd.Flags().IntVarP(&offset, "offset", "o", 0, "Topic offset")
	consumeCmd.Flags().IntVarP(&partition, "partition", "p", 0, "Partition")
	consumeCmd.Flags().BoolVarP(&exit, "exit", "e", false, "Exit at end of stream")
}

var consumeCmd = &cobra.Command{
	Use:   "consume <topic>",
	Short: "Consume messages fromn a topic",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, arguments []string) {
		topic := cmd.Flags().Arg(0)
		client.Consume(client.ConsumerArgs{
			Topic:       topic,
			Partition:   partition,
			Offset:      offset,
			Exit:        exit,
			ClusterArgs: cluster.GetClusterArgs(clusterName),
		})
	},
}
