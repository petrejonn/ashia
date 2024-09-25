package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.54

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/gosimple/slug"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/petrejonn/naytife/internal/db"
	"github.com/petrejonn/naytife/internal/graph/generated"
	"github.com/petrejonn/naytife/internal/graph/model"
)

// ID is the resolver for the id field.
func (r *categoryResolver) ID(ctx context.Context, obj *model.Category) (string, error) {
	// Return the base64-encoded ID
	return EncodeRelayID("Category", obj.ID), nil
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, category model.CreateCategoryInput) (model.CreateCategoryPayload, error) {
	shopID := ctx.Value("shop_id").(int64)
	param := db.CreateShopCategoryParams{
		Title:       category.Title,
		Description: pgtype.Text{String: category.Description, Valid: true},
		Slug:        slug.MakeLang(category.Title, "en"),
		ShopID:      shopID,
	}
	param.CategoryAttributes = []byte("{}")
	if category.ParentID != nil {
		_, id, err := DecodeRelayID(*category.ParentID)
		if err != nil {
			return nil, errors.New("could not find parent category")
		}
		parent, err := r.Repository.GetShopCategory(ctx, *id)
		if err != nil {
			return nil, errors.New("could not find parent category")
		}
		param.ParentID = parent.ParentID
		param.CategoryAttributes = parent.CategoryAttributes
	}
	cat, err := r.Repository.CreateShopCategory(ctx, param)
	if err != nil {
		return nil, err
	}
	return &model.CreateCategorySuccess{
		Category: &model.Category{
			ID:                strconv.FormatInt(cat.CategoryID, 10),
			Slug:              cat.Slug,
			Title:             cat.Title,
			Description:       cat.Description.String,
			Parent:            &model.Category{},
			Children:          []model.Category{},
			Products:          nil,
			Images:            nil,
			AllowedAttributes: nil,
			CreatedAt:         cat.CreatedAt.Time,
			UpdatedAt:         cat.UpdatedAt.Time,
		}}, nil
}

// UpdateCategory is the resolver for the updateCategory field.
func (r *mutationResolver) UpdateCategory(ctx context.Context, categoryID string, category model.UpdateCategoryInput) (model.UpdateCategoryPayload, error) {
	_, catId, err := DecodeRelayID(categoryID)
	if err != nil {
		return nil, errors.New("invalid category ID")
	}
	params := db.UpdateShopCategoryParams{
		CategoryID:  *catId,
		Title:       pgTextFromStringPointer(category.Title),
		Description: pgTextFromStringPointer(category.Description),
	}

	if category.ParentID != nil {
		_, id, err := DecodeRelayID(*category.ParentID)
		if err != nil {
			return nil, errors.New("could not find parent category")
		}
		params.ParentID = pgtype.Int8{Int64: *id, Valid: true}
	}

	dbCat, err := r.Repository.UpdateShopCategory(ctx, params)
	if err != nil {
		return nil, errors.New("could not update category")
	}
	attributes, err := unmarshalCategoryAttributes(dbCat.CategoryAttributes)
	if err != nil {
		return nil, errors.New("could not unmarshal category attributes")
	}
	return &model.UpdateCategorySuccess{
		Category: &model.Category{
			ID:                strconv.FormatInt(dbCat.CategoryID, 10),
			Slug:              dbCat.Slug,
			Title:             dbCat.Title,
			Description:       dbCat.Description.String,
			Parent:            &model.Category{},
			Children:          []model.Category{},
			Products:          nil,
			Images:            nil,
			AllowedAttributes: attributes,
			CreatedAt:         dbCat.CreatedAt.Time,
			UpdatedAt:         dbCat.UpdatedAt.Time,
		}}, nil
}

