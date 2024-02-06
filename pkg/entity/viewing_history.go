package entity

type ViewingHistory struct {
	UserID        string `json:"userID"`
	EpisodeID     string `json:"episodeID"`
	SeriesID      string `json:"seriesID"`
	SeasonID      string `json:"seasonID"`
	IsWatched     bool   `json:"isWatched"`     // 視聴完了済み判定フラグ
	Position      int64  `json:"position"`      // どの時点まで番組を見たか
	Duration      int64  `json:"duration"`      // 視聴時間
	LastViewingAt int64  `json:"lastViewingAt"` // 最後の視聴日時
}

type ViewingHistories []*ViewingHistory
