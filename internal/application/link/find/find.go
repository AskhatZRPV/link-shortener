package find

import (
	"link-shortener/internal/core/tx"
	"link-shortener/internal/core/usecase"
	"link-shortener/internal/domain/link"
)

type Payload struct {
	Alias string
}

type Result = link.Link

type UseCase = usecase.UseCase[*Payload, *Result]

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
