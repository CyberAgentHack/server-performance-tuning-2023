package http

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/usecase"
)

func (s *Service) routeEpisode(r chi.Router) {
	r.Get("/", s.listEpisodes)
}

func (s *Service) listEpisodes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &usecase.ListEpisodesRequest{
		Limit:    QueryIntDefault(r, "limit", 20),
		Offset:   QueryInt(r, "offset"),
		SeasonID: Query(r, "seasonId"),
		SeriesID: Query(r, "seriesId"),
	}
	resp, err := s.usecase.ListEpisodes(ctx, req)
	if err != nil {
		s.Error(err, w, r)
		return
	}
	s.OK(&entity.ListEpisodesResponse{Episodes: resp.Episodes, Series: resp.Series, Seasons: resp.Seasons}, w, r)
}
