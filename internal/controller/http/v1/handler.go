package v1

import (
	"link-shortener/internal/controller/http/dto"
	"link-shortener/internal/domain/entity"
	link_usecase "link-shortener/internal/domain/usecase/link"
	"link-shortener/pkg/httperrors"
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const (
	createURL   = "/"
	redirectURL = "/:alias"
)

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
	var linkDto dto.CreateLink

	c.BodyParser(&linkDto)

	link := &entity.Link{
		URL:    linkDto.URL,
		Hash:   linkDto.Hash,
		Domain: linkDto.DomainName,
	}

	link, err := h.lu.Create(c.Context(), linkDto)
	if err != nil {
		h.log.Error(err.Error())
		return c.Status(http.StatusBadRequest).JSON(httperrors.BuildHttpError(err))
	}

	return c.JSON(link)

}

func (h *handler) Redirect(c *fiber.Ctx) error {
	link, err := h.lu.GetByHash(c.Context(), c.Params("alias"))
	if err != nil {
		h.log.Error(err.Error())
		return c.Status(http.StatusNotFound).JSON(httperrors.BuildHttpError(err))
	}

	return c.Redirect(link.URL)
}

func (h *handler) Get(c *fiber.Ctx) error {
	var linkDto dto.GetLink

	err := c.BodyParser(&linkDto)
	if err != nil {
		h.log.Error(err.Error())
		return err
	}

	link, err := h.lu.GetByID(c.Context(), linkDto.ID)
	if err != nil {
		h.log.Error(err.Error())
		return c.Status(http.StatusNotFound).JSON(httperrors.BuildHttpError(err))
	}

	return c.JSON(link)
}
