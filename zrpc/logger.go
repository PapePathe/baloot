package zrpc

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func GrpcLogger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	logger := log.Info()

	peer, ok := peer.FromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "missing peer info")
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
	}

	statusCode := codes.Unknown
	startTime := time.Now()
	result, err := handler(ctx, req)
	duration := time.Since(startTime)

	st, ok := status.FromError(err)
	if ok {
		statusCode = st.Code()
	}

	if err != nil {

		switch st.Code() {
		case codes.OK:
			logger = log.Info()

		case codes.Canceled,
			codes.InvalidArgument,
			codes.NotFound,
			codes.AlreadyExists,
			codes.PermissionDenied,
			codes.FailedPrecondition,
			codes.Aborted,
			codes.OutOfRange,
			codes.Unimplemented,
			codes.Unauthenticated:
			logger = log.Warn().Err(err)

		default:
			logger = log.Error().Err(err)
		}
	}

	logger.
		Int("StatusCode", int(statusCode)).
		Str("StatusText", statusCode.String()).
		Str("Latency", duration.String()).
		Str("Method", info.FullMethod).
		Str("Peer", peer.Addr.String()).
		Interface("UserAgent", md["user-agent"]).
		Msg("")

	return result, err
}
