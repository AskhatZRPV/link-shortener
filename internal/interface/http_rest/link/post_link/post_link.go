package post_link

import (
	"link-shortener/internal/application/link/create"
	"link-shortener/internal/interface/http_rest/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase create.UseCase
}

func New(usecase create.UseCase) common.Handler {
	return &handler{usecase}
}

func (h *handler) Pattern() string {
	return "/"
}

func (h *handler) Method() string {
	return http.MethodPost
}

func (h *handler) Middleware() []fiber.Handler {
	return []fiber.Handler{}
}
