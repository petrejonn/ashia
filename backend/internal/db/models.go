// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Category struct {
	CategoryID         int64
	Slug               string
	Title              string
	Description        pgtype.Text
	ParentID           pgtype.Int8
	CreatedAt          pgtype.Timestamptz
	UpdatedAt          pgtype.Timestamptz
	ShopID             int64
	CategoryAttributes []byte
}

type Facebook struct {
	FacebookID int64
	Handle     string
	Url        string
	ShopID     int64
}

type Product struct {
	ProductID         int64
	Title             string
	Description       string
	AllowedAttributes []byte
	CreatedAt         pgtype.Timestamptz
	UpdatedAt         pgtype.Timestamptz
	Status            string
	CategoryID        int64
	ShopID            int64
}

type ProductImage struct {
	ProductImageID int64
	Url            string
	Alt            string
	ProductID      int64
	ShopID         int64
}

type Shop struct {
	ShopID         int64
	OwnerID        uuid.UUID
	Title          string
	Domain         string
	Email          string
	CurrencyCode   string
	Status         string
	About          pgtype.Text
	Address        pgtype.Text
	PhoneNumber    pgtype.Text
	SeoDescription pgtype.Text
	SeoKeywords    []string
	SeoTitle       pgtype.Text
	UpdatedAt      pgtype.Timestamptz
	CreatedAt      pgtype.Timestamptz
}

type ShopImage struct {
	ShopImageID   int64
	FaviconUrl    pgtype.Text
	LogoUrl       pgtype.Text
	BannerUrl     pgtype.Text
	CoverImageUrl pgtype.Text
	ShopID        int64
}

type User struct {
	UserID            uuid.UUID
	Auth0Sub          pgtype.Text
	Email             string
	Name              pgtype.Text
	ProfilePictureUrl pgtype.Text
	CreatedAt         pgtype.Timestamp
	LastLogin         pgtype.Timestamp
}

type Whatsapp struct {
	WhatsappID  int64
	PhoneNumber string
	Url         string
	ShopID      int64
}
