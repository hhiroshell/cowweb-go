FROM golang:1.14-buster as builder

WORKDIR /build

COPY ./cmd ./cmd
#COPY ./pkg ./pkg
COPY ./go.mod ./go.mod

# create dependency cache for repeated build during local development.
RUN go mod download
RUN go build -v -o ./cowweb ./cmd/cowweb/main.go

FROM debian:buster-slim

RUN groupadd -r cowweb && useradd -r -g cowweb cowweb \
 && chown cowweb:cowweb ./
USER cowweb

WORKDIR /home/cowweb

COPY --from=builder --chown=cowweb:cowweb /build/cowweb .

ENTRYPOINT ["./cowweb"]
