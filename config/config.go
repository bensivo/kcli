package config

import (
	"flag"
	"os"
)

type ClusterConfig struct {
	BootstrapServer string
}

type ProducerConfig struct {
	Topic         string
	Partition     int
	ClusterConfig ClusterConfig
}

type ConsumerConfig struct {
	Topic         string
	Partition     int
	Offset        int
	ClusterConfig ClusterConfig
}

func GetProducerConfig() ProducerConfig {
	flags := flag.NewFlagSet("producer", flag.ExitOnError)
	bootstrapServer := flags.String("b", "", "BootstrapServer")
	topic := flags.String("t", "", "Topic")
	partition := flags.Int("p", 0, "Partition")
	flags.Parse(os.Args[2:])

	return ProducerConfig{
		Topic:     *topic,
		Partition: *partition,
		ClusterConfig: ClusterConfig{
			BootstrapServer: *bootstrapServer,
		},
	}
}

func GetConsumerConfig() ConsumerConfig {
	flags := flag.NewFlagSet("consumer", flag.ExitOnError)
	bootstrapServer := flags.String("b", "", "BootstrapServer")
	topic := flags.String("t", "", "Topic")
	partition := flags.Int("p", 0, "Partition")
	offset := flags.Int("o", 0, "Offset")
	flags.Parse(os.Args[2:])

	return ConsumerConfig{
		Topic:     *topic,
		Partition: *partition,
		Offset:    *offset,
		ClusterConfig: ClusterConfig{
			BootstrapServer: *bootstrapServer,
		},
	}
}
