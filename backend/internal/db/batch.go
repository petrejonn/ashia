// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: batch.go

package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrBatchAlreadyClosed = errors.New("batch already closed")
)

const deleteProductVariations = `-- name: DeleteProductVariations :batchexec
DELETE FROM product_variations
WHERE shop_id = $1 AND product_id = $2
AND product_variation_id != ALL($3::bigint[])
`

type DeleteProductVariationsBatchResults struct {
	br     pgx.BatchResults
	tot    int
	closed bool
}

type DeleteProductVariationsParams struct {
	ShopID              int64   `json:"shop_id"`
	ProductID           int64   `json:"product_id"`
	ProductVariationIds []int64 `json:"product_variation_ids"`
}

func (q *Queries) DeleteProductVariations(ctx context.Context, arg []DeleteProductVariationsParams) *DeleteProductVariationsBatchResults {
	batch := &pgx.Batch{}
	for _, a := range arg {
		vals := []interface{}{
			a.ShopID,
			a.ProductID,
			a.ProductVariationIds,
		}
		batch.Queue(deleteProductVariations, vals...)
	}
	br := q.db.SendBatch(ctx, batch)
	return &DeleteProductVariationsBatchResults{br, len(arg), false}
}

func (b *DeleteProductVariationsBatchResults) Exec(f func(int, error)) {
	defer b.br.Close()
	for t := 0; t < b.tot; t++ {
		if b.closed {
			if f != nil {
				f(t, ErrBatchAlreadyClosed)
			}
			continue
		}
		_, err := b.br.Exec()
		if f != nil {
			f(t, err)
		}
	}
}

func (b *DeleteProductVariationsBatchResults) Close() error {
	b.closed = true
	return b.br.Close()
}

const upsertProductVariation = `-- name: UpsertProductVariation :batchmany
INSERT INTO product_variations (slug, description, price, available_quantity, attributes, seo_description, seo_keywords, seo_title, product_id, shop_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
ON CONFLICT (slug, shop_id)
DO UPDATE SET
    description = EXCLUDED.description,
    price = EXCLUDED.price,
    available_quantity = EXCLUDED.available_quantity,
    attributes = EXCLUDED.attributes,
    seo_description = EXCLUDED.seo_description,
    seo_keywords = EXCLUDED.seo_keywords,
    seo_title = EXCLUDED.seo_title
RETURNING product_variation_id, slug, description, price, available_quantity, attributes, seo_description, seo_keywords, seo_title, created_at, updated_at, product_id, shop_id
`

type UpsertProductVariationBatchResults struct {
	br     pgx.BatchResults
	tot    int
	closed bool
}

type UpsertProductVariationParams struct {
	Slug              string         `json:"slug"`
	Description       string         `json:"description"`
	Price             pgtype.Numeric `json:"price"`
	AvailableQuantity int64          `json:"available_quantity"`
	Attributes        []byte         `json:"attributes"`
	SeoDescription    *string        `json:"seo_description"`
	SeoKeywords       []string       `json:"seo_keywords"`
	SeoTitle          *string        `json:"seo_title"`
	ProductID         int64          `json:"product_id"`
	ShopID            int64          `json:"shop_id"`
}

func (q *Queries) UpsertProductVariation(ctx context.Context, arg []UpsertProductVariationParams) *UpsertProductVariationBatchResults {
	batch := &pgx.Batch{}
	for _, a := range arg {
		vals := []interface{}{
			a.Slug,
			a.Description,
			a.Price,
			a.AvailableQuantity,
			a.Attributes,
			a.SeoDescription,
			a.SeoKeywords,
			a.SeoTitle,
			a.ProductID,
			a.ShopID,
		}
		batch.Queue(upsertProductVariation, vals...)
	}
	br := q.db.SendBatch(ctx, batch)
	return &UpsertProductVariationBatchResults{br, len(arg), false}
}

func (b *UpsertProductVariationBatchResults) Query(f func(int, []ProductVariation, error)) {
	defer b.br.Close()
	for t := 0; t < b.tot; t++ {
		var items []ProductVariation
		if b.closed {
			if f != nil {
				f(t, items, ErrBatchAlreadyClosed)
			}
			continue
		}
		err := func() error {
			rows, err := b.br.Query()
			if err != nil {
				return err
			}
			defer rows.Close()
			for rows.Next() {
				var i ProductVariation
				if err := rows.Scan(
					&i.ProductVariationID,
					&i.Slug,
					&i.Description,
					&i.Price,
					&i.AvailableQuantity,
					&i.Attributes,
					&i.SeoDescription,
					&i.SeoKeywords,
					&i.SeoTitle,
					&i.CreatedAt,
					&i.UpdatedAt,
					&i.ProductID,
					&i.ShopID,
				); err != nil {
					return err
				}
				items = append(items, i)
			}
			return rows.Err()
		}()
		if f != nil {
			f(t, items, err)
		}
	}
}

func (b *UpsertProductVariationBatchResults) Close() error {
	b.closed = true
	return b.br.Close()
}
