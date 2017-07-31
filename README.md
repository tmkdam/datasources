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

Every binary requires access to a local copy of `config.toml`, `dbmapping.yaml`, and `searchmapping.yaml`. Alternatively, you can reference a folder that contains `config.toml` by setting the environment variable `BLOOM_CONFIG`. For example, the following would be used to add a data source to BloomAPI (from the directory output by make `go/bin/linux_amd64/datasources/sources/usgov.hhs.npi`):

```
BLOOM_CONFIG=../../ ./npi bootstrap
BLOOM_CONFIG=../../ ./npi fetch
BLOOM_CONFIG=../../ ./npi search-index
```

It may make sense to also update a datasource via cron. If you wanted to update a given datasource, run:

```
BLOOM_CONFIG=../../ ./npi fetch
BLOOM_CONFIG=../../ ./npi search-index
```

Every command will have the commands:

* `bootstrap`: creates database tables for the data source
* `fetch`: downloads + imports the datasource and related metadata
* `search-index`: loads the dataset into ElasticSearch and allows it to be queried by BloomAPI
* `drop`: drops the database tables and related metadata
* `search-drop`: drops the ElasticSearch index for this source
* `schema`: Generates a raw schema file for the datasource (given an implementation of fetch functionality). This is a tool for development rather than for use in production.