package handler

import (
	"go_templ/model"
	"go_templ/view/view_user"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (u UserHandler) HandleUserShow(c echo.Context) error {
	user_model := model.User{
		Email: "blabla@gmail.com",
	}

	return view_user.Show(user_model).Render(c.Request().Context(), c.Response())
}
