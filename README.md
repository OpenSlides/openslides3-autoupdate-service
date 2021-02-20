# OpenSlides3 Autoupdate Service

The OpenSlides3 autoupdate service is a messaging system for openslides 3.x.

To work, the service needs running OpenSlides 3 instance. It reads the data from
redis, gets autoupdates via redis and authenticates request via the openslides
service.

IMPORTANT: The data are sent via an open http-connection. All browsers limit the
amount of open http1.1 connections to a domain. For this service to work, the
browser has to connect to the service with http2 and therefore needs https.


## Install and Start

To start the service, the file `/run/secrets/django` has to exist and has to
contain the same django secret that is used be the server. It can be created
with:
```
mkdir -p /run/secrets
echo "DJANGO_SECRET_KEY='MY_SECRET'" > /run/secrets/django
```


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

### With Docker

```
docker build . --tag openslides3-autoupdate
docker run --network host openslides3-autoupdate
```

This example uses the host network to connect to redis.


## Example requests with curl

Curl needs the flag `-N / --no-buffer` or it can happen, that the output is not
printed immediately. With a self signed certificat (the default of the
autoupdate-service) is also needs the flag `-k / --insecure`.


### Autoupdate

To get all data:

```
curl -N localhost:8002/system/autoupdate
```

To get all data after a change id:

```
curl -N localhost:8002/system/autoupdate?change_id=133188953000
```

To test an authenticated request, login to OpenSlides and find the given session
id. Afterwards the session cookie can be used with curl:

```
curl -N --cookie "OpenSlidesSessionID=3e38tw8kpx64p4gxq80qf2hg4k60ix6w" localhost:8002/system/autoupdate
```


### Projector

To get the projector data for a list of projectors:

```
curl -N localhost:8002/system/projector?projector_ids=1,2,3
```


### Notify

To listen for messages:

```
curl -N localhost:8002/system/notify
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


To send a messages:

To send a notify message, you need to be logged-in. A seesion cookies has to be created and
used. See the --cookie flag above.

```
curl localhost:8002/system/notify/send -d '{"channel_id":"foo:1:0", "name":"title", "to_all":true, "message": "some json"}'
```

The body has to be a valid json object with at least the fields `channel_id`,
`name`, and `message`.

```
{
  "channel_id": string -> channel_id of the sender.
  "name": string -> Name of the message.
  "message": json -> valid json containing the message.
  "to_all": bool -> If true, message is send to every connection.
  "to_users": list[int] -> List of user ids that should receive the message.
  "to_channels": list[string] -> List of channel ids that should receive the message.
}
```


### Applause

The notify system also handels applause. To send an applause, send the following request:

```
curl localhost:8002/system/applause
```

For this to work, a sessionID is required (see above)


### Vote Cache

To send a vote use:

```
curl localhost:8002/system/vote/motion/1 -d `{"data":"Y"}`
```

If the poll is in the correct state and the payload is correct, the vote is
saved into redis.


## Run Test

To run the tests, call:

```
go test ./...
```

To replace the export.json copy it to internal/test/export.json and run:

```
go generate ./...
```

## Environment variables

The service can be configured with the following environment variables:

* `AUTOUPDATE_PORT`:Port to listen on. The default is `8002`.
* `AUTOUPDATE_HOST`: The device where the service starts. The default is an
  empty string which starts the service on every device.
* `MESSAGE_BUS_HOST`: Host of the redis server for reading. The default is
  `localhost`.
* `MESSAGE_BUS_PORT`: Port of the redis server for reading. The default is
  `6379`.
* `REDIS_WRITE_HOST`: Host of the redis server for writing. The default is the
  same as `MESSAGE_BUS_HOST`.
* `REDIS_WRITE_PORT`: Port of the redis server for writing. The default is the
  same as `MESSAGE_BUS_PORT`.
* `APPLAUSE_INTERVAL_MS`: Time to calc the applause in milliseconds (Default:
  `1000`)
* `COOKIE_NAME`: Name of the auth-session-cookie (Default: `OpenSlidesSessionID`).
* `SESSION_PREFIX`: Prefix of the redis session keys (Default: `session:`).
* `FAKE_AUTH`: If set, the auth system is skipped and the given user id is used
  for all requests (Default: `-1`).
