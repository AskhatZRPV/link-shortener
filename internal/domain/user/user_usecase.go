package user

import (
	"context"
	"time"
)

// package user

// import (
// 	"context"
// 	"time"
// )

type Repository interface {
	Save(context.Context, *User) error
	FindByUsername(context.Context, string) (*User, error)
	UpdateLastLoginAttempt(context.Context, ID, time.Time) error
}

// type Repository interface {
// 	SignUp(ctx context.Context, signUpDto dto.SignUpDto) (*entity.User, error)
// 	SignIn(ctx context.Context, signInDto dto.SignInDto) (string, error)
// 	SignOut(ctx context.Context, id string) (*entity.User, error)
// 	GetByUsername(ctx context.Context, hash string) (*entity.User, error)
// }

// type user struct {
// 	ur       usermongo.Repository
// 	jwtToken jwt_token.JWTToken
// }

// func NewUsecase(ur usermongo.Repository, jwtKey string) User {
// 	return &user{ur: ur, jwtKey: jwtKey}
// }
