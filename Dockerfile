FROM golang:1.15.8-alpine3.12 as builder
LABEL maintainer="OpenSlides Team <info@openslides.com>"

WORKDIR /root/
# Install git, it is needed to install the dependencies
RUN apk --no-cache add git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build ./cmd/autoupdate

# Development build.
FROM builder as development

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
EXPOSE 9012
CMD CompileDaemon -log-prefix=false -build="go build ./cmd/autoupdate" -command="./autoupdate"

# Productive build
FROM alpine:3.13.1

WORKDIR /root/
COPY --from=builder /root/autoupdate .
ENV FORCE_HTTP2 yes
EXPOSE 8002
CMD ./autoupdate
