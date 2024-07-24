package main

import (
	"backend/internal/analysis"
	roof "backend/internal/proto"
	"backend/internal/service"
	blob "backend/pkg/blobstore"
	"backend/pkg/logging"
	"context"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// Serve spawns an analysis gRPC server that listens on the specified port.

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	logger := logging.NewLoggerFromEnv()
	ctx = logging.WithLogger(ctx, logger)

	defer func() {
		done()
		if r := recover(); r != nil {
			logger.Fatalw("application panic", "panic", r)
		}
	}()

	err := realMain(ctx)

	done()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("successful shutdown")
}

func realMain(ctx context.Context) error {
	logger := logging.FromContext(ctx)
	logger.Infow("starting server")

	blobstore, err := blob.BlobstoreFor(ctx, &blob.Config{
		Type:    os.Getenv("BLOBSTORE_TYPE"),
		RootDir: os.Getenv("BLOBSTORE_ROOT_DIR"),
	})
	if err != nil {
		return err
	}

	proc := &analysis.Processor{
		Blobstore: blobstore,
	}
	svc := service.NewRoofService(proc)

	port := os.Getenv("PORT")
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	roof.RegisterRoofServiceProceduresServer(server, &svc)

	logger.Info("listening on ", listener.Addr())
	err = server.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}
