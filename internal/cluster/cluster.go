package cluster

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ClusterArgs struct {
	BootstrapServer string
	Timeout         int64
}

type ConfigFileContents map[string]ClusterArgs

func AddCluster(name string, options ClusterArgs) {
	fmt.Printf("Adding cluster %s\n", name)
	config := ReadConfig()

	config[name] = options
	WriteConfig(config)
}

func RemoveCluster(name string) {
	fmt.Printf("Removing cluster %s\n", name)
	config := ReadConfig()
	delete(config, name)
	WriteConfig(config)
}

func ListClusters() {
	fmt.Printf("Cluster: \n")

	config := ReadConfig()
	for clusterName := range config {
		fmt.Printf("  - %s\n", clusterName)
	}

}

func ReadConfig() ConfigFileContents {
	bytes, err := os.ReadFile("kcli.yaml")
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		fmt.Println("Error reading kcli file", err)
		os.Exit(1)
	}

	var config ConfigFileContents = ConfigFileContents{}
	if bytes != nil {
		err = yaml.Unmarshal(bytes, &config)
		if err != nil {
			fmt.Println("Error while parsing config file", err)
			os.Exit(1)
		}
	}

	return config
}

func WriteConfig(config ConfigFileContents) {
	bytes, err := yaml.Marshal(config)
	if err != nil {
		fmt.Println("Error writing config file", err)
		os.Exit(1)
	}

	os.WriteFile("kcli.yaml", bytes, 0666)
}
