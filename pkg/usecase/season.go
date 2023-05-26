package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
	"github.com/aws/aws-xray-sdk-go/xray"
)

type ListSeasonsRequest struct {
	Limit    int
	Offset   int
	SeriesID string
}

type ListSeasonsResponse struct {
	Seasons entity.Seasons
}

func (u *UsecaseImpl) ListSeasons(ctx context.Context, req *ListSeasonsRequest) (*ListSeasonsResponse, error) {
	ctx, seg := xray.BeginSubsegment(ctx, "usecase.UsecaseImpl#ListSeasons")
	defer seg.Close(nil)

	key := fmt.Sprintf("%v", req)
	resp := &ListSeasonsResponse{}
	hit, err := u.redis.Get(ctx, key, resp)
	if err != nil {
		return nil, errcode.New(err)
	}
	if hit {
		return resp, nil
	}

	params := &repository.ListSeasonsParams{
		Limit:    req.Limit,
		Offset:   req.Offset,
		SeriesID: req.SeriesID,
	}
	seasons, err := u.db.Season.List(ctx, params)
	if err != nil {
		return nil, errcode.New(err)
	}

	resp = &ListSeasonsResponse{
		Seasons: seasons,
	}
	u.redis.Set(ctx, key, resp, time.Second*10)
	return resp, nil
}
