package main

import (
	"link-shortener/internal/application/auth/login"
	"link-shortener/internal/application/auth/register"
	"link-shortener/internal/application/link/create"
	"link-shortener/internal/application/link/find"
	"link-shortener/internal/core/config"
	"link-shortener/internal/infrastructure/crosscutting/bcrypt"
	"link-shortener/internal/infrastructure/crosscutting/jwt_token"
	"link-shortener/internal/infrastructure/crosscutting/pgsql_client"
	linkpgsql "link-shortener/internal/infrastructure/repository/link/pgsql"
	sessionpgsql "link-shortener/internal/infrastructure/repository/session/pgsql"
	userpgsql "link-shortener/internal/infrastructure/repository/user/pgsql"

	"link-shortener/internal/infrastructure/repository/tx/pgsqltx"
	httprest "link-shortener/internal/interface/http_rest"
	"link-shortener/internal/interface/http_rest/auth/post_login"
	"link-shortener/internal/interface/http_rest/auth/post_register"
	"link-shortener/internal/interface/http_rest/link/get_link"
	"link-shortener/internal/interface/http_rest/link/post_link"
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
			linkpgsql.New,
		),
		fx.Provide(
			jwt_token.New,
		),
		fx.Provide(
			register.New,
			login.New,
			create.New,
			find.New,
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
			fx.Annotate(
				post_link.New,
				fx.ResultTags(`group:"handlers"`),
			),
			fx.Annotate(
				get_link.New,
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
