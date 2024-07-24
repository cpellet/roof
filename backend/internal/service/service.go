package service

import (
	"backend/internal/analysis"
	roof "backend/internal/proto"
	"backend/pkg/logging"
	"context"
)

// RoofService implements the interface generated from grpc.
// UnimplementedRoofServiceProceduresServer must be embedded
// to have forward compatible implementations.
type RoofService struct {
	roof.UnimplementedRoofServiceProceduresServer
	processor *analysis.Processor
}

// NewRoofService creates a new RoofService.
func NewRoofService(p *analysis.Processor) RoofService {
	return RoofService{processor: p}
}

// Ping is a simple ping procedure.
func (r *RoofService) Ping(ctx context.Context, req *roof.PingMessage) (*roof.PingMessage, error) {
	logger := logging.FromContext(ctx)
	logger.Infow("received ping", "message", req.Message)
	return &roof.PingMessage{Message: req.Message}, nil
}

// PerformAnalysis performs an analysis on the cmap and emap.
func (r *RoofService) PerformAnalysis(ctx context.Context, req *roof.PerformAnalysisRequest) (*roof.
	PerformAnalysisResponse,
	error) {
	logger := logging.FromContext(ctx)
	_, id, err := r.processor.ProcessRequest(ctx, req.Cmap, req.Emap)
	if err != nil {
		logger.Errorw("failed to process request", "error", err)
		return nil, err
	}
	logger.Infow("analysis complete", "mask_id", id)
	return &roof.PerformAnalysisResponse{
		Id: id,
	}, nil
}

// RetrieveAnalysis retrieves the cmap and msmap for the given id.
func (r *RoofService) RetrieveAnalysis(ctx context.Context, req *roof.RetrieveAnalysisRequest) (*roof.
	RetrieveAnalysisResponse,
	error) {
	logger := logging.FromContext(ctx)
	id := req.Id
	cmap, err := r.processor.Blobstore.GetObject(ctx, id, analysis.BlobKeyImage)
	if err != nil {
		logger.Errorw("failed to retrieve cmap", "error", err)
		return nil, err
	}
	msmap, err := r.processor.Blobstore.GetObject(ctx, id, analysis.BlobKeyMask)
	if err != nil {
		logger.Errorw("failed to retrieve msmap", "error", err)
		return nil, err
	}
	return &roof.RetrieveAnalysisResponse{
		Cmap:  cmap,
		Msmap: msmap,
	}, nil
}
