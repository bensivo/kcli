package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go/sasl/scram"
	"gitlab.com/bensivo/kcli/internal/cluster"
)

func GetSaslMechanism(cfg cluster.ClusterArgs) sasl.Mechanism {
	if strings.ToLower(cfg.SaslMechanism) == "plain" {
		return plain.Mechanism{
			Username: cfg.SaslUsername,
			Password: cfg.SaslPassword,
		}
	}
	if strings.ToLower(cfg.SaslMechanism) == "scram-sha-512" {
		mechanism, err := scram.Mechanism(scram.SHA512, cfg.SaslUsername, cfg.SaslPassword)
		if err != nil {
			fmt.Println("Error configuring scram-sha-512 auth")
			os.Exit(1)
		}
		return mechanism
	}
	if strings.ToLower(cfg.SaslMechanism) == "scram-sha-256" {
		mechanism, err := scram.Mechanism(scram.SHA256, cfg.SaslUsername, cfg.SaslPassword)
		if err != nil {
			fmt.Println("Error configuring scram-sha-256 auth")
			os.Exit(1)
		}
		return mechanism
	}

	return nil
}

func GetTLSConfig(cfg cluster.ClusterArgs) *tls.Config {
	if !cfg.SSLEnabled {
		return nil
	}

	var tlsConfig tls.Config = tls.Config{
		Certificates: []tls.Certificate{},
	}

	if cfg.SSLCaCertificatePath != "" {
		caCert, err := ioutil.ReadFile("./test/clusters/ssl/ca_authority/ca-cert")
		if err != nil {
			log.Fatal(err)
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		tlsConfig.RootCAs = caCertPool
	}

	if cfg.SSLSkipVerification {
		tlsConfig.InsecureSkipVerify = true
	}

	return &tlsConfig
}

func GetDialer(cfg cluster.ClusterArgs) *kafka.Dialer {
	fmt.Println("Connecting to cluster " + cfg.BootstrapServer)

	dialer := &kafka.Dialer{
		Timeout:       time.Second * 10,
		DualStack:     true,
		SASLMechanism: GetSaslMechanism(cfg),
		TLS:           GetTLSConfig(cfg),
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
