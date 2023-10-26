package login

import (
	"link-shortener/internal/core/hasher"
	"link-shortener/internal/core/tx"
	"link-shortener/internal/core/usecase"
	"link-shortener/internal/domain/session"
	"link-shortener/internal/domain/token"
	"link-shortener/internal/domain/user"
)

type Payload struct {
	Username, Password string
}

type Result = token.GrantResult

type UseCase = usecase.UseCase[*Payload, *Result]

type implementation struct {
	txManager     tx.TransactionManager
	ph            hasher.Hasher
	userRepo      user.Repository
	tokenProvider token.Provider
	refreshRepo   session.Repository
}

func New(
	txManager tx.TransactionManager,
	ph hasher.Hasher,
	userRepo user.Repository,
	tokenProvider token.Provider,
	refreshRepo session.Repository,
) UseCase {
	return &implementation{txManager, ph, userRepo, tokenProvider, refreshRepo}
}
