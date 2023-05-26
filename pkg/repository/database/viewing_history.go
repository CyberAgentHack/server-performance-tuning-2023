package database

import (
	"context"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/aws/aws-xray-sdk-go/xray"
)

type ViewingHistory struct {
}

func NewViewingHistory() *ViewingHistory {
	return &ViewingHistory{}
}

func (e *ViewingHistory) Create(ctx context.Context, viewingHistory *entity.ViewingHistory) (*entity.ViewingHistory, error) {
	_, seg := xray.BeginSubsegment(ctx, "database.ViewingHistory#Create")
	defer seg.Close(nil)
	// TODO
	return &entity.ViewingHistory{ID: "id"}, nil
}

func (e *ViewingHistory) BatchGet(ctx context.Context, ids []string, userID string) (entity.ViewingHistories, error) {
	_, seg := xray.BeginSubsegment(ctx, "database.ViewingHistory#BatchGet")
	defer seg.Close(nil)
	// TODO
	return entity.ViewingHistories{{ID: "id"}}, nil
}
