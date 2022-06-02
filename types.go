package main

import (
	"time"
)

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

// CLIOpts defines the CLI Arguements
type CLIOpts struct {
	Config string
	Source string
	Dir    string
}

// Config struct for webapp config at the top level
type Config struct {
	Application ApplicationYaml `yaml:"app"`
}

// ApplicationYaml is what is defined for this Application when running as a server
type ApplicationYaml struct {
	ServerEnabled bool     `yaml:"server_enabled"`
	Server        Server   `yaml:"server,omitempty"`
	Assets        []Assets `yaml:"assets,omitempty"`
}

// Server configures the HTTP server
type Server struct {
	// Host is the local machine IP Address to bind the HTTP Server to
	Host string `yaml:"host"`

	BasePath       string `yaml:"base_path"`
	FileSystemPath string `yaml:"fs_path"`

	// Port is the local machine TCP Port to bind the HTTP Server to
	Port    string `yaml:"port"`
	Timeout struct {
		// Server is the general server timeout to use
		// for graceful shutdowns
		Server time.Duration `yaml:"server"`

		// Write is the amount of time to wait until an HTTP server
		// write opperation is cancelled
		Write time.Duration `yaml:"write"`

		// Read is the amount of time to wait until an HTTP server
		// read operation is cancelled
		Read time.Duration `yaml:"read"`

		// Read is the amount of time to wait
		// until an IDLE HTTP session is closed
		Idle time.Duration `yaml:"idle"`
	} `yaml:"timeouts"`
}

// Assets is the slice of a dictionary of assets that will be downloaded from a source to a destination path and served via the web server
type Assets struct {
	// Source is the source URL to download the assets from
	Source string `yaml:"source"`
	// Destination is the local path to download the assets to
	Destination string `yaml:"destination"`
	// Overwrite is a boolean to determine if the destination file should be overwritten
	Overwrite bool `yaml:"overwrite,omitempty"`
}

// HTTPReqInfo describes info about HTTP request
type HTTPReqInfo struct {
	// GET etc.
	method  string
	uri     string
	referer string
	ipaddr  string
	// response code, like 200, 404
	code int
	// how long did it take to
	duration  time.Duration
	userAgent string
}
