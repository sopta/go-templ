package domain

import "context"

type MoviesAllQuery interface {
	All(ctx context.Context)
}
