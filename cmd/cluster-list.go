/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/bensivo/kcli/internal/cluster"
)

var listClustersCmd = &cobra.Command{
	Aliases: []string{"ls"},
	Use:     "list",
	Short:   "('ls') List kafka clusters",
	Run: func(cmd *cobra.Command, args []string) {
		cluster.ListClusters()
	},
}

func init() {
	clusterCmd.AddCommand(listClustersCmd)
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addClusterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
