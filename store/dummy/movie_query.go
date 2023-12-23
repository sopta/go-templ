package dummy

import (
	"context"
	"log"
)

type MoviesQuery struct {
}

func (q MoviesQuery) All(ctx context.Context) { // error
	log.Println("user returned: ")
}
