package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/rajatjain007/GoGraphQL/database"
	"github.com/rajatjain007/GoGraphQL/graph/generated"
	"github.com/rajatjain007/GoGraphQL/graph/model"
)

// CreateDog is the resolver for the createDog field.
func (r *mutationResolver) CreateDog(ctx context.Context, input *model.NewDog) (*model.Dog, error) {
	// panic(fmt.Errorf("not implemented: CreateDog - createDog"))
	return db.SaveDog(input), nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	// panic(fmt.Errorf("not implemented: CreateUser - createUser"))
	return db.SaveUser((input)), nil
}

// Dog is the resolver for the dog field.
func (r *queryResolver) Dog(ctx context.Context, id *string) (*model.Dog, error) {
	// panic(fmt.Errorf("not implemented: Dog - dog"))
	return db.FindDogByID(*id), nil
}

// Dogs is the resolver for the dogs field.
func (r *queryResolver) Dogs(ctx context.Context) ([]*model.Dog, error) {
	// panic(fmt.Errorf("not implemented: Dogs - dogs"))
	return db.FindAllDogs(), nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id *string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return db.FindAllUsers(),nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
var db = database.Connect()
