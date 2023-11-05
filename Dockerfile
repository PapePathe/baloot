# syntax=docker/dockerfile:1

FROM golang:1.21 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
 RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build cmd/websocket/server.go

FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build-stage /app/server /app/server
EXPOSE 7777
USER nonroot:nonroot
ENTRYPOINT ["/app/server"]
