package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// PreflightSetup just makes sure the stage is set before starting the application
func PreflightSetup() {
	logStdOut("Preflight complete!")
}

// ServerPreflightSetup just makes sure the stage is set before starting the HTTP server
func ServerPreflightSetup() {
	logStdOut("Server Mode Preflight complete!")
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath CLIOpts) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath.Config)
	checkAndFail(err)
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	readConfig = config

	// Pass in some basic debug information
	logStdOut("[Server Mode]: Configuration file: " + configPath.Config)
	logStdOut("[Server Mode]: Server Port: " + config.Application.Server.Port)
	logStdOut("[Server Mode]: Skip TLS Verification: " + fmt.Sprint(config.Application.Server.SkipTLSVerify))
	logStdOut("[Server Mode]: Timeouts: " + fmt.Sprint(config.Application.Server.Timeouts))
	logStdOut("[Server Mode]: Listening at: " + config.Application.Server.Host + ":" + config.Application.Server.Port + config.Application.Server.BasePath)
	logStdOut("[Server Mode]: Serving path: " + config.Application.Server.FileSystemPath)

	return config, nil
}

// ParseFlags will create and parse the CLI flags
// and return the path to be used elsewhere
func ParseFlags() (CLIOpts, error) {
	// String that contains the configured configuration path
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", "", "path to config file, eg '-config=./config.yml'")

	// Actually parse the flags
	flag.Parse()

	if configPath == "" {
		return CLIOpts{}, Stoerr("[Server Mode]: No server configuration defined! (-config=./config.yml)")
	} else {
		// Validate the path first
		if err := ValidateConfigPath(configPath); err != nil {
			return CLIOpts{}, err
		}
	}

	SetCLIOpts := CLIOpts{
		Config: configPath}

	// Return the configuration path
	return SetCLIOpts, nil
}