// CreateCategoryAttribute is the resolver for the createCategoryAttribute field.
func (r *mutationResolver) CreateCategoryAttribute(ctx context.Context, categoryID string, attribute model.CreateCategoryAttributeInput) (model.CreateCategoryAttributePayload, error) {
	_, catId, err := DecodeRelayID(categoryID)
	if err != nil {
		return nil, errors.New("invalid category ID")
	}
	attributesDB, err := r.Repository.CreateCategoryAttribute(ctx, db.CreateCategoryAttributeParams{
		CategoryID: *catId,
		Title:      attribute.Title,
		DataType:   attribute.DataType.String(),
	})
	if err != nil {
		return nil, errors.New("could not create category attribute")
	}
	// Unmarshal JSONB ([]byte) into a Go map
	var attributesMap map[string]interface{}
	if err := json.Unmarshal(attributesDB, &attributesMap); err != nil {
		return nil, errors.New("could not create category attribute")
	}
	attributes := make([]model.AllowedCategoryAttributes, 0, len(attributesMap))

	// Iterate over the map and populate the attribute list
	for title, dataType := range attributesMap {
		attributes = append(attributes, model.AllowedCategoryAttributes{
			Title:    title,                                             // The map key is the title (string)
			DataType: model.ProductAttributeDataType(dataType.(string)), // The map value is the data type
		})
	}
	return &model.CreateCategoryAttributeSuccess{
		Attributes: attributes,
	}, nil
}

// DeleteCategoryAttribute is the resolver for the deleteCategoryAttribute field.
func (r *mutationResolver) DeleteCategoryAttribute(ctx context.Context, categoryID string, attribute string) (model.DeleteCategoryAttributePayload, error) {
	_, catId, err := DecodeRelayID(categoryID)
	if err != nil {
		return nil, errors.New("invalid category ID")
	}
	attributesDB, err := r.Repository.DeleteCategoryAttribute(ctx, db.DeleteCategoryAttributeParams{
		CategoryID: *catId, Attribute: attribute,
	})
	if err != nil {
		return nil, errors.New("could not fetch category attribute")
	}
	// Unmarshal JSONB ([]byte) into a Go map
	var attributesMap map[string]interface{}
	if err := json.Unmarshal(attributesDB, &attributesMap); err != nil {
		return nil, errors.New("could not fetch category attribute")
	}
	attributes := make([]model.AllowedCategoryAttributes, 0, len(attributesMap))

	// Iterate over the map and populate the attribute list
	for title, dataType := range attributesMap {
		attributes = append(attributes, model.AllowedCategoryAttributes{
			Title:    title,                                             // The map key is the title (string)
			DataType: model.ProductAttributeDataType(dataType.(string)), // The map value is the data type
		})
	}
	return &model.DeleteCategoryAttributeSuccess{
		Attributes: attributes,
	}, nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]model.Category, error) {
	categoriesDB, err := r.Repository.GetShopCategories(ctx)
	if err != nil {
		return nil, errors.New("could not fetch categories")
	}

	categories := make([]model.Category, 0, len(categoriesDB))
	for _, cat := range categoriesDB {
		attributes, err := unmarshalCategoryAttributes(cat.CategoryAttributes)
		if err != nil {
			return nil, errors.New("could not understand category")
		}
		categories = append(categories, model.Category{
			ID:                strconv.FormatInt(cat.CategoryID, 10),
			Slug:              cat.Slug,
			Title:             cat.Title,
			Description:       cat.Description.String,
			Parent:            &model.Category{},
			Children:          []model.Category{},
			Products:          nil,
			Images:            nil,
			AllowedAttributes: attributes,
			CreatedAt:         cat.CreatedAt.Time,
			UpdatedAt:         cat.UpdatedAt.Time,
		})
	}
	return categories, nil
}

// Category is the resolver for the category field.
func (r *queryResolver) Category(ctx context.Context, id string) (*model.Category, error) {
	_, catId, err := DecodeRelayID(id)
	if err != nil {
		return nil, errors.New("invalid category ID")
	}
	cat, err := r.Repository.GetShopCategory(ctx, *catId)
	if err != nil {
		return nil, errors.New("could not find category")
	}
	attributes, err := unmarshalCategoryAttributes(cat.CategoryAttributes)
	if err != nil {
		return nil, errors.New("could not understand category")
	}
	return &model.Category{
		ID:                strconv.FormatInt(cat.CategoryID, 10),
		Slug:              cat.Slug,
		Title:             cat.Title,
		Description:       cat.Description.String,
		Parent:            &model.Category{},
		Children:          []model.Category{},
		Products:          nil,
		Images:            nil,
		AllowedAttributes: attributes,
		CreatedAt:         cat.CreatedAt.Time,
		UpdatedAt:         cat.UpdatedAt.Time,
	}, nil
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

type categoryResolver struct{ *Resolver }
