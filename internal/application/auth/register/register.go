package register

import (
	"link-shortener/internal/core/hasher"
	"link-shortener/internal/core/tx"
	"link-shortener/internal/core/usecase"
	"link-shortener/internal/domain/user"
)

type Payload struct {
	Username, Password string
}

type UseCase = usecase.Interactor[*Payload]

type implementation struct {
	txManager tx.TransactionManager
	ph        hasher.Hasher
	userRepo  user.Repository
}

func New(
	txManager tx.TransactionManager,
	ph hasher.Hasher,
	userRepo user.Repository,
) UseCase {
	return &implementation{txManager, ph, userRepo}
}
