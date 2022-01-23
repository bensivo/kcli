package client

import (
	"fmt"
	"os"

	"bensivo.com/kcli/internal/cluster"
	"github.com/segmentio/kafka-go"
)

func ListTopics(cfg cluster.ClusterArgs) {
	conn := Dial(cfg)
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, p := range partitions {
		fmt.Println(p.Topic, p.ID)
	}
}

type CreateTopicArgs struct {
	Topic             string
	Partitions        int
	ReplicationFactor int
	ClusterArgs       cluster.ClusterArgs
}

func CreateTopic(cfg CreateTopicArgs) {
	fmt.Printf("Creating topic %s with %d partitions, %d replicas\n", cfg.Topic, cfg.Partitions, cfg.ReplicationFactor)
	conn := Dial(cfg.ClusterArgs)
	defer conn.Close()

	topicArgss := []kafka.TopicConfig{
		{
			Topic:             cfg.Topic,
			NumPartitions:     cfg.Partitions,
			ReplicationFactor: cfg.ReplicationFactor,
		},
	}

	err := conn.CreateTopics(topicArgss...)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Topic %s created successfully\n", cfg.Topic)
}
