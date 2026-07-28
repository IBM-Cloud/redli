# Redli - A humane alternative to redis-cli

[![Release](https://img.shields.io/github/release/IBM-Cloud/redli.svg)](https://github.com/IBM-Cloud/redli/releases/latest)

## About

Redli is a Go-based alternative to the official Redis-cli application. Its major feature is that it mimics the `redis-cli` command line arguments while also understanding `redis:` protocols and supporting a `--tls` flag, allowing it to connect to TLS/SSL secured Redis without the need for tunnels. It also has a number of flags and environment variables for passing server certificates as files or base64 strings.

_Note_: Passing a certificate turns TLS on by default.

## Installation

### Homebrew

If you use Homebrew, run the following command to install the package:

`brew install redli` 

### Direct download

You can download the binary for your OS from the [releases page](https://github.com/IBM-Cloud/redli/releases). Un-tar the file, then `chmod +x` the binary, and move it into your `$PATH`.

### Compile from source

You can also compile Redli with [Go](https://go.dev) using these steps (Go 1.8+ required):

1. `go get -u github.com/IBM-Cloud/redli`
2. `go install github.com/IBM-Cloud/redli`

#### Usage

Complete instructions are available in the man page (`man redli`).

_Note_: Be aware of interactions with wild cards and special characters in the shell; Quote and escape values as appropriate.

## License

Redli is &copy; IBM Corporation 2018-2024. All rights reserved.

It is available under the [Apache 2 License](https://www.apache.org/licenses/LICENSE-2.0).

The application's build process includes the `commands.json` file from the Redis-docs repository. This file is available under a CC-BY-SA 4.0 license (see [Copyright](https://github.com/redis/docs/blob/main/LICENSE)).

Attribution: The [commands.json](https://github.com/redis/redis-doc/blob/master/commands.json) file is by [Salvatore Sanfillipo](http://invece.org).
