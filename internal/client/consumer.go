package client

import (
	"context"
	"fmt"
	"os"
	"time"

	"bensivo.com/kcli/internal/args"
	"bensivo.com/kcli/internal/util"
	"github.com/segmentio/kafka-go"
)

//  Uses the Reader API instead of the connection API.
// Allows the user to use consumer groups, but also doesn't
func Consume(cfg args.ConsumerArgs) {
	// Make an initial connection to the leader to get the current offset
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(cfg.ClusterArgs.Timeout)*time.Second)
	bootstrapServer := cfg.ClusterArgs.BootstrapServer
	conn, err := kafka.DialLeader(ctx, "tcp", bootstrapServer, cfg.Topic, cfg.Partition) // TODO: Write to round-robin partitions
	if err != nil {
		fmt.Println("Failed to dial leader", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Calculate actual offset position from the given parameter
	// i.e. offset 0 -> earliest available
	// 		offset -1 -> latest available - 1
	var seekPos int
	if cfg.Offset != 0 {
		if cfg.Offset > 0 {
			seekPos = kafka.SeekStart
		} else {
			seekPos = kafka.SeekEnd
		}

		_, err := conn.Seek(int64(util.AbsInt(cfg.Offset)), seekPos)
		if err != nil {
			fmt.Println("Failed to seek offset", err)
			return
		}
	}

	var last int64
	if cfg.Exit {
		last, err = conn.ReadLastOffset()
		if err != nil {
			fmt.Println("Failed to read offset", err)
		}
	}

	// Make another connection to actually read messages from the broker
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{bootstrapServer},
		Topic:   cfg.Topic,
	})

	reader.SetOffset(int64(seekPos))
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

		if cfg.Exit && m.Offset == last-1 {
			os.Exit(0)
		}
	}

	if err := reader.Close(); err != nil {
		fmt.Println("failed to close reader:", err)
	}
}
