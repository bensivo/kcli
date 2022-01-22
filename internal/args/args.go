package args

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
