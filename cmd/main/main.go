package main

import (
	"context"
	"link-shortener/internal/config"
	link_usecase "link-shortener/internal/domain/link/usecase"
	user_usecase "link-shortener/internal/domain/user/usecase"
	linkmongo "link-shortener/internal/infra/persistence/link/mongodb"
	usermongo "link-shortener/internal/infra/persistence/user/mongodb"
	apihttp "link-shortener/internal/interface/delivery/api_http"
	"link-shortener/pkg/client/mongo"
	"link-shortener/pkg/logging"

	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	cfg := config.MustLoad()

	log := logging.InitLogger(cfg.Env)

	mongo, err := mongo.NewClient(
		context.Background(),
		cfg.Mongo.Host,
		cfg.Mongo.Port,
		cfg.Mongo.Username,
		cfg.Mongo.Password,
		cfg.Mongo.Database,
	)
	if err != nil {
		panic(err)
	}

	lr := linkmongo.NewMongoRepository(mongo.Collection("link"))
	ur := usermongo.NewMongoRepository(mongo.Collection("user"))

	lu := link_usecase.NewUsecase(lr)
	uu := user_usecase.NewUsecase(ur)

	log = log.With(slog.String("env", cfg.Env))
	log.Info("initializing server", slog.String("address", cfg.Address))

	app := fiber.New()
	// Logging Request ID
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

	apihttp.Register(app, lu, uu)

	app.Listen(cfg.Address)
}
