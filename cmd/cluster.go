/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

var clusterCmd = &cobra.Command{
	Aliases: []string{"cl"},
	Use:     "cluster",
	Short:   "('cl') Manage multiple kafka clusters",
}

func init() {
	rootCmd.AddCommand(clusterCmd)
}
