# Redli - a humane alternative to redis-cli

[![Release](https://img.shields.io/github/release/IBM-Cloud/redli.svg)](https://github.com/IBM-Cloud/redli/releases/latest)

## About

Redli is a Go-based alternative to the official Redis-cli application. It's major feature is that it mimics the redis-cli command line argurments while also understanding rediss: protocols and supporting a `--tls` flag allowing it to connect to TLS/SSL secured Redis without the need for tunnels. It also has a number of flags and environment variables for passing server certificates over as files or base64 strings. Note, passing a certificate turns TLS on by default.

## Installation

You can download the binary for your OS from the [releases page](https://github.com/IBM-Cloud/redli/releases). Un-tar the file, then `chmod +x` the binary and move it to your path.

You can also compile Redli with **go** using these steps (Go 1.8+ required):

1. `go get -u github.com/IBM-Cloud/redli`
2. `go install github.com/IBM-Cloud/redli`

## Usage

```text
 redli [<flags>] [<commands>...]

      --help               Show context-sensitive help (also try --help-long and --help-man).
      --debug              Enable debug mode.
      --long               Enable long prompt with host/port
  -u, --uri=URI            URI to connect to
  -h, --host="127.0.0.1"   Host to connect to
  -p, --port=6379          Port to connect to
  -a, --auth=AUTH          Password to use when connecting
  -n, --ndb=0              Redis database to access
      --tls                Enable TLS/SSL
      --skipverify         Insecure option to skip server certificate validation
      --certfile=CERTFILE  Self-signed certificate file for validation
      --certb64=CERTB64    Self-signed certificate string as base64 for validation
      --raw                Produce raw output
      --eval=EVAL          Evaluate a Lua script file, follow with keys a , and args
      
Args:
  [<commands>]  Redis commands and values
```

* `URI`  URI to connect To. It follow the format of [the provisional IANA spec for Redis URLs](https://www.iana.org/assignments/uri-schemes/prov/redis), but with the option to denote a TLS secured connection with the protocol rediss:.

e.g. `INFO KEYSPACE`

Be aware of interactions with wild cards and special characters in the shell; quote and escape as appropriate.

## License

Redli is (c) IBM Corporation 2018. All rights reserved.

Redli is released under the Apache 2 License.

Attribution: The `commands.json` file is by Salvatore Sanfillipo.

In the process of building the application, the commands.json file of the Redis-docs repository is retrieved and incorporated into the code. This file is distributed under a CC-BY-SA 4.0 license (see [Copyright](https://github.com/antirez/redis-doc/blob/master/COPYRIGHT)).
