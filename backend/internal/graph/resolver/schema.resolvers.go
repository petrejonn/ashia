package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.54

import (
	"context"
	"errors"

	"github.com/petrejonn/naytife/internal/db"
	"github.com/petrejonn/naytife/internal/graph/generated"
	"github.com/petrejonn/naytife/internal/graph/model"
)

// SignInUser is the resolver for the signInUser field.
func (r *mutationResolver) SignInUser(ctx context.Context, input model.SignInInput) (model.SignInUserPayload, error) {
	fakeAuthSub := "9vgPO5K5ipI424xe84HUrtqQJMWT3e7f@clients"
	fakeEmail := "fake@email.com"
	fakeName := "Fake Name"
	fakeUrl := "https://fake-url.com/profile-picture.png"
	// Check if the user exists in the database
	user, err := r.Repository.UpsertUser(ctx, db.UpsertUserParams{
		Auth0Sub:          &fakeAuthSub,
		Email:             fakeEmail,
		Name:              &fakeName,
		ProfilePictureUrl: &fakeUrl,
	})
	if err != nil {
		return nil, errors.New("user not found")
	}
	// Return the user data and a JWT token
	return &model.SignInUserSuccess{
		User: &model.User{
			ID:                user.UserID.String(),
			Email:             user.Email,
			Name:              user.Name,
			ProfilePictureURL: user.ProfilePictureUrl,
		},
	}, nil
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	typ, _, err := decodeRelayID(id)
	if err != nil {
		return nil, err
	}

	switch typ {
	case "Shop":
		return r.Shop(ctx)
	case "Category":
		return r.Category(ctx, id)
	default:
		return nil, errors.New("unknown node type")
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
