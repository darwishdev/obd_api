// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: center.sql

package db

import (
	"context"
)

const centerCreate = `-- name: CenterCreate :one
INSERT INTO
    centers (
        name,
        phone,
        location,
        address,
        area_id,
        lat,
        long
    )
VALUES
    ($1, $2, $3, $4, $5, $6, $7) RETURNING center_id, name, phone, location, address, area_id, lat, long, created_at, deleted_at
`

type CenterCreateParams struct {
	Name     string  `json:"name"`
	Phone    string  `json:"phone"`
	Location string  `json:"location"`
	Address  string  `json:"address"`
	AreaID   int64   `json:"area_id"`
	Lat      float32 `json:"lat"`
	Long     float32 `json:"long"`
}

func (q *Queries) CenterCreate(ctx context.Context, arg CenterCreateParams) (Center, error) {
	row := q.db.QueryRowContext(ctx, centerCreate,
		arg.Name,
		arg.Phone,
		arg.Location,
		arg.Address,
		arg.AreaID,
		arg.Lat,
		arg.Long,
	)
	var i Center
	err := row.Scan(
		&i.CenterID,
		&i.Name,
		&i.Phone,
		&i.Location,
		&i.Address,
		&i.AreaID,
		&i.Lat,
		&i.Long,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const centersList = `-- name: CentersList :many
SELECT center_id, name, phone, location, address, area_id, lat, long, created_at, avg_rate, reviews_count, distance FROM 
    find_centers(
        $1,
        $2
    )
`

type CentersListParams struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

func (q *Queries) CentersList(ctx context.Context, arg CentersListParams) ([]CenterInfo, error) {
	rows, err := q.db.QueryContext(ctx, centersList, arg.Lat, arg.Long)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CenterInfo{}
	for rows.Next() {
		var i CenterInfo
		if err := rows.Scan(
			&i.CenterID,
			&i.Name,
			&i.Phone,
			&i.Location,
			&i.Address,
			&i.AreaID,
			&i.Lat,
			&i.Long,
			&i.CreatedAt,
			&i.AvgRate,
			&i.ReviewsCount,
			&i.Distance,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
