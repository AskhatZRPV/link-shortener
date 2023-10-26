package link_handler

// type LinkHandler struct {
// 	lu  link_usecase.Link
// 	log *slog.Logger
// }

// func NewHandler(lu link_usecase.Link) *LinkHandler {
// 	return &LinkHandler{lu: lu}
// }

// func (handler LinkHandler) Route(app *fiber.App, jwtKey string) {
// 	app.Get(redirectURL, handler.Redirect)
// 	app.Use(jwtware.New(jwtware.Config{
// 		SigningKey: jwtware.SigningKey{
// 			JWTAlg: jwtware.HS512,
// 			Key:    []byte(jwtKey)},
// 	}))
// 	app.Post(createURL, handler.Create)
// }

// func (h LinkHandler) Create(c *fiber.Ctx) error {
// 	var linkDto dto.CreateLinkDto

// 	if err := c.BodyParser(&linkDto); err != nil {
// 		h.log.Error(err.Error())
// 		return err
// 	}

// 	link := &entity.Link{
// 		URL:    linkDto.URL,
// 		Hash:   linkDto.Hash,
// 		Domain: linkDto.DomainName,
// 	}

// 	link, err := h.lu.Create(c.Context(), linkDto)
// 	if err != nil {
// 		h.log.Error(err.Error())
// 		return c.Status(http.StatusBadRequest).JSON(errors.BuildHttpError(err))
// 	}

// 	return c.JSON(link)

// }

// func (h LinkHandler) Redirect(c *fiber.Ctx) error {
// 	link, err := h.lu.GetByHash(c.Context(), c.Params("alias"))
// 	if err != nil {
// 		h.log.Error(err.Error())
// 		return c.Status(http.StatusNotFound).JSON(errors.BuildHttpError(err))
// 	}
// 	return c.Redirect(link.URL)
// }

// func (h LinkHandler) Get(c *fiber.Ctx) error {
// 	var linkDto dto.GetLinkDto

// 	if err := c.BodyParser(&linkDto); err != nil {
// 		h.log.Error(err.Error())
// 		return err
// 	}

// 	link, err := h.lu.GetByID(c.Context(), linkDto.ID)
// 	if err != nil {
// 		h.log.Error(err.Error())
// 		return c.Status(http.StatusNotFound).JSON(errors.BuildHttpError(err))
// 	}

// 	return c.JSON(link)
// }
