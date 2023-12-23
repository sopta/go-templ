package sqlite

import (
	"context"
	"go_templ/ent"
	"log"
)

type MoviesQuery struct {
	Client *ent.Client
}

func (q MoviesQuery) All(ctx context.Context) { // error
	u, _ := q.Client.Movie.Query().All(ctx)

	/*
		if err != nil {
			return nil, fmt.Errorf("(q MoviesQuery) All : querying user: %w", err)
		}*/

	log.Println("user returned: ", u)
}
