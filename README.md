# User Service

This is the User service

Generated with

```
micro new user --namespace=github.com/tonymj76/micro-postgres --type=service
```

### which is Inspired by `johanbrandhorst` [blog](https://jbrandhorst.com/post/postgres/)


## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: github.com/tonymj76/micro-postgres.service.user
- Type: service
- Alias: user

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./user-service
```

Build a docker image
```
make docker
```