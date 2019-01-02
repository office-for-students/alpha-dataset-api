alpha-dataset-api
==================
An app for accessing the Office for Students data (specifically institutions and courses)

### Installation

As the service are written in Go, make sure you have version 1.10.0 or greater installed.

Using [Homebrew](https://brew.sh/) to install go
* Run `brew install go` or `brew upgrade go`
* Set your `GOPATH` environment variable, this specifies the location of your workspace

#### Database

* Run `brew install mongodb`
* Run `brew services restart mongodb`

#### Loading Mongo Data from scripts

See [scripts](https://github.com/office-for-students/alpha-scripts/tree/develop/README.md)

#### Running Service

* Run `make debug`

#### Running tests

* Run `make test`

### Configuration

| Environment variable      | Default                | Description
| ------------------------- | ---------------------- | ----------------------------------------------------------------
| BIND_ADDR                 | :10000                 | The host and port to bind to
| HOST_NAME                 | http://localhost       | The scheme and host name
| GRACEFUL_SHUTDOWN_TIMEOUT | 5s                     | The graceful shutdown timeout in seconds
| MONGO_ADDR                | localhost:27017        | The address of the mongo database to retrieve dataset data from
| MONGO_COLLECTION          | localhost:27017        | The address of the mongo database to retrieve dataset data from
| MONGO_DATABASE            | localhost:27017        | The address of the mongo database to retrieve dataset data from


### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

See [LICENSE](LICENSE.md) for details.
