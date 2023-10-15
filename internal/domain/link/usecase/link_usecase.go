package usecase

import (
	"context"
	"link-shortener/internal/domain/link/entity"
	"link-shortener/internal/infra/persistence/mongodb"
	"link-shortener/internal/interface/delivery/api_http/link/dto"
	"link-shortener/pkg/utils/random"
)

type Link interface {
	Create(ctx context.Context, link dto.CreateLinkDto) (*entity.Link, error)
	GetAll(ctx context.Context) ([]entity.Link, error)
	GetByID(ctx context.Context, id string) (*entity.Link, error)
	GetByHash(ctx context.Context, hash string) (*entity.Link, error)
}

type link struct {
	lr mongodb.Repository
}

func NewLinkUsecase(lr mongodb.Repository) Link {
	return &link{lr: lr}
}

func (l *link) Create(ctx context.Context, link dto.CreateLinkDto) (*entity.Link, error) {
	var urlHash string

	if len(link.Hash) != 0 {
		urlHash = link.Hash
		_, err := l.lr.GetOneByHash(urlHash)
		if err == nil {
			return nil, err
		}
	} else {
		for {
			urlHash = random.NewRandomString(6)
			_, err := l.lr.GetOneByHash(urlHash)
			if err != nil {
				break
			}
		}
	}

	res, err := l.lr.Create(&entity.Link{
		URL:    link.URL,
		Hash:   urlHash,
		Domain: link.DomainName,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (l *link) GetAll(ctx context.Context) ([]entity.Link, error) {
	return nil, nil
}

func (l *link) GetByID(ctx context.Context, id string) (*entity.Link, error) {
	link, err := l.lr.GetOne(id)
	if err != nil {
		return nil, err
	}

	return link, nil
}

func (l *link) GetByHash(ctx context.Context, hash string) (*entity.Link, error) {
	link, err := l.lr.GetOneByHash(hash)
	if err != nil {
		return nil, err
	}

	return link, nil
}

func (l *link) Delete(ctx context.Context, id string) {

}
