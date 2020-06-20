# OpenSlides3 Autoupdate Service

The OpenSlides3 autoupdate service is a messaging system for openslides 3.x. It
replaces django channels.

To work, the service needs running OpenSlides 3 instance. It reads the data from
redis, gets autoupdates via redis and authenticates request via the openslides service.


## Install and Start

### With Go

When Go is installed, the service can be installed with:

```
go get github.com/OpenSlides/openslides3-autoupdate-service/cmd/autoupdate
```

This command is case sensitive!

After this command, the binary is installed inside the
[gopath](https://github.com/golang/go/wiki/GOPATH). When the gopath is
integrated correctly, the autoupdate service can be started with:

```
autoupdate
```

### With Checked out repository and Go

When the repository is checked out, the service can be build and started with

```
go build ./cmd/autoupdate && ./autoupdate
```


## Example requests with curl

To get all data:

```
curl localhost:8002/system/autoupdate
```

To get all data after a change id:

```
curl localhost:8002/system/autoupdate?change_id=133188953000
```

To test an authenticated request, login to OpenSlides and find the given session
id. Afterwards the session cookie can be used with curl:

```
curl --cookie "OpenSlidesSessionID=3e38tw8kpx64p4gxq80qf2hg4k60ix6w" localhost:8002/system/autoupdate
```


## Run Test

To run the tests, call:

```
go test ./...
```

## Environment variables

The service can be configured with the following environment variables:

* `LISTEN_HTTP_ADDR`: Host and port on which the services listen on (Default: `:8002`).
* `REDIS_ADDR`: Host and port of the redis service (Default: `localhost:6379`).
* `WORKER_ADDR`: Host and port of the OpenSlides worker (Default: `http://localhost:8000`).
* `KEEP_ALIVE_DURATION`: Time in seconds how often a keep alive packet is sent
  to the client. The value `0` means not to send any keep alive packages.
  (Default: `30`).