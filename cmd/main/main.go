package main

import (
	"link-shortener/internal/application/auth/login"
	"link-shortener/internal/application/auth/register"
	"link-shortener/internal/core/config"
	"link-shortener/internal/infrastructure/crosscutting/bcrypt"
	"link-shortener/internal/infrastructure/crosscutting/jwt_token"
	"link-shortener/internal/infrastructure/crosscutting/pgsql_client"
	sessionpgsql "link-shortener/internal/infrastructure/repository/session/pgsql"
	"link-shortener/internal/infrastructure/repository/tx/pgsqltx"
	userpgsql "link-shortener/internal/infrastructure/repository/user/pgsql"
	httprest "link-shortener/internal/interface/http_rest"
	"link-shortener/internal/interface/http_rest/auth/post_login"
	"link-shortener/internal/interface/http_rest/auth/post_register"
	"time"

	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		fx.Provide(config.New),
		fx.Provide(bcrypt.New),
		fx.Provide(pgsql_client.New),
		fx.Provide(pgsqltx.New),
		fx.Provide(
			sessionpgsql.New,
			userpgsql.New,
		),
		fx.Provide(
			jwt_token.New,
		),
		fx.Provide(
			register.New,
			login.New,
		),
		fx.Provide(
			fx.Annotate(
				post_login.New,
				fx.ResultTags(`group:"handlers"`),
			),
			fx.Annotate(
				post_register.New,
				fx.ResultTags(`group:"handlers"`),
			),
		),
		fx.Invoke(
			fx.Annotate(httprest.New, fx.ParamTags(``, ``, `group:"handlers"`)),
		),
		fx.StartTimeout(time.Second),
	)

	app.Run()
}
