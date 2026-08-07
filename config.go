package main

import (
	"flag"
	"log"
	"os"

	"go.yaml.in/yaml/v4"
)

// Structure for incoming parameters
type config struct {
	// The port to listen on
	Port string `yaml:"port"`
	// The directory to scan for repositories
	ListDir string `yaml:"repo_dir"`
}

type configuration struct {
	// Directory with configuration file
	configFile string
}

// Handle the command line flags and return a configuration struct
func parseFlags() config {
	// Define flags for the configuration file path
	configFile := flag.String(
		"configuration",
		"config.yaml",
		"Path to config file",
	)
	flag.Parse()

	// Process the configuration file
	var cfg config
	cfg, err := readConfigFile(*configFile)
	if err != nil {
		// Handle error, e.g., log and exit
		log.Fatal(err)
	}

	return cfg
}


// Function to read the config file and return a config struct
func readConfigFile(path string) (config, error) {
	// Implement reading the config file and returning a config struct
	// If the file does not exist, return an empty config struct and no error
	// If the file exists but is invalid, return an error
	// If the file exists and is valid, return the config struct with the values from the file

	// Read the config file
	data, err := os.ReadFile(path)
	if err != nil {
		// If the file does not exist, return an empty config struct and no error
		if os.IsNotExist(err) {
			return config{}, nil
		}
		// If the file exists but is invalid, return an error
		return config{}, err
	}

	// Unmarshal the yaml data into a config struct
	var cfg config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		// If the file exists but is invalid, return an error
		return config{}, err
	}
	log.Println(cfg)

	return cfg, nil
}
