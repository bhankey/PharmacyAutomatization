FROM golang:1.17

RUN mkdir "build"

COPY ./ /h1:KmX6BPdI08NWTb3/sm4ZGu5ShLoqVDhKgpiN924inxo=
WORKDIR /github.com/aws/aws-lambda-go v1.13.3/go.mod h1:4UKl9IzQMoD+QF79YdCuzCwp8VbmG4VAQwij/eHl5CU=
github.com/aws/aws-sdk-go v0.14.3216/go.mod h1:KmX6BPdI08NWTb3/sm4ZGu5ShLoqVDhKgpiN924inxo=

RUN GOOGS=linux CGO_ENABLED=0 go build -o ./app ./cmd/server/main.go

FROM alpine:latest
COPY --from=0 /build/app ./
COPY ./config/config-docker.yaml ./config/config.yaml

CMD ["./app"]
