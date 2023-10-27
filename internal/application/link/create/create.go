package create

import (
	"link-shortener/internal/core/tx"
	"link-shortener/internal/core/usecase"
	"link-shortener/internal/domain/link"
)

type Payload struct {
	Alias, RedirectUrl string
}

type UseCase = usecase.Interactor[*Payload]

type implementation struct {
	txManager tx.TransactionManager
	linkRepo  link.Repository
}

func New(
	txManager tx.TransactionManager,
	linkRepo link.Repository,
) UseCase {
	return &implementation{txManager, linkRepo}
}
