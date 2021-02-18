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

RUN mkdir -p /run/secrets
RUN echo "DJANGO_SECRET_KEY='development'" > /run/secrets/django

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
EXPOSE 8002
CMD CompileDaemon -log-prefix=false -build="go build ./cmd/autoupdate" -command="./autoupdate"

# Productive build
FROM alpine:3.13.2

WORKDIR /root/
COPY --from=builder /root/autoupdate .
EXPOSE 8002
CMD ./autoupdate
