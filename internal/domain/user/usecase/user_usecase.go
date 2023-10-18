package user_usecase

import (
	"context"
	"fmt"
	"link-shortener/internal/domain/user/entity"
	usermongo "link-shortener/internal/infra/persistence/user/mongodb"
	"link-shortener/internal/interface/delivery/api_http/auth/dto"
	"link-shortener/pkg/pwd"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var tokenEncodeString string = "something"

type User interface {
	SignUp(ctx context.Context, signUpDto dto.SignUpDto) (*entity.User, error)
	SignIn(ctx context.Context, signInDto dto.SignInDto) (string, error)
	SignOut(ctx context.Context, id string) (*entity.User, error)
	GetByUsername(ctx context.Context, hash string) (*entity.User, error)
}

type user struct {
	ur usermongo.Repository
}

func NewUsecase(ur usermongo.Repository) User {
	return &user{ur: ur}
}

func (u *user) SignUp(ctx context.Context, sud dto.SignUpDto) (*entity.User, error) {
	var user entity.User = entity.User{
		Username: sud.Username,
		Password: pwd.HashPassword(sud.Password),
	}

	res, err := u.ur.Create(&user)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *user) SignIn(ctx context.Context, sid dto.SignInDto) (string, error) {
	user, err := u.ur.GetOneByUsername(sid.Username)
	if err != nil {
		return "", err
	}

	if !pwd.ComparePasswords(user.Password, sid.Password) {
		return "", fmt.Errorf("")
	}

	claims := jwt.MapClaims{
		"exp":        time.Now().Add(time.Hour).Unix(),
		"authorized": true,
		"user":       user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(tokenEncodeString))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (u *user) SignOut(ctx context.Context, id string) (*entity.User, error) {
	return nil, nil
}

func (u *user) GetByUsername(ctx context.Context, hash string) (*entity.User, error) {
	return nil, nil
}
