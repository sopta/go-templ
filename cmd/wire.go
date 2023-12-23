//go:build wireinject
// +build wireinject

package main

import "go_templ/handler"

func InitializeMovieHandler() handler.MovieHandler {
	return handler.MovieHandler{}
}
