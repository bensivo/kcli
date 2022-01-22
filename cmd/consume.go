package cmd

import (
	"fmt"

	"bensivo.com/kcli/internal/args"
	"bensivo.com/kcli/internal/client"
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
		fmt.Printf("Consuming from %s:%d starting at offset %d\n", topic, partition, offset)
		client.Consume(args.ConsumerArgs{
			Topic:     topic,
			Partition: partition,
			Offset:    offset,
			Exit:      exit,
			ClusterArgs: args.ClusterArgs{
				BootstrapServer: bootstrapServer,
				Timeout:         timeoutSec,
			},
		})
	},
}
