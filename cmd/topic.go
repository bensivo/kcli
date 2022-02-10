/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

var topicCmd = &cobra.Command{
	Aliases: []string{"t"},
	Use:     "topic",
	Short:   "('t') Manage kafka topics",
}

func init() {
	rootCmd.AddCommand(topicCmd)
}
