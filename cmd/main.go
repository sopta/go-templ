package main

import (
	"context"
	"fmt"
	"go_templ/ent"
	"go_templ/handler"
	"go_templ/store/sqlite"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

var db_client *ent.Client

func main() {
	var db_client_err error

	db_client, db_client_err = ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if db_client_err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", db_client_err)
	}
	defer db_client.Close()

	// Run the auto migration tool.
	if err := db_client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	movie_create(context.Background(), "Movie Title 1", "Director Name 1")
	movie_create(context.Background(), "Movie Title 2", "Director Name 2")
	movie_create(context.Background(), "Movie Title 3", "Director Name 3")
	movie_create(context.Background(), "Movie Title 4", "Director Name 4")

	count, _ := db_client.Movie.Query().Count(context.Background())
	log.Println(count)

	app := echo.New()
	app.Use(withUser)

	userHandler := handler.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)

	movieHandler := handler.MovieHandler{
		Query: sqlite.MoviesQuery{
			Client: db_client,
		},
	}
	//app.GET("/movies", movieHandler.HandleMovieIndex)
	app.GET("/movies", InitializeMovieHandler().HandleMovieIndex)
	app.POST("/add-movie", movieHandler.HandleAddMovie)

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	app.Logger.Fatal(app.Start(":1323"))
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "user", "kokot_hoho")

		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

func movie_create(ctx context.Context, title string, director string) (*ent.Movie, error) {
	u, err := db_client.Movie.
		Create().
		SetTitle(title).
		SetDirector(director).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating movie: %w", err)
	}

	log.Println("movie was created: ", u)
	return u, nil
}
