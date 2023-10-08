package main

import (
	"context"
	"link-shortener/internal/adapters/db/mongodb"
	"link-shortener/internal/config"
	v1 "link-shortener/internal/controller/http/v1"
	link "link-shortener/internal/domain/usecase/link"
	"link-shortener/pkg/client/mongo"
	"link-shortener/pkg/logging"

	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.MustLoad()

	log := logging.InitLogger(cfg.Env)

	mongoClient, err := mongo.NewClient(
		context.Background(),
		cfg.Mongo.Host,
		cfg.Mongo.Port,
		cfg.Mongo.Username,
		cfg.Mongo.Password,
		cfg.Mongo.Database,
		cfg.Mongo.Collection,
	)
	if err != nil {
		panic(err)
	}

	mongoRepo := mongodb.NewMongoRepository(mongoClient)
	lu := link.NewLinkUsecase(mongoRepo)

	log = log.With(slog.String("env", cfg.Env))
	log.Info("initializing server", slog.String("address", cfg.Address))

	router := fiber.New()

	v1 := v1.NewHandler(lu)
	v1.Register(router)

	router.Listen(cfg.Address)
}
