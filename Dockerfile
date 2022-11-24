FROM golang:1.19-buster as builder

WORKDIR /build

COPY ./main.go ./main.go
COPY ./cmd ./cmd
COPY ./pkg ./pkg
COPY ./go.mod ./go.mod

# create dependency cache for repeated build during local development.
RUN go mod tidy
RUN go build -v -o ./cowweb .

FROM gcr.io/distroless/base

COPY --from=builder /build/cowweb .

ENTRYPOINT ["./cowweb"]
