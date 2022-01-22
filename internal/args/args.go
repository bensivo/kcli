package args

import "bensivo.com/kcli/internal/cluster"

type ProducerArgs struct {
	Topic       string
	Partition   int
	ClusterArgs cluster.ClusterArgs
}

type ConsumerArgs struct {
	Topic       string
	Partition   int
	Offset      int
	ClusterArgs cluster.ClusterArgs
	Exit        bool
}

type CreateTopicArgs struct {
	Topic             string
	Partitions        int
	ReplicationFactor int
	ClusterArgs       cluster.ClusterArgs
}
