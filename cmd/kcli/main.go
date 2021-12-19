package main

import (
	"fmt"
	"os"

	"bensivo.com/kcli/internal/args"
	"bensivo.com/kcli/internal/client"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No subcommand specified")
		fmt.Printf("Options are:\n  produce\n  consume\n  topic\n")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "produce":
		client.Produce(args.GetProducerArgs())
	case "consume":
		client.Consume(args.GetConsumerArgs())
	case "topic":
		subcommand := os.Args[2]
		switch subcommand {
		case "list":
			client.ListTopics(args.GetTopicListArgs())
		case "create":
			client.CreateTopic(args.GetCreateTopicArgs())
		}
	default:
		fmt.Printf("Command not recognized: %s\n", cmd)
		fmt.Printf("Options are:\n  produce\n  consume\n  topic\n")
	}
}
