package config

import (
	"flag"
	"os"
)

type ClusterConfig struct {
	BootstrapServer string
	Timeout         int64
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

type CreateTopicConfig struct {
	Topic             string
	Partitions        int
	ReplicationFactor int
	ClusterConfig     ClusterConfig
}

func GetProducerConfig() ProducerConfig {
	flags := flag.NewFlagSet("producer", flag.ExitOnError)
	bootstrapServer := flags.String("b", "", "BootstrapServer")
	connectionTimeout := flags.Int64("timeout", 10, "ConnectionTimeout")
	topic := flags.String("t", "", "Topic")
	partition := flags.Int("p", 0, "Partition")
	flags.Parse(os.Args[2:])

	return ProducerConfig{
		Topic:     *topic,
		Partition: *partition,
		ClusterConfig: ClusterConfig{
			BootstrapServer: *bootstrapServer,
			Timeout:         *connectionTimeout,
		},
	}
}

func GetConsumerConfig() ConsumerConfig {
	flags := flag.NewFlagSet("consumer", flag.ExitOnError)
	bootstrapServer := flags.String("b", "", "BootstrapServer")
	connectionTimeout := flags.Int64("timeout", 10, "ConnectionTimeout")
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
			Timeout:         *connectionTimeout,
		},
	}
}

func GetClusterConfig() ClusterConfig {
	flags := flag.NewFlagSet("cluster", flag.ExitOnError)
	bootstrapServer := flags.String("b", "", "BootstrapServer")
	connectionTimeout := flags.Int64("timeout", 10, "ConnectionTimeout")

	flags.Parse(os.Args[2:])

	return ClusterConfig{
		BootstrapServer: *bootstrapServer,
		Timeout:         *connectionTimeout,
	}
}

func GetCreateTopicConfig() CreateTopicConfig {
	flags := flag.NewFlagSet("cluster", flag.ExitOnError)
	bootstrapServer := flags.String("b", "", "BootstrapServer")
	connectionTimeout := flags.Int64("timeout", 10, "ConnectionTimeout")
	topic := flags.String("t", "", "Topic")
	partitions := flags.Int("p", 1, "Partitions (default 1)")
	replicationFactor := flags.Int("r", 1, "ReplicationFactor (default 1)")

	flags.Parse(os.Args[3:])

	return CreateTopicConfig{
		Topic:             *topic,
		Partitions:        *partitions,
		ReplicationFactor: *replicationFactor,
		ClusterConfig: ClusterConfig{
			BootstrapServer: *bootstrapServer,
			Timeout:         *connectionTimeout,
		},
	}
}
