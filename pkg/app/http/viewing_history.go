package http

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/usecase"
)

func (s *Service) routeViewingHistory(r chi.Router) {
	r.Get("/", s.listViewingHistories)
}

func (s *Service) listViewingHistories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	episodeIDs := QueryStrings(r, "episodeIds")
	if len(episodeIDs) == 0 {
		s.OK(nil, w, r)
		return
	}
	req := &usecase.BatchGetViewingHistoriesRequest{
		UserID:     r.Header.Get("userId"),
		EpisodeIDs: QueryStrings(r, "episodeIds"),
	}
	resp, err := s.usecase.BatchGetViewingHistories(ctx, req)
	if err != nil {
		s.Error(err, w, r)
		return
	}

	s.OK(resp.ViewingHistories, w, r)
}
