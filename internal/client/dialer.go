package client

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"bensivo.com/kcli/internal/cluster"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func GetDialer(cfg cluster.ClusterArgs) *kafka.Dialer {
	fmt.Println("Connecting to cluster " + cfg.BootstrapServer)

	dialer := &kafka.Dialer{
		Timeout: time.Second * 10,
	}
	if strings.ToLower(cfg.SaslMechanism) == "plain" {
		dialer.SASLMechanism = plain.Mechanism{
			Username: cfg.SaslUsername,
			Password: cfg.SaslPassword,
		}
	}
	return dialer
}

func Dial(cfg cluster.ClusterArgs) *kafka.Conn {
	dialer := GetDialer(cfg)

	conn, err := dialer.Dial("tcp", cfg.BootstrapServer)
	if err != nil {
		fmt.Println("Failed to dial leader", err)
		os.Exit(1)
	}

	return conn
}

func DialLeader(cfg cluster.ClusterArgs, topic string, partition int) *kafka.Conn {
	dialer := GetDialer(cfg)

	fmt.Printf("Dialing leader for %s:%s:%d\n", cfg.BootstrapServer, topic, partition)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	conn, err := dialer.DialLeader(ctx, "tcp", cfg.BootstrapServer, topic, partition)
	if err != nil {
		fmt.Println("Failed to dial leader", err)
		os.Exit(1)
	}

	return conn
}
