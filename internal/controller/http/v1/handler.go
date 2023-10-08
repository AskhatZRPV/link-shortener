package v1

import (
	"fmt"
	"link-shortener/internal/controller/http/dto"
	"link-shortener/internal/domain/entity"
	link_usecase "link-shortener/internal/domain/usecase/link"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

const (
	createURL   = "/"
	redirectURL = "/:alias"
)

type LinkHandler interface {
	Register(router *fiber.App)
	Create(c *fiber.Ctx)
	Get(c *fiber.Ctx)
	Redirect(c *fiber.Ctx)
}

type handler struct {
	lu  link_usecase.Link
	log *slog.Logger
}

func NewHandler(lu link_usecase.Link) *handler {
	return &handler{lu: lu}
}

func (h *handler) Register(router *fiber.App) {
	router.Post(createURL, h.Create)
	router.Get(redirectURL, h.Redirect)
	router.Get(createURL, h.Get)

}

func (h *handler) Create(c *fiber.Ctx) error {
	var linkDto *dto.Link

	c.BodyParser(&linkDto)

	link := &entity.Link{
		URL:    linkDto.URL,
		Hash:   linkDto.Hash,
		Domain: linkDto.DomainName,
	}
	link, err := h.lu.Create(c.Context(), linkDto)
	if err != nil {
		return err
	}

	return c.JSON(link)

}

func (h *handler) Redirect(c *fiber.Ctx) error {
	link, err := h.lu.GetByHash(c.Context(), c.Params("alias"))
	if err != nil {
		return err
	}
	return c.Redirect(link.URL)
}

func (h *handler) Get(c *fiber.Ctx) error {
	fmt.Println("123213132131321")
	var link dto.LinkID
	err := c.BodyParser(&link)
	if err != nil {
		return err
	}

	link2, err := h.lu.GetByID(c.Context(), link.ID)
	if err != nil {
		return err
	}
	// b, err := json.Marshal(*link2)
	// if err != nil {
	// 	return err
	// }
	return c.JSON(link2)
}
