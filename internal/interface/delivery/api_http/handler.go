package apihttp

import (
	link_usecase "link-shortener/internal/domain/link/usecase"
	user_usecase "link-shortener/internal/domain/user/usecase"
	"link-shortener/internal/interface/delivery/api_http/auth"
	"link-shortener/internal/interface/delivery/api_http/link"

	"github.com/gofiber/fiber/v2"
)

func Register(
	router *fiber.App,

	link_usecase link_usecase.Link,
	auth_usecase user_usecase.User,
) {
	ah := auth.NewHandler(auth_usecase)
	lh := link.NewHandler(link_usecase)

	ah.Route(router)
	lh.Route(router)
}
