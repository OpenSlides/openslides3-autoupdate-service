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

### Autoupdate

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


### Projector

To get the projector data for a list of projectors:

```
curl localhost:8002/system/projector?projector_ids=1,2,3
```


### Notify

The notify system needs logged-in users. A seesion cookies has to be created and
used. See the --cookie flag above.

To listen for messages:

```
curl localhost:8002/system/notify
```

It returns a message like this one to tell the channel id:

`{"channelID": "lVbO2irQ:1:0"}`

The value of channelID is just a string. The format may change in the future.
Currently the first part is a random string that is constant for a server
instance. The second part is the user id. The last part is a counter.

If a message is received, it has the format:

```
{
  "sender_user_id": int -> User id of the sender.
  "sender_channel_id": string -> Channel id of the sender.
  "name": string -> Name of the message.
  "message": json -> Message send by the sender.
}
```

From time to time the server sends empty keep alive messages. A keep alive
message is an empfy json object.


To listen for messages:

```
curl localhost:8002/system/notify/send -d '{"channel_id":"foo:1:0", "to_all":true, "message": "some json"}'
```

The body has to be a valid json object with at least the fields `channel_id` and
`message`.

```
{
  "channel_id": string -> channel_id of the sender.
  "Name": string -> Name of the message.
  "message": json -> valid json containing the message.
  "to_all": bool -> If true, message is send to every connection.
  "to_users": list[int] -> List of user ids that should receive the message.
  "to_channels": list[string] -> List of channel ids that should receive the message.
}
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