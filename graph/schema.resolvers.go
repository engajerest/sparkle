package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
"io"
"strconv"
"time"
"sync"
"errors"
"bytes"
gqlparser "github.com/vektah/gqlparser/v2"
"github.com/vektah/gqlparser/v2/ast"
"github.com/99designs/gqlgen/graphql"
"github.com/99designs/gqlgen/graphql/introspection"
"github.com/engajerest/sparkle/Models/subscription"
"github.com/engajerest/sparkle/graph/generated"
"github.com/engajerest/sparkle/graph/model")

	"github.com/engajerest/auth/Models/users"

	"github.com/engajerest/sparkle/graph/generated"
	"github.com/engajerest/sparkle/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.Data) (*model.SubCategory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Sparkle(ctx context.Context) (*model.Sparkle, error) {
	var cat []*model.Category
	var sub []*model.SubCategory
	var mod []*model.Module
	
	var userGetAll []users.User
	for _, user := range userGetAll {
		cat = append(cat, &model.Category{CategoryID: 1, Name: "cate", Type: 1, SortOrder: 1, Status: "Active"})
		print(user.ID)
	}
	for _, user := range userGetAll {
		sub = append(sub, &model.SubCategory{CategoryID: 1, SubCategoryID: 3, Name: "cate", Type: 1, SortOrder: 1, Status: "Active"})
		print(user.ID)
	}
	for _, user := range userGetAll {
		mod = append(mod, &model.Module{CategoryID: 1, Name: "cate", ModuleID: 1, Content: "ontt", ImageURL: "cecedc", LogoURL: "bhszga"})
		print(user.ID)
	}
	return &model.Sparkle{
		Category:    cat,
		Subcategory: sub,
		Module:      mod,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
