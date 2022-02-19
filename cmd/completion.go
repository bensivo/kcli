/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate bash completion script",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(`
function __kcli_complete() {
	local cur prev

	cur=${COMP_WORDS[COMP_CWORD]}
	prev=${COMP_WORDS[COMP_CWORD-1]}
	pprev=${COMP_WORDS[COMP_CWORD-2]}

	case ${COMP_CWORD} in
		1) # First word 'kcli'
			COMPREPLY=($(compgen -W "cluster consume produce topic" -- ${cur}))
			;;
		2) # second word 'use prev to see what the command was'
			case ${prev} in
				cluster)
					COMPREPLY=($(compgen -W "add list remove use" -- ${cur}))
					;;
				topic)
					COMPREPLY=($(compgen -W "create list" -- ${cur}))
					;;
				consume)
					TOPICS=$(kcli topic list | cut -d ' ' -f 1 | sort)
					COMPREPLY=($(compgen -W "$TOPICS" -- ${cur}))
					;;
				produce)
					TOPICS=$(kcli topic list | cut -d ' ' -f 1 | sort)
					COMPREPLY=($(compgen -W "$TOPICS" -- ${cur}))
					;;
			esac
			;;
		3)
			case ${pprev} in 
				cluster)
					case ${prev} in
						use)
							CLUSTERS=$(kcli cluster list | grep "-" | sed 's/^ *- //g' | sed 's/ (Active)//g')
							COMPREPLY=($(compgen -W "$CLUSTERS" -- ${cur}))
							;;
						remove)
							CLUSTERS=$(kcli cluster list | grep "-" | sed 's/^ *- //g' | sed 's/ (Active)//g')
							COMPREPLY=($(compgen -W "$CLUSTERS" -- ${cur}))
							;;
					esac
					;;
			esac
			;;
		*)
			COMPREPLY=()
			;;
	esac
}

complete -F __kcli_complete kcli
		`)
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
