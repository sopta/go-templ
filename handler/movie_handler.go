package handler

import (
	"go_templ/domain"
	"go_templ/model"
	"go_templ/view/view_movie"
	"net/http"

	"github.com/labstack/echo/v4"
)

var movies = []model.Movie{
	{Title: "Movie Title 1", Director: "Director Name 1"},
	{Title: "Movie Title 2", Director: "Director Name 2"},
	{Title: "Movie Title 3", Director: "Director Name 3"},
	{Title: "Movie Title 4", Director: "Director Name 4"},
	{Title: "Movie Title 5", Director: "Director Name 5"},
}

type MovieHandler struct {
	Query domain.MoviesAllQuery
}

type addMoviePost struct {
	Title    string `form:"title" query:"title"`
	Director string `form:"director" query:"director"`
}

func (h MovieHandler) HandleMovieIndex(c echo.Context) error {
	h.Query.All(c.Request().Context())

	return view_movie.Index(movies).Render(c.Request().Context(), c.Response())
}

func (h MovieHandler) HandleAddMovie(c echo.Context) error {
	//time.Sleep(3 * time.Second)
	u := new(addMoviePost)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	movie := model.Movie{
		Title:    u.Title,
		Director: u.Director,
	}

	return view_movie.MovieListItem(movie).Render(c.Request().Context(), c.Response())
}
