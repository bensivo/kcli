package client

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"bensivo.com/kcli/internal/args"
	"github.com/segmentio/kafka-go"
)

func Produce(cfg args.ProducerArgs) {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(cfg.ClusterArgs.Timeout)*time.Second)
	bootstrapServer := cfg.ClusterArgs.BootstrapServer
	conn, err := kafka.DialLeader(ctx, "tcp", bootstrapServer, cfg.Topic, cfg.Partition) // TODO: Read from all partitions
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
