package client

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"bensivo.com/kcli/internal/cluster"
	"github.com/segmentio/kafka-go"
)

type ProducerArgs struct {
	Topic       string
	Partition   int
	ClusterArgs cluster.ClusterArgs
}

func Produce(cfg ProducerArgs) {
	conn := DialLeader(cfg.ClusterArgs, cfg.Topic, cfg.Partition)
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
