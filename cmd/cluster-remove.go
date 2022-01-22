/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var removeClusterCmd = &cobra.Command{
	Use:   "remove <name>",
	Short: "Remove a kafka cluster",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Remove Cluster called")
	},
}

func init() {
	clusterCmd.AddCommand(removeClusterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addClusterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addClusterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
