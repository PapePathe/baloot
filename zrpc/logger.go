package zrpc

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcLogger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	logger := log.Info()

	statusCode := codes.Unknown
	startTime := time.Now()
	result, err := handler(ctx, req)
	duration := time.Since(startTime)

	if st, ok := status.FromError(err); ok {
		statusCode = st.Code()
	}

	if err != nil {
		logger = log.Error().Err(err)
	}

	logger.
		Str("Method", info.FullMethod).
		Int("StatusCode", int(statusCode)).
		Str("StatusText", statusCode.String()).
		Dur("Latency", duration).
		Msg("")

	return result, err
}
