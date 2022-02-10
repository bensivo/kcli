package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/bensivo/kcli/internal/client"
	"gitlab.com/bensivo/kcli/internal/cluster"
)

func init() {
	rootCmd.AddCommand(produceCmd)

	produceCmd.Flags().StringVarP(&clusterName, "cluster-name", "c", "", "Cluster name")
	produceCmd.Flags().IntVarP(&partition, "partition", "p", 0, "Partition")
}

var produceCmd = &cobra.Command{
	Aliases: []string{"p"},
	Use:     "produce <topic>",
	Short:   "('p') Produce messages to a topic",
	Args:    cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, arguments []string) {
		topic := cmd.Flags().Arg(0)
		filepath := cmd.Flags().Arg(1)
		client.Produce(client.ProducerArgs{
			Topic:       topic,
			Partition:   partition,
			ClusterArgs: cluster.GetClusterArgs(clusterName),
			Filepath:    filepath,
		})
	},
}
