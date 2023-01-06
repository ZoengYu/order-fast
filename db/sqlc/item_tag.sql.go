// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: item_tag.sql

package db

import (
	"context"
)

const createMenuItemTag = `-- name: CreateMenuItemTag :one
INSERT INTO item_tag (
	item_id,
	item_tag
) VALUES (
	$1, $2
) RETURNING id, item_id, item_tag
`

type CreateMenuItemTagParams struct {
	ItemID  int64  `json:"item_id"`
	ItemTag string `json:"item_tag"`
}

func (q *Queries) CreateMenuItemTag(ctx context.Context, arg CreateMenuItemTagParams) (ItemTag, error) {
	row := q.db.QueryRowContext(ctx, createMenuItemTag, arg.ItemID, arg.ItemTag)
	var i ItemTag
	err := row.Scan(&i.ID, &i.ItemID, &i.ItemTag)
	return i, err
}

const getMenuItemTag = `-- name: GetMenuItemTag :one
SELECT id, item_id, item_tag FROM item_tag
WHERE item_id = $1 AND item_tag = $2
`

type GetMenuItemTagParams struct {
	ItemID  int64  `json:"item_id"`
	ItemTag string `json:"item_tag"`
}

func (q *Queries) GetMenuItemTag(ctx context.Context, arg GetMenuItemTagParams) (ItemTag, error) {
	row := q.db.QueryRowContext(ctx, getMenuItemTag, arg.ItemID, arg.ItemTag)
	var i ItemTag
	err := row.Scan(&i.ID, &i.ItemID, &i.ItemTag)
	return i, err
}

const listMenuItemTag = `-- name: ListMenuItemTag :many
SELECT item_tag.item_tag FROM item_tag
WHERE item_id = $1
`

func (q *Queries) ListMenuItemTag(ctx context.Context, itemID int64) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listMenuItemTag, itemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var item_tag string
		if err := rows.Scan(&item_tag); err != nil {
			return nil, err
		}
		items = append(items, item_tag)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeMenuItemTag = `-- name: RemoveMenuItemTag :exec
DELETE FROM item_tag
WHERE item_id = $1 AND item_tag = $2
`

type RemoveMenuItemTagParams struct {
	ItemID  int64  `json:"item_id"`
	ItemTag string `json:"item_tag"`
}

func (q *Queries) RemoveMenuItemTag(ctx context.Context, arg RemoveMenuItemTagParams) error {
	_, err := q.db.ExecContext(ctx, removeMenuItemTag, arg.ItemID, arg.ItemTag)
	return err
}