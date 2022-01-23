package cluster

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type ClusterArgs struct {
	BootstrapServer string
	Timeout         int64

	SaslMechanism string
	SaslUsername  string
	SaslPassword  string
}

func getConfigFilepath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Could not determine user home directory")
		os.Exit(1)
	}

	err = os.MkdirAll(filepath.Join(home, ".kcli"), os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return filepath.Join(home, ".kcli", "clusters.conf")
}

type KcliConfig struct {
	Active   string
	Clusters map[string]ClusterArgs
}

func GetActiveClusterArgs() ClusterArgs {
	config := ReadConfig()
	return config.Clusters[config.Active]
}

func UseCluster(name string) {
	config := ReadConfig()
	config.Active = name
	WriteConfig(config)
}

func AddCluster(name string, options ClusterArgs) {
	fmt.Printf("Adding cluster %s\n", name)
	config := ReadConfig()

	config.Clusters[name] = options
	config.Active = name
	WriteConfig(config)
}

func RemoveCluster(name string) {
	fmt.Printf("Removing cluster %s\n", name)

	config := ReadConfig()
	if config.Active == name {
		config.Active = ""
	}

	delete(config.Clusters, name)
	WriteConfig(config)
}

func ListClusters() {
	fmt.Printf("Clusters: \n")

	config := ReadConfig()
	for clusterName := range config.Clusters {
		if clusterName == config.Active {
			fmt.Printf("  - %s (Active)\n", clusterName)
		} else {
			fmt.Printf("  - %s\n", clusterName)
		}
	}
}

func ReadConfig() KcliConfig {
	bytes, err := os.ReadFile(getConfigFilepath())
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		fmt.Println("Error reading kcli file", err)
		os.Exit(1)
	}

	var config KcliConfig = KcliConfig{
		Active:   "",
		Clusters: map[string]ClusterArgs{},
	}
	if bytes != nil {
		err = yaml.Unmarshal(bytes, &config)
		if err != nil {
			fmt.Println("Error while parsing config file", err)
			os.Exit(1)
		}
	}

	return config
}

func WriteConfig(config KcliConfig) {
	bytes, err := yaml.Marshal(config)
	if err != nil {
		fmt.Println("Error writing config file", err)
		os.Exit(1)
	}

	os.WriteFile(getConfigFilepath(), bytes, 0666)
}
