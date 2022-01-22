/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bensivo.com/kcli/internal/client"
	"bensivo.com/kcli/internal/cluster"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List topics, partition counts, and current offsets",
	Run: func(cmd *cobra.Command, arguments []string) {
		client.ListTopics(cluster.GetDefaultClusterArgs())
	},
}
