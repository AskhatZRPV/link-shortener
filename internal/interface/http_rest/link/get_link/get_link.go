package get_link

import (
	"link-shortener/internal/application/link/find"
	"link-shortener/internal/interface/http_rest/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase find.UseCase
}

func New(usecase find.UseCase) common.Handler {
	return &handler{usecase}
}

func (h *handler) Pattern() string {
	return "/:alias"
}

func (h *handler) Method() string {
	return http.MethodGet
}

func (h *handler) Middleware() []fiber.Handler {
	return []fiber.Handler{}
}
