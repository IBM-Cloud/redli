# Redli - a humane alternative to redis-cli

## About

Redli is a Go-based alternative to the official Redis-cli application. It's major feature is that it understands redis: and rediss: URLs and, unlike Redis-cli currently, can connect to TLS/SSL secured Redis without the need for tunnels. It also has a number of flags and environment variables for passing server certificates over as files or base64 strings.

## Usage

```
redli [<flags>] <url>
```

### Flags:
  * `--help` Show context-sensitive help (also try --help-long and --help-man).
  * `--debug` Enable debug mode.
  * `--long` Enable long prompt with host/port
  * `--certfile=CERTFILE` Self-signed certificate file for validation
  * `--certb64=CERTB64` Self-signed certificate string as base64 for validation

#### Args:

* `<url>`  URL to connect To. The URL is, at this point in development, mandatory. It follow the format of [the provisional IANA spec for Redis URLs](https://www.iana.org/assignments/uri-schemes/prov/redis), but with the option to denote a TLS secured connection with the protocol rediss:.

## License

Redli is released under the Apache 2 License.

