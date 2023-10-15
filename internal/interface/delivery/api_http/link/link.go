package linkhandler

import (
	"link-shortener/internal/domain/link/entity"
	"link-shortener/internal/domain/link/usecase"
	"link-shortener/internal/interface/delivery/api_http/errors"
	"link-shortener/internal/interface/delivery/api_http/link/dto"
	"net/http"

	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type LinkHandler struct {
	lu  usecase.Link
	log *slog.Logger
}

func NewHandler(lu usecase.Link) *LinkHandler {
	return &LinkHandler{lu: lu}
}

func (handler LinkHandler) Route(app *fiber.App) {
	app.Post(createURL, handler.Create)
	app.Get(redirectURL, handler.Redirect)
}

func (h LinkHandler) Create(c *fiber.Ctx) error {
	var linkDto dto.CreateLinkDto

	if err := c.BodyParser(&linkDto); err != nil {

	}

	link := &entity.Link{
		URL:    linkDto.URL,
		Hash:   linkDto.Hash,
		Domain: linkDto.DomainName,
	}

	link, err := h.lu.Create(c.Context(), linkDto)
	if err != nil {
		h.log.Error(err.Error())
		return c.Status(http.StatusBadRequest).JSON(errors.BuildHttpError(err))
	}

	return c.JSON(link)

}

func (h LinkHandler) Redirect(c *fiber.Ctx) error {
	link, err := h.lu.GetByHash(c.Context(), c.Params("alias"))
	if err != nil {
		h.log.Error(err.Error())
		return c.Status(http.StatusNotFound).JSON(errors.BuildHttpError(err))
	}
	return c.Redirect(link.URL)
}

func (h LinkHandler) Get(c *fiber.Ctx) error {
	var linkDto dto.GetLinkDto

	if err := c.BodyParser(&linkDto); err != nil {
		h.log.Error(err.Error())
		return err
	}

	link, err := h.lu.GetByID(c.Context(), linkDto.ID)
	if err != nil {
		h.log.Error(err.Error())
		return c.Status(http.StatusNotFound).JSON(errors.BuildHttpError(err))
	}

	return c.JSON(link)
}
