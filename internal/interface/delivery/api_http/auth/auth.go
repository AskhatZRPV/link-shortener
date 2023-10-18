package auth

import (
	usecase "link-shortener/internal/domain/user/usecase"
	"link-shortener/internal/interface/delivery/api_http/auth/dto"
	"link-shortener/internal/interface/delivery/api_http/errors"
	"net/http"

	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	lu  usecase.User
	log *slog.Logger
}

func NewHandler(lu usecase.User) *AuthHandler {
	return &AuthHandler{lu: lu}
}

func (h AuthHandler) Route(app *fiber.App) {
	user := app.Group(basePath)
	user.Post(signUp, h.SignUp)
	user.Post(signIn, h.SignIn)
	user.Post(signOut, h.SignOut)
}

func (h AuthHandler) SignUp(c *fiber.Ctx) error {
	var su dto.SignUpDto

	if err := c.BodyParser(&su); err != nil {
		return c.Status(http.StatusBadRequest).JSON("")
	}

	user, err := h.lu.SignUp(c.Context(), su)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func (h AuthHandler) SignIn(c *fiber.Ctx) error {
	var si dto.SignInDto

	if err := c.BodyParser(&si); err != nil {
		return c.Status(http.StatusBadRequest).JSON(errors.BuildHttpError(err))
	}

	token, err := h.lu.SignIn(c.Context(), si)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(errors.BuildHttpError(err))
	}

	var l LoginResponse = LoginResponse{
		AccessToken: token,
	}
	return c.JSON(l)
}

func (h AuthHandler) SignOut(c *fiber.Ctx) error {
	return nil
}
