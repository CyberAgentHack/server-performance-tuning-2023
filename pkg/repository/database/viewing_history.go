package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
	"github.com/aws/aws-xray-sdk-go/xray"
)

type ViewingHistory struct {
	db *sql.DB
}

func NewViewingHistory(db *sql.DB) *ViewingHistory {
	return &ViewingHistory{db: db}
}

func (e *ViewingHistory) Get(ctx context.Context, id string, userID string) (*entity.ViewingHistory, error) {
	_, seg := xray.BeginSubsegment(ctx, "database.ViewingHistory#Get")
	defer seg.Close(nil)

	query := `SELECT userID, episodeID, seriesID, seasonID, isWatched, position, duration, lastViewingAt FROM viewingHistories WHERE userID = ? AND episodeID = ?`
	row := e.db.QueryRowContext(ctx, query, userID, id)
	if err := row.Err(); err != nil {
		return nil, errcode.New(err)
	}

	viewingHistory := &entity.ViewingHistory{}
	err := row.Scan(
		&viewingHistory.UserID,
		&viewingHistory.EpisodeID,
		&viewingHistory.SeriesID,
		&viewingHistory.SeasonID,
		&viewingHistory.IsWatched,
		&viewingHistory.Position,
		&viewingHistory.Duration,
		&viewingHistory.LastViewingAt,
	)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, errcode.New(err)
	}

	return viewingHistory, nil
}

func (e *ViewingHistory) BatchGet(ctx context.Context, ids []string, userID string) (entity.ViewingHistories, error) {
	_, seg := xray.BeginSubsegment(ctx, "database.ViewingHistory#BatchGet")
	defer seg.Close(nil)
	return nil, errcode.New(errcode.NewCode(errcode.CodeUnimplemented))
}
