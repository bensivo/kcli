package cmd

import (
	"bensivo.com/kcli/internal/client"
	"bensivo.com/kcli/internal/cluster"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(consumeCmd)

	consumeCmd.Flags().IntVarP(&offset, "offset", "o", 0, "Topic offset")
	consumeCmd.Flags().IntVarP(&partition, "partition", "p", 0, "Partition")
	consumeCmd.Flags().BoolVarP(&exit, "exit", "e", false, "Exit at end of stream")
}

var consumeCmd = &cobra.Command{
	Use:   "consume <topic>",
	Short: "Consume messages fromn a topic",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, arguments []string) {
		topic := cmd.Flags().Arg(0)
		client.Consume(client.ConsumerArgs{
			Topic:       topic,
			Partition:   partition,
			Offset:      offset,
			Exit:        exit,
			ClusterArgs: cluster.GetDefaultClusterArgs(),
		})
	},
}

// kcli consume messages -e  0.13s user 0.16s system 5% cpu 5.542 total
