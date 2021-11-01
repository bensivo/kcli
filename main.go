package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"time"

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
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(cfg.ClusterConfig.Timeout)*time.Second)
	bootstrapServer := cfg.ClusterConfig.BootstrapServer
	conn, err := kafka.DialLeader(ctx, "tcp", bootstrapServer, cfg.Topic, cfg.Partition) // TODO: Read from all partitions
	if err != nil {
		fmt.Println("Failed to dial leader", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("connected")

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

		res, err := conn.WriteMessages(
			kafka.Message{Value: []byte(message)},
		)
		if err != nil {
			fmt.Println("Failed to write messages", err)
			os.Exit(1)
		}

		fmt.Println(res)
	}
}

func consume(cfg config.ConsumerConfig) {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(cfg.ClusterConfig.Timeout)*time.Second)
	bootstrapServer := cfg.ClusterConfig.BootstrapServer
	conn, err := kafka.DialLeader(ctx, "tcp", bootstrapServer, cfg.Topic, cfg.Partition) // TODO: Write to round-robin partitions
	if err != nil {
		fmt.Println("Failed to dial leader", err)
		os.Exit(1)
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

	last, err := conn.ReadLastOffset()
	if err != nil {
		fmt.Println("Failed to read offset", err)
	}
	if last == 0 {
		fmt.Println("End of topic")
	}

	for {
		msg, err := conn.ReadMessage(10e6)
		if err != nil {
			fmt.Println("Failed to read message", err)
			break
		}

		fmt.Printf("%d: %s\n", msg.Offset, string(msg.Value)) // TODO: Don't actually print offsets

		last, err = conn.ReadLastOffset()
		if err != nil {
			fmt.Println("Failed to read offset", err)
		}

		if msg.Offset == last-1 {
			fmt.Println("End of topic") // TODO: Option for exiting once you reach the end of the topic
		}
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
