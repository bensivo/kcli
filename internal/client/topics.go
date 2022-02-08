package client

import (
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
	"gitlab.com/bensivo/kcli/internal/cluster"
)

type TopicMetadata struct {
	Name          string
	NumPartitions int
}

func ListTopics(cfg cluster.ClusterArgs) map[string]TopicMetadata {
	conn := Dial(cfg)
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	topics := make(map[string]TopicMetadata)
	for _, partition := range partitions {
		name := partition.Topic

		topic, exists := topics[name]
		if !exists {
			topics[name] = TopicMetadata{
				Name:          name,
				NumPartitions: 1,
			}
			continue
		} else {
			topic.NumPartitions++
			topics[name] = topic
		}
	}

	return topics
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
