/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/bensivo/kcli/internal/cluster"
)

var getClusterCmd = &cobra.Command{
	Aliases: []string{"g"},
	Use:     "get <name>",
	Short:   "('g') Get a cluster configuration",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flags().Arg(0)

		currentArgs := cluster.GetClusterArgs(name)
		currentArgs.Print()
	},
}

func init() {
	clusterCmd.AddCommand(getClusterCmd)
}
