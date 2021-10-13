package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func absInt64(n int64) int64 {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func main() {
	bootstrapServerPtr := flag.String("b", "", "Bootstrap server")
	topicPtr := flag.String("t", "", "Topic")
	consumePtr := flag.Bool("c", false, "Consume")
	productPtr := flag.Bool("p", false, "Produce")
	offsetPtr := flag.Int64("o", 0, "Offset (use negative for from-end)")
	flag.Parse()

	if *productPtr && len(flag.Args()) > 0 {
		conn, err := kafka.DialLeader(context.Background(), "tcp", *bootstrapServerPtr, *topicPtr, 0)
		if err != nil {
			fmt.Println("Failed to dial leader", err)
		}
		defer conn.Close()

		_, err = conn.WriteMessages(
			kafka.Message{Value: []byte(flag.Args()[0])},
		)
		if err != nil {
			fmt.Println("Failed to write messages", err)
		}
	}

	if *consumePtr {
		conn, err := kafka.DialLeader(context.Background(), "tcp", *bootstrapServerPtr, *topicPtr, 0)
		if err != nil {
			fmt.Println("Failed to dial leader", err)
		}
		defer conn.Close()

		if *offsetPtr != 0 {
			var seekPos int
			if *offsetPtr > 0 {
				seekPos = kafka.SeekStart
			} else {
				seekPos = kafka.SeekEnd
			}

			_, err := conn.Seek(absInt64(*offsetPtr), seekPos)
			if err != nil {
				fmt.Println("Failed to seek offset", err)
				return
			}
		}

		for {
			msg, err := conn.ReadMessage(10e6)
			if err != nil {
				fmt.Println("Failed to read message", err)
				break
			}

			fmt.Println(string(msg.Value))
		}
	}
}
