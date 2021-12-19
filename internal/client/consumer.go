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

func Consume(cfg args.ConsumerArgs) {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(cfg.ClusterArgs.Timeout)*time.Second)
	bootstrapServer := cfg.ClusterArgs.BootstrapServer
	conn, err := kafka.DialLeader(ctx, "tcp", bootstrapServer, cfg.Topic, cfg.Partition) // TODO: Write to round-robin partitions
	if err != nil {
		fmt.Println("Failed to dial leader", err)
		os.Exit(1)
	}
	defer conn.Close()

	if cfg.Offset != 0 {
		var seekPos int
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

	last, err := conn.ReadLastOffset()
	if err != nil {
		fmt.Println("Failed to read offset", err)
	}
	if last == 0 {
		if cfg.Exit {
			os.Exit(0)
		}
	}

	for {
		msg, err := conn.ReadMessage(10e6)
		if err != nil {
			fmt.Println("Failed to read message", err)
			break
		}

		// fmt.Printf("%d: %s\n", msg.Offset, string(msg.Value)) // TODO: Don't actually print offsets
		fmt.Printf("%s\n", string(msg.Value)) // TODO: Don't actually print offsets

		last, err = conn.ReadLastOffset()
		if err != nil {
			fmt.Println("Failed to read offset", err)
		}

		if msg.Offset == last-1 {
			if cfg.Exit {
				os.Exit(0)
			}
		}
	}
}
