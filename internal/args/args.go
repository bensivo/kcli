package args

import (
	"flag"
	"os"
)

type ClusterArgs struct {
	BootstrapServer string
	Timeout         int64
}

type ProducerArgs struct {
	Topic       string
	Partition   int
	ClusterArgs ClusterArgs
}

type ConsumerArgs struct {
	Topic       string
	Partition   int
	Offset      int
	ClusterArgs ClusterArgs
	Exit        bool
}

type CreateTopicArgs struct {
	Topic             string
	Partitions        int
	ReplicationFactor int
	ClusterArgs       ClusterArgs
}

func GetProducerArgs() ProducerArgs {
	flags := flag.NewFlagSet("producer", flag.ExitOnError)
	bootstrapServer := flags.String("b", "", "BootstrapServer")
	connectionTimeout := flags.Int64("timeout", 10, "ConnectionTimeout")
	topic := flags.String("t", "", "Topic")
	partition := flags.Int("p", 0, "Partition")
	flags.Parse(os.Args[2:])

	return ProducerArgs{
		Topic:     *topic,
		Partition: *partition,
		ClusterArgs: ClusterArgs{
			BootstrapServer: *bootstrapServer,
			Timeout:         *connectionTimeout,
		},
	}
}

func GetConsumerArgs() ConsumerArgs {
	flags := flag.NewFlagSet("consumer", flag.ExitOnError)
	bootstrapServer := flags.String("b", "", "BootstrapServer")
	connectionTimeout := flags.Int64("timeout", 10, "ConnectionTimeout")
	topic := flags.String("t", "", "Topic")
	partition := flags.Int("p", 0, "Partition")
	offset := flags.Int("o", 0, "Offset")
	exit := flags.Bool("e", false, "Exit")
	flags.Parse(os.Args[2:])

	return ConsumerArgs{
		Topic:     *topic,
		Partition: *partition,
		Offset:    *offset,
		Exit:      *exit,
		ClusterArgs: ClusterArgs{
			BootstrapServer: *bootstrapServer,
			Timeout:         *connectionTimeout,
		},
	}
}

func GetTopicListArgs() ClusterArgs {
	flags := flag.NewFlagSet("cluster", flag.ExitOnError)
	bootstrapServer := flags.String("b", "", "BootstrapServer")
	connectionTimeout := flags.Int64("timeout", 10, "ConnectionTimeout")

	flags.Parse(os.Args[3:])

	return ClusterArgs{
		BootstrapServer: *bootstrapServer,
		Timeout:         *connectionTimeout,
	}
}

func GetCreateTopicArgs() CreateTopicArgs {
	flags := flag.NewFlagSet("cluster", flag.ExitOnError)
	bootstrapServer := flags.String("b", "", "BootstrapServer")
	connectionTimeout := flags.Int64("timeout", 10, "ConnectionTimeout")
	topic := flags.String("t", "", "Topic")
	partitions := flags.Int("p", 1, "Partitions (default 1)")
	replicationFactor := flags.Int("r", 1, "ReplicationFactor (default 1)")

	flags.Parse(os.Args[3:])

	return CreateTopicArgs{
		Topic:             *topic,
		Partitions:        *partitions,
		ReplicationFactor: *replicationFactor,
		ClusterArgs: ClusterArgs{
			BootstrapServer: *bootstrapServer,
			Timeout:         *connectionTimeout,
		},
	}
}
