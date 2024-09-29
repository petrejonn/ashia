// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: product.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products ( title, description, category_id, shop_id, allowed_attributes, status)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING product_id, title, description, allowed_attributes, created_at, updated_at, status, category_id, shop_id
`

type CreateProductParams struct {
	Title             string
	Description       string
	CategoryID        int64
	ShopID            int64
	AllowedAttributes []byte
	Status            string
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProduct,
		arg.Title,
		arg.Description,
		arg.CategoryID,
		arg.ShopID,
		arg.AllowedAttributes,
		arg.Status,
	)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.Title,
		&i.Description,
		&i.AllowedAttributes,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Status,
		&i.CategoryID,
		&i.ShopID,
	)
	return i, err
}

const getProduct = `-- name: GetProduct :one
SELECT product_id, title, description, created_at, updated_at, status, category_id
FROM products
WHERE shop_id = $1 AND product_id = $2
`

type GetProductParams struct {
	ShopID    int64
	ProductID int64
}

type GetProductRow struct {
	ProductID   int64
	Title       string
	Description string
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
	Status      string
	CategoryID  int64
}

func (q *Queries) GetProduct(ctx context.Context, arg GetProductParams) (GetProductRow, error) {
	row := q.db.QueryRow(ctx, getProduct, arg.ShopID, arg.ProductID)
	var i GetProductRow
	err := row.Scan(
		&i.ProductID,
		&i.Title,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Status,
		&i.CategoryID,
	)
	return i, err
}

const getProducts = `-- name: GetProducts :many
SELECT product_id, title, description, created_at, updated_at, status, category_id
FROM products
WHERE shop_id = $1 AND product_id > $2
LIMIT $3
`

type GetProductsParams struct {
	ShopID int64
	After  int64
	Limit  int32
}

type GetProductsRow struct {
	ProductID   int64
	Title       string
	Description string
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
	Status      string
	CategoryID  int64
}

func (q *Queries) GetProducts(ctx context.Context, arg GetProductsParams) ([]GetProductsRow, error) {
	rows, err := q.db.Query(ctx, getProducts, arg.ShopID, arg.After, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProductsRow
	for rows.Next() {
		var i GetProductsRow
		if err := rows.Scan(
			&i.ProductID,
			&i.Title,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Status,
			&i.CategoryID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductsByCategory = `-- name: GetProductsByCategory :many
SELECT product_id, title, description, created_at, updated_at, status, category_id
FROM products
WHERE category_id = $1 AND product_id > $2
LIMIT $3
`

type GetProductsByCategoryParams struct {
	CategoryID int64
	After      int64
	Limit      int32
}

type GetProductsByCategoryRow struct {
	ProductID   int64
	Title       string
	Description string
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
	Status      string
	CategoryID  int64
}

func (q *Queries) GetProductsByCategory(ctx context.Context, arg GetProductsByCategoryParams) ([]GetProductsByCategoryRow, error) {
	rows, err := q.db.Query(ctx, getProductsByCategory, arg.CategoryID, arg.After, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProductsByCategoryRow
	for rows.Next() {
		var i GetProductsByCategoryRow
		if err := rows.Scan(
			&i.ProductID,
			&i.Title,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Status,
			&i.CategoryID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
