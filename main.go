package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/segmentio/kafka-go"

	"bensivo.com/kcli/config"
)

func absInt(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func produce(cfg config.ProducerConfig) {
	bootstrapServer := cfg.ClusterConfig.BootstrapServer
	conn, err := kafka.DialLeader(context.Background(), "tcp", bootstrapServer, cfg.Topic, cfg.Partition)
	if err != nil {
		fmt.Println("Failed to dial leader", err)
		os.Exit(1)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			os.Exit(0)
		}

		if err != nil {
			fmt.Println("Failed to read input", err)
			os.Exit(1)
		}

		message := input[:len(input)-1]

		_, err = conn.WriteMessages(
			kafka.Message{Value: []byte(message)},
		)
		if err != nil {
			fmt.Println("Failed to write messages", err)
			os.Exit(1)
		}
	}
}

func consume(cfg config.ConsumerConfig) {
	bootstrapServer := cfg.ClusterConfig.BootstrapServer
	conn, err := kafka.DialLeader(context.Background(), "tcp", bootstrapServer, cfg.Topic, cfg.Partition)
	if err != nil {
		fmt.Println("Failed to dial leader", err)
	}
	defer conn.Close()

	if cfg.Offset != 0 {
		var seekPos int
		if cfg.Offset > 0 {
			seekPos = kafka.SeekStart
		} else {
			seekPos = kafka.SeekEnd
		}

		_, err := conn.Seek(int64(absInt(cfg.Offset)), seekPos)
		if err != nil {
			fmt.Println("Failed to seek offset", err)
			return
		}
	}

	for {
		msg, err := conn.ReadMessage(10e6)
		if err != nil {
			fmt.Println("Failed to read message", err)
			break
		}

		fmt.Printf("%d: %s\n", msg.Offset, string(msg.Value))
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No subcommand specified")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "produce":
		produce(config.GetProducerConfig())
	case "consume":
		consume(config.GetConsumerConfig())
	default:
		fmt.Printf("Command not recognized: %s\n", cmd)
	}
}
