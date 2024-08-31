package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"github.com/petrejonn/ashia/graph/model"
)

func strPtr(s string) *string {
	return &s
}

// CreateShop is the resolver for the createShop field.
func (r *mutationResolver) CreateShop(ctx context.Context, shop model.CreateShopInput) (*model.Shop, error) {
	panic(fmt.Errorf("not implemented: CreateShop - createShop"))
}

// Shop is the resolver for the shop field.
func (r *queryResolver) Shop(ctx context.Context, id string) (*model.Shop, error) {
	shop := &model.Shop{
		ID:     id,
		Title:  faker.Word(options.WithRandomStringLength(10)),
		Domain: faker.DomainName(),
		ContactPhone: &model.PhoneNumber{
			Number:      faker.Phonenumber(),
			CountryCode: "+234",
		},
		ContactEmail: strPtr(faker.Email()),
		Location: &model.Location{
			Address: faker.GetRealAddress().Address,
			State:   faker.GetRealAddress().State,
			Country: faker.GetRealAddress().State,
		},
		// Products     *ProductConnection
		WhatsApp: &model.WhatsApp{
			URL: faker.URL(),
			Number: &model.PhoneNumber{
				Number:      faker.Phonenumber(),
				CountryCode: "+234",
			},
		},
		Facebook: &model.Facebook{
			URL:    faker.URL(),
			Handle: faker.Word(),
		},
		SiteLogoURL: strPtr(faker.URL()),
		FaviconURL:  strPtr(faker.URL()),
		Currency:    strPtr(faker.Currency()),
		Status:      (*model.ShopStatus)(strPtr(model.ShopStatusPublished.String())),
		About:       strPtr(faker.Paragraph()),
		UpdatedAt:   strPtr(faker.Timestamp()),
		CreatedAt:   strPtr(faker.Timestamp()),
	}
	return shop, nil
}

// MyShops is the resolver for the myShops field.
func (r *queryResolver) MyShops(ctx context.Context) ([]*model.Shop, error) {
	shops := make([]*model.Shop, 5) // Generate 5 fake users
	for i := range shops {
		shops[i] = &model.Shop{
			ID:     faker.UUIDHyphenated(),
			Title:  faker.LastName(),
			Domain: faker.DomainName(),
			ContactPhone: &model.PhoneNumber{
				Number:      faker.Phonenumber(),
				CountryCode: "+234",
			},
			ContactEmail: strPtr(faker.Email()),
			Location: &model.Location{
				Address: faker.GetRealAddress().Address,
				State:   faker.GetRealAddress().State,
				Country: faker.GetRealAddress().State,
			},
			// Products     *ProductConnection
			WhatsApp: &model.WhatsApp{
				URL: faker.URL(),
				Number: &model.PhoneNumber{
					Number:      faker.Phonenumber(),
					CountryCode: "+234",
				},
			},
			Facebook: &model.Facebook{
				URL:    faker.URL(),
				Handle: faker.Word(),
			},
			SiteLogoURL: strPtr(faker.URL()),
			FaviconURL:  strPtr(faker.URL()),
			Currency:    strPtr(faker.Currency()),
			Status:      (*model.ShopStatus)(strPtr(model.ShopStatusPublished.String())),
			About:       strPtr(faker.Paragraph()),
			UpdatedAt:   strPtr(faker.Timestamp()),
			CreatedAt:   strPtr(faker.Timestamp()),
		}
	}
	return shops, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
