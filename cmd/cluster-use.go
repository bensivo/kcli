/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/bensivo/kcli/internal/cluster"
)

var useClusterCmd = &cobra.Command{
	Aliases: []string{"u"},
	Use:     "use <cluster>",
	Short:   "('u') Set a cluster as the default",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cluster.UseCluster(cmd.Flags().Arg(0))
	},
}

func init() {
	clusterCmd.AddCommand(useClusterCmd)
}
