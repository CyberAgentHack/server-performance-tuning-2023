package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
	"github.com/aws/aws-xray-sdk-go/xray"
	"go.uber.org/multierr"
)

type Season struct {
	db *sql.DB
}

func NewSeason(db *sql.DB) *Season {
	return &Season{db: db}
}

func (e *Season) List(ctx context.Context, params *repository.ListSeasonsParams) (entity.Seasons, error) {
	ctx, seg := xray.BeginSubsegment(ctx, "database.Season#List")
	defer seg.Close(nil)

	fields := []string{
		"seasonID",
		"seriesID",
		"displayName",
		"imageURL",
		"displayOrder",
	}

	clauses := make([]string, 0, 3)
	args := make([]any, 0, 3)
	if params.SeriesID != "" {
		clauses = append(clauses, "seriesID = ?")
		args = append(args, params.SeriesID)
	}
	if params.SeasonID != "" {
		clauses = append(clauses, "seasonID = ?")
		args = append(args, params.SeasonID)
	}

	var whereClause string
	if len(clauses) != 0 {
		whereClause = fmt.Sprintf("WHERE %s", strings.Join(clauses, " AND "))
	}

	query := fmt.Sprintf(
		"SELECT %s FROM seasons %s ORDER BY displayOrder LIMIT %d OFFSET %d",
		strings.Join(fields, ", "),
		whereClause,
		params.Limit,
		params.Offset,
	)

	rows, err := e.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errcode.New(err)
	}

	var seasons entity.Seasons
	for rows.Next() {
		season := &entity.Season{}
		err = rows.Scan(
			&season.ID,
			&season.SeriesID,
			&season.DisplayName,
			&season.ImageURL,
			&season.DisplayOrder,
		)
		if err != nil {
			break
		}
		seasons = append(seasons, season)
	}

	if closeErr := rows.Close(); closeErr != nil {
		return nil, errcode.New(closeErr)
	}
	if err != nil {
		return nil, errcode.New(err)
	}
	return seasons, nil
}

func (e *Season) Get(ctx context.Context, id string) (*entity.Season, error) {
	ctx, seg := xray.BeginSubsegment(ctx, "database.Season#Get")
	defer seg.Close(nil)

	fields := []string{
		"seasonID",
		"seriesID",
		"displayName",
		"imageURL",
		"displayOrder",
	}

	query := fmt.Sprintf(
		"SELECT %s FROM seasons WHERE seasonID = ?",
		strings.Join(fields, ", "),
	)
	row := e.db.QueryRowContext(ctx, query, id)
	if err := row.Err(); err != nil {
		return nil, errcode.New(err)
	}

	season := &entity.Season{}
	err := row.Scan(
		&season.ID,
		&season.SeriesID,
		&season.DisplayName,
		&season.ImageURL,
		&season.DisplayOrder,
	)
	return season, errcode.New(err)
}

func (e *Season) BatchGet(ctx context.Context, ids []string) (entity.Seasons, error) {
	ctx, seg := xray.BeginSubsegment(ctx, "database.Season#BatchGet")
	defer seg.Close(nil)

	if len(ids) == 0 {
		return nil, nil
	}

	newIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		if id != "" {
			newIDs = append(newIDs, id)
		}
	}

	fields := []string{
		"seasonID",
		"seriesID",
		"displayName",
		"imageURL",
		"displayOrder",
	}

	query := fmt.Sprintf(
		"SELECT %s FROM seasons WHERE seasonID IN(?%s)",
		strings.Join(fields, ","),
		strings.Repeat(",?", len(newIDs)-1),
	)
	rows, err := e.db.QueryContext(ctx, query, newIDs)
	if err != nil {
		return nil, errcode.New(err)
	}

	var seasons entity.Seasons
	var multiErr error
	for rows.Next() {
		var s entity.Season
		if err = rows.Scan(&s); err != nil {
			multiErr = multierr.Append(multiErr, err)
			continue
		}
		seasons = append(seasons, &s)
	}

	if cerr := rows.Close(); cerr != nil {
		return nil, errcode.New(cerr)
	}

	if multiErr != nil {
		return nil, errcode.New(multiErr)
	}

	return seasons, nil
}
