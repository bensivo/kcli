package client

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"bensivo.com/kcli/internal/args"
	"bensivo.com/kcli/internal/cluster"
	"github.com/segmentio/kafka-go"
)

func ListTopics(cfg cluster.ClusterArgs) {
	bootstrapServer := cfg.BootstrapServer
	conn, _ := kafka.Dial("tcp", bootstrapServer)
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

func CreateTopic(cfg args.CreateTopicArgs) {
	conn, err := kafka.Dial("tcp", cfg.ClusterArgs.BootstrapServer)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicArgss := []kafka.TopicConfig{
		{
			Topic:             cfg.Topic,
			NumPartitions:     cfg.Partitions,
			ReplicationFactor: cfg.ReplicationFactor,
		},
	}

	err = controllerConn.CreateTopics(topicArgss...)
	if err != nil {
		panic(err.Error())
	}
}
