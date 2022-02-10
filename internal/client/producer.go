package client

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/segmentio/kafka-go"
	"gitlab.com/bensivo/kcli/internal/cluster"
)

type ProducerArgs struct {
	Topic       string
	Partition   int
	ClusterArgs cluster.ClusterArgs
	Filepath    string
}

func Produce(cfg ProducerArgs) {
	conn := DialLeader(cfg.ClusterArgs, cfg.Topic, cfg.Partition)
	defer conn.Close()

	var bytes []byte
	var err error

	if cfg.Filepath == "" {
		bytes, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Failed to read input", err)
			os.Exit(1)
		}
	} else {
		bytes, err = ioutil.ReadFile(cfg.Filepath)
		if err != nil {
			fmt.Println("Failed to read file", err)
			os.Exit(1)
		}
	}

	_, err = conn.WriteMessages(
		kafka.Message{Value: bytes},
	)
	if err != nil {
		fmt.Println("Failed to write messages", err)
		os.Exit(1)
	}
}
