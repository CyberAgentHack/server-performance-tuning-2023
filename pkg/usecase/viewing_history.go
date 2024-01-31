package usecase

import (
	"context"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
	"github.com/aws/aws-xray-sdk-go/xray"
)

type CreateViewingHistoryRequest struct {
	ViewingHistory *entity.ViewingHistory
}

type CreateViewingHistoryResponse struct {
	ViewingHistory *entity.ViewingHistory
}

func (c *CreateViewingHistoryRequest) validate() error {
	if c.ViewingHistory == nil {
		return errcode.NewInvalidArgument("usecase: Article is required")
	}
	return nil
}

func (u *UsecaseImpl) CreateViewingHistory(ctx context.Context, req *CreateViewingHistoryRequest) (*CreateViewingHistoryResponse, error) {
	ctx, seg := xray.BeginSubsegment(ctx, "usecase.UsecaseImpl#CreateViewingHistory")
	defer seg.Close(nil)

	if err := req.validate(); err != nil {
		return nil, errcode.New(err)
	}

	viewingHistory, err := u.db.ViewingHistory.Create(ctx, req.ViewingHistory)
	if err != nil {
		return nil, errcode.New(err)
	}
	return &CreateViewingHistoryResponse{ViewingHistory: viewingHistory}, nil
}

type BatchGetViewingHistoriesRequest struct {
	UserID     string   `validate:"required"`
	EpisodeIDs []string `validate:"required"`
}

type BatchGetViewingHistoriesResponse struct {
	ViewingHistories entity.ViewingHistories
}

func (u *UsecaseImpl) BatchGetViewingHistories(ctx context.Context, req *BatchGetViewingHistoriesRequest) (*BatchGetViewingHistoriesResponse, error) {
	ctx, seg := xray.BeginSubsegment(ctx, "usecase.UsecaseImpl#BatchGetViewingHistories")
	defer seg.Close(nil)

	if err := u.validate.Struct(req); err != nil {
		return nil, errcode.New(err)
	}

	viewingHistories := make(entity.ViewingHistories, 0, len(req.EpisodeIDs))

	for _, episodeID := range req.EpisodeIDs {
		viewingHistory, err := u.db.ViewingHistory.Get(ctx, episodeID, req.UserID)
		if err != nil {
			return nil, errcode.New(err)
		}
		viewingHistories = append(viewingHistories, viewingHistory)
	}

	return &BatchGetViewingHistoriesResponse{ViewingHistories: viewingHistories}, nil
}
