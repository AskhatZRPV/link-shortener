package get_link

import (
	"link-shortener/internal/application/link/create"
	"link-shortener/internal/interface/http_rest/common"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func (h *handler) Handle(c *fiber.Ctx) error {
	var body requestBody
	if err := c.ParamsParser(&body); err != nil {
		return errors.Wrap(err, "failed to parse login body")
	}

	res, err := h.usecase.Execute(c.Context(), body.toUsecasePayload())
	if err != nil {
		return h.resolveErr(c, err)
	}

	return c.Redirect(res.RedirectUrl, fiber.StatusFound)
}

func (h *handler) resolveErr(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, create.ErrLinkAlreadyExists):
		return common.ErrorBuilder(err).Detail("link", "Link already exists").Build()
	}

	return err
}
