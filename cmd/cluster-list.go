/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listClustersCmd = &cobra.Command{
	Use:   "list",
	Short: "List kafka clusters",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listClusters called")
	},
}

func init() {
	clusterCmd.AddCommand(listClustersCmd)
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addClusterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
