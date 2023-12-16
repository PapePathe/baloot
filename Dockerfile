# syntax=docker/dockerfile:1

FROM golang:1.21 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o zinx cmd/websocket/server.go
RUN CGO_ENABLED=0 GOOS=linux go build -o zinx-takes-history cmd/services/take-history/main.go

FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build-stage /app/zinx /app/zinx
COPY --from=build-stage /app/public /app/public
COPY --from=build-stage /app/zinx-takes-history /app/zinx-takes-history
EXPOSE 7777
ENV ZINX_STATIC_PATH /app/public
USER nonroot:nonroot
CMD ["/app/zinx"]
