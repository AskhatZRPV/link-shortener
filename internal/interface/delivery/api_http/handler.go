package apihttp

import (
	"link-shortener/internal/domain/link/usecase"
	linkhandler "link-shortener/internal/interface/delivery/api_http/link"

	"github.com/gofiber/fiber/v2"
)

func Register(
	router *fiber.App,
	link_usecase usecase.Link,
) {
	lh := linkhandler.NewHandler(link_usecase)

	lh.Route(router)
}
