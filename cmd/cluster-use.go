/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bensivo.com/kcli/internal/cluster"
	"github.com/spf13/cobra"
)

var useClusterCmd = &cobra.Command{
	Use:   "use <cluster>",
	Short: "Set a cluster as the default",
	Run: func(cmd *cobra.Command, args []string) {
		cluster.UseCluster(cmd.Flags().Arg(0))
	},
}

func init() {
	clusterCmd.AddCommand(useClusterCmd)
}
