package httprest

import (
	"context"
	"fmt"
	"link-shortener/internal/core/config"
	"link-shortener/internal/interface/http_rest/common"
	"time"

	"go.uber.org/fx"

	"github.com/gofiber/fiber/v2"
)

func New(lc fx.Lifecycle, config *config.Config, handlers []common.Handler) {
	f := fiber.New(
		fiber.Config{
			ErrorHandler: common.ErrorHandler,
			ReadTimeout:  time.Second * 3,
		},
	)
	for _, h := range handlers {
		fmt.Println(h)
		f.Add(h.Method(), h.Pattern(), append(h.Middleware(), h.Handle)...)
	}

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				err := f.Listen(fmt.Sprintf("localhost:%s", config.Port))
				if err != nil {
					panic(err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return f.ShutdownWithContext(ctx)
		},
	})
}
