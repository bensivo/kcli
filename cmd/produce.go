package cmd

import (
	"bensivo.com/kcli/internal/client"
	"bensivo.com/kcli/internal/cluster"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(produceCmd)

	produceCmd.Flags().StringVarP(&clusterName, "cluster-name", "c", "", "Cluster name")
	produceCmd.Flags().IntVarP(&partition, "partition", "p", 0, "Partition")
}

var produceCmd = &cobra.Command{
	Use:   "produce <topic>",
	Short: "Produce messages to a topic",
	Run: func(cmd *cobra.Command, arguments []string) {
		topic := cmd.Flags().Arg(0)
		client.Produce(client.ProducerArgs{
			Topic:       topic,
			Partition:   partition,
			ClusterArgs: cluster.GetClusterArgs(clusterName),
		})
	},
}
