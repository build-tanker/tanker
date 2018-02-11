# Tanker

Tanker is an Open source, Self Hosted, Mobile Releases Deployment tool for iOS and Android.

*WIP : We are not ready for production yet. You can sign up for our [newsletter here](https://goo.gl/forms/Ck0lnzdWQNyfD5Il2), we will reach out to you as soon as we are ready.*

## Getting Started

Tanker consists of two different components.

[Shipper](https://github.com/sudhanshuraheja/shipper), which connects to your CI or your local machine and uploads binaries. 

[Tanker](https://github.com/sudhanshuraheja/tanker), is the backend that saves binaries on Google Cloud Storage and makes it available for user's who should have access.

Please follow the instructions for both the repos.

### Prerequisites

Here are the things that you would require before you get started

1. [Install git](https://www.atlassian.com/git/tutorials/install-git)
1. [Install golang](https://golang.org/doc/install)
1. [Install docker](https://docs.docker.com/install/#supported-platforms), we use it both for deployment and development

### Installing

Clone the repo and build it

```bash
git clone https://github.com/sudhanshuraheja/tanker.git
make build_fresh
```

Start the binary by running tanker

```bash
tanker start
```

## Running the tests

If you would like to run the automated tests for the complete package, run this

```bash
make coverage
open ./coverage.html
```

### And coding style tests

We use the default golang coding conventions. Run the following to test for those

```bash
make fmt
make vet
make lint
```

## Deployment

We will have deployment instructions as soon as we are ready. You can sign up for our [newsletter here](https://goo.gl/forms/Ck0lnzdWQNyfD5Il2), we will reach out to you as soon as we are ready.

## Built With

* [DEP](https://github.com/golang/dep) - For dependency management
* [CLI](github.com/urfave/cli) - For accessing the binary on CLI
* [VIPER](github.com/spf13/viper) - For configuration management
* [LOGRUS](github.com/sirupsen/logrus) - For logging
* [NEGRONI](github.com/urfave/negroni) - HTTP Middleware
* [MUX](github.com/gorilla/mux) - For routing each request to the correct place
* [PQ](github.com/lib/pq) - SQL driver for postgres
* [SQLX](github.com/jmoiron/sqlx) - For connecting to postgres
* [MIGRATE](github.com/mattes/migrate) - For migrating postgres
* [GC STORAGE](cloud.google.com/go/storage) - For storing files on google cloud
* [GO.UUID](github.com/satori/go.uuid) - For getting UUIDs
* [TESTIFY](github.com/stretchr/testify) - For asserting tests
* [GO-SQLMOCK](github.com/DATA-DOG/go-sqlmock) - For mocking postgres

## Contributing

Please read [CONTRIBUTING.md](https://github.com/sudhanshuraheja/tanker/blob/master/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](https://semver.org/spec/v2.0.0.html) for versioning based on the recommendation from [Dave Chaney](https://dave.cheney.net/2016/06/24/gophers-please-tag-your-releases). For the versions available, see the [tags on this repository](https://github.com/sudhanshuraheja/tanker/tags).

## Authors

* [Sudhanshu Raheja](https://github.com/sudhanshuraheja)

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](https://github.com/sudhanshuraheja/tanker/blob/master/LICENSE) file for details

## Acknowledgments

* Hat tip to fabric.io and hockeyapp, we have been looking at them for inspiration
