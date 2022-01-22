/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var useClusterCmd = &cobra.Command{
	Use:   "use <cluster>",
	Short: "Set a cluster as active",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use Cluster called")
	},
}

func init() {
	clusterCmd.AddCommand(useClusterCmd)
}
