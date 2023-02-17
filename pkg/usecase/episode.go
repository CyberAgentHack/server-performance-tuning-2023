package usecase

import (
	"context"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

type ListEpisodesRequest struct {
	Limit    int
	Offset   int
	SeasonID string
	SeriesID string
}

type ListEpisodesResponse struct {
	Episodes entity.Episodes
	Series   entity.SeriesMulti
	Seasons  entity.Seasons
}

func (u *UsecaseImpl) ListEpisodes(ctx context.Context, req *ListEpisodesRequest) (*ListEpisodesResponse, error) {
	ctx, span := tracer.Start(ctx, "usecase.UsecaseImpl#ListEpisodes")
	defer span.End()

	params := &repository.ListEpisodesParams{
		Limit:    req.Limit,
		Offset:   req.Offset,
		SeasonID: req.SeasonID,
	}
	episodes, err := u.db.Episode.List(ctx, params)
	if err != nil {
		return nil, errcode.New(err)
	}

	srids := make([]string, 0, len(episodes))
	for i := range episodes {
		srids = append(srids, episodes[i].SeriesID)
	}
	series, err := u.db.Series.BatchGet(ctx, srids)
	if err != nil {
		return nil, errcode.New(err)
	}

	seids := make([]string, 0, len(episodes))
	for i := range episodes {
		if episodes[i].SeasonID == nil {
			continue
		}
		seids = append(seids, *episodes[i].SeasonID)
	}
	seasons, err := u.db.Season.BatchGet(ctx, seids)
	if err != nil {
		return nil, errcode.New(err)
	}

	return &ListEpisodesResponse{
		Episodes: episodes,
		Series:   series,
		Seasons:  seasons,
	}, nil
}
