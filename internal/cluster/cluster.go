package cluster

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ClusterArgs struct {
	BootstrapServer string
	Timeout         int64
}

func AddCluster(options ClusterArgs) {
	// TODO: Use a standard ~/.kcli.yaml file
	// TODO: Read the existing file and append a cluster to it
	bytes, err := yaml.Marshal(options)
	if err != nil {
		fmt.Println("Error occurred marshalling", err)
		os.Exit(1)
	}

	os.WriteFile("./kcli.yaml", bytes, 0666)
}
