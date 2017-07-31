Bloom Public Data Sources
=========================

This project is all the data-source specific code for loading
BloomAPI's public data APIs. Each data-loader is its own binary.

The `runner` directory contains a tool to run all the binaries
available to update any given datasource. All other directories
contain data-source specific loading code.

To build the loader binaries + runner:

1. Run `go run configure.go`
2. Run `make` (make sure that your GOPATH environment variable is set)
