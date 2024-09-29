package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.54

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/petrejonn/naytife/internal/db"
	"github.com/petrejonn/naytife/internal/graph/generated"
	"github.com/petrejonn/naytife/internal/graph/model"
)

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, product model.CreateProductInput) (model.CreateProductPayload, error) {
	shopID := ctx.Value("shop_id").(int64)
	_, catID, err := DecodeRelayID(product.CategoryID)
	if err != nil {
		return nil, errors.New("invalid category ID")
	}
	cat, err := r.Repository.GetCategory(ctx, db.GetCategoryParams{ShopID: shopID, CategoryID: *catID})
	if err != nil {
		return &model.CategoryNotFoundError{Message: "category not found", Code: model.ErrorCodeNotFoundCategory}, nil
	}
	param := db.CreateProductParams{
		Title:             product.Title,
		Description:       product.Description,
		ShopID:            shopID,
		CategoryID:        *catID,
		AllowedAttributes: cat.CategoryAttributes,
		Status:            model.ProductStatusDraft.String(),
	}
	dbProduct, err := r.Repository.CreateProduct(ctx, param)
	if err != nil {
		log.Println(err)
		return nil, errors.New("server error")
	}
	attributes, err := unmarshalAllowedProductAttributes(dbProduct.AllowedAttributes)
	if err != nil {
		return nil, errors.New("could not understand category")
	}
	return model.CreateProductSuccess{Product: &model.Product{
		ID:                strconv.FormatInt(dbProduct.ProductID, 10),
		Title:             dbProduct.Title,
		Description:       dbProduct.Description,
		AllowedAttributes: attributes,
		Status:            (*model.ProductStatus)(&dbProduct.Status),
		CreatedAt:         dbProduct.CreatedAt.Time,
		UpdatedAt:         dbProduct.UpdatedAt.Time,
	}}, nil
}

// UpdateProduct is the resolver for the updateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, productID string, product model.UpdateProductInput) (model.UpdateProductPayload, error) {
	panic(fmt.Errorf("not implemented: UpdateProduct - updateProduct"))
}

// CreateProductAttribute is the resolver for the createProductAttribute field.
func (r *mutationResolver) CreateProductAttribute(ctx context.Context, productID string, attribute model.CreateProductAttributeInput) (model.CreateProductAttributePayload, error) {
	panic(fmt.Errorf("not implemented: CreateProductAttribute - createProductAttribute"))
}

// DeleteProductAttribute is the resolver for the deleteProductAttribute field.
func (r *mutationResolver) DeleteProductAttribute(ctx context.Context, productID string, attribute string) (model.DeleteCategoryAttributePayload, error) {
	panic(fmt.Errorf("not implemented: DeleteProductAttribute - deleteProductAttribute"))
}

// ID is the resolver for the id field.
func (r *productResolver) ID(ctx context.Context, obj *model.Product) (string, error) {
	return EncodeRelayID("Product", obj.ID), nil
}

// DefaultVariant is the resolver for the defaultVariant field.
func (r *productResolver) DefaultVariant(ctx context.Context, obj *model.Product) (*model.ProductVariant, error) {
	panic(fmt.Errorf("not implemented: DefaultVariant - defaultVariant"))
}

// Variants is the resolver for the variants field.
func (r *productResolver) Variants(ctx context.Context, obj *model.Product) ([]model.ProductVariant, error) {
	panic(fmt.Errorf("not implemented: Variants - variants"))
}

// AllowedAttributes is the resolver for the allowedAttributes field.
func (r *productResolver) AllowedAttributes(ctx context.Context, obj *model.Product) ([]model.AllowedProductAttributes, error) {
	panic(fmt.Errorf("not implemented: AllowedAttributes - allowedAttributes"))
}

// Images is the resolver for the images field.
func (r *productResolver) Images(ctx context.Context, obj *model.Product) ([]model.Image, error) {
	panic(fmt.Errorf("not implemented: Images - images"))
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context, first *int, after *string) (*model.ProductConnection, error) {
	// productsDB, err := r.Repository.GetProducts(ctx)
	// if err != nil {
	// 	log.Println(err)
	// 	return nil, errors.New("server error")
	// }
	// products := make([]model.Product, 0, len(productsDB))
	// for _, productDB := range productsDB {
	// 	if err != nil {
	// 		return nil, errors.New("could not understand category")
	// 	}
	// 	products = append(products, model.Product{
	// 		Title:       productDB.Title,
	// 		Description: productDB.Description,
	// 		Status:      (*model.ProductStatus)(&productDB.Status),
	// 		CreatedAt:   productDB.CreatedAt.Time,
	// 		UpdatedAt:   productDB.UpdatedAt.Time,
	// 	})
	// }

	// return products, nil
	panic(fmt.Errorf("not implemented: Product - product"))
}

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	panic(fmt.Errorf("not implemented: Product - product"))
}

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

type productResolver struct{ *Resolver }
