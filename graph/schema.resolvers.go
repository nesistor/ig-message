package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.57

import (
	"context"
	"ig-message/graph/model"
	"time"
)

// GetInstagramLikes resolves the `getInstagramLikes` query
func (r *queryResolver) GetInstagramLikes(ctx context.Context, userID string) ([]*model.Like, error) {
	// This is a dummy resolver. Replace with actual logic to fetch Instagram likes.
	// For now, returning a sample "like" data
	return []*model.Like{
		{UserID: userID, Timestamp: time.Now().String()},
	}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
