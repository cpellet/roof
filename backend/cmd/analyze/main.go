package main

import (
	"backend/internal/analysis"
	blob "backend/pkg/blobstore"
	"backend/pkg/logging"
	"context"
	"os"
	"os/signal"
	"syscall"
)

// Analyze does a one-shot analysis of the image and elevation map files, writing the mask to the specified blobstore.

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
	imgFilename := os.Args[1]
	emapFilename := os.Args[2]

	logger.Infow("starting analysis")
	eBytes, err := analysis.ReadFileBytes(emapFilename)
	if err != nil {
		return err
	}
	imgBytes, err := analysis.ReadFileBytes(imgFilename)
	if err != nil {
		return err
	}

	blobstore, err := blob.BlobstoreFor(ctx, &blob.Config{
		Type:    os.Getenv("BLOBSTORE_TYPE"),
		RootDir: os.Getenv("BLOBSTORE_ROOT_DIR"),
	})
	if err != nil {
		return err
	}

	processor := &analysis.Processor{Blobstore: blobstore}
	_, id, err := processor.ProcessRequest(ctx, imgBytes, eBytes)
	if err != nil {
		return err
	}

	defer logger.Infof("analysis complete, mask written at %s", id)
	return nil
}
