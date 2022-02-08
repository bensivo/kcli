package client

import (
	"fmt"
	"sync"

	"github.com/segmentio/kafka-go"
	"gitlab.com/bensivo/kcli/internal/cluster"
	"gitlab.com/bensivo/kcli/internal/util"
)

type ConsumeManyArgs struct {
	Topic       string
	Partitions  []int
	Offset      int
	ClusterArgs cluster.ClusterArgs
	Exit        bool
}

func ConsumeAllPartitions(cfg ConsumeManyArgs) {
	topics := ListTopics(cfg.ClusterArgs)
	numPartitions := topics[cfg.Topic].NumPartitions

	var wg sync.WaitGroup
	for i := 0; i < numPartitions; i++ {
		wg.Add(1)
		go Consume(&wg, ConsumeArgs{
			Topic:       cfg.Topic,
			Partition:   i,
			Offset:      cfg.Offset,
			ClusterArgs: cfg.ClusterArgs,
			Exit:        cfg.Exit,
		})
	}

	wg.Wait()
}

type ConsumeArgs struct {
	Topic       string
	Partition   int
	Offset      int
	ClusterArgs cluster.ClusterArgs
	Exit        bool
}

func Consume(wg *sync.WaitGroup, cfg ConsumeArgs) {
	defer wg.Done()

	conn := DialLeader(cfg.ClusterArgs, cfg.Topic, cfg.Partition)
	defer conn.Close()

	// Set the reference seek position based on whether the given offset is positive or negative
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

	// Get the actual current offset from the leader
	last, err := conn.ReadLastOffset()
	if err != nil {
		fmt.Println("Failed to read offset", err)
	}
	if last == 0 {
		if cfg.Exit {
			return
		}
	}

	// Read messages
	for {
		msg, err := conn.ReadMessage(10e6)
		if err != nil {
			fmt.Println("Failed to read message", err)
			break
		}

		fmt.Printf("%s\n", string(msg.Value))

		// TODO: Do we need to re-read the offset after every message?
		// last, err = conn.ReadLastOffset()
		// if err != nil {
		// 	fmt.Println("Failed to read offset", err)
		// }

		if msg.Offset == last-1 {
			if cfg.Exit {
				return
			}
		}
	}
}

// // Uses the Reader API instead of the connection API.
// // Allows the user to use consumer groups, but also doesn't
// func ConsumeWithGroup(cfg client.ConsumerArgs) {
// 	// Make an initial connection to the leader to get the current offset
// 	ctx, _ := context.WithTimeout(context.Background(), time.Duration(cfg.ClusterArgs.Timeout)*time.Second)
// 	bootstrapServer := cfg.ClusterArgs.BootstrapServer
// 	conn, err := kafka.DialLeader(ctx, "tcp", bootstrapServer, cfg.Topic, cfg.Partition) // TODO: Write to round-robin partitions
// 	if err != nil {
// 		fmt.Println("Failed to dial leader", err)
// 		os.Exit(1)
// 	}
// 	defer conn.Close()

// 	// Calculate actual offset position from the given parameter
// 	// i.e. offset 0 -> earliest available
// 	// 		offset -1 -> latest available - 1
// 	var seekPos int
// 	if cfg.Offset != 0 {
// 		if cfg.Offset > 0 {
// 			seekPos = kafka.SeekStart
// 		} else {
// 			seekPos = kafka.SeekEnd
// 		}

// 		_, err := conn.Seek(int64(util.AbsInt(cfg.Offset)), seekPos)
// 		if err != nil {
// 			fmt.Println("Failed to seek offset", err)
// 			return
// 		}
// 	}

// 	var last int64
// 	if cfg.Exit {
// 		last, err = conn.ReadLastOffset()
// 		if err != nil {
// 			fmt.Println("Failed to read offset", err)
// 		}
// 	}

// 	// Make another connection to actually read messages from the broker
// 	reader := kafka.NewReader(kafka.ReaderConfig{
// 		Brokers: []string{bootstrapServer},
// 		Topic:   cfg.Topic,
// 	})

// 	reader.SetOffset(int64(seekPos))
// 	for {
// 		m, err := reader.ReadMessage(context.Background())
// 		if err != nil {
// 			fmt.Println(err)
// 			break
// 		}
// 		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

// 		if cfg.Exit && m.Offset == last-1 {
// 			os.Exit(0)
// 		}
// 	}

// 	if err := reader.Close(); err != nil {
// 		fmt.Println("failed to close reader:", err)
// 	}
// }
