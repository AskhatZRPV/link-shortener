package post_link

import (
	"link-shortener/internal/application/link/create"
)

type requestBody struct {
	Alias       string `json:"alias"`
	RedirectUrl string `json:"redirect_url"`
}

func (r *requestBody) toUsecasePayload() *create.Payload {
	return &create.Payload{
		Alias:       r.Alias,
		RedirectUrl: r.RedirectUrl,
	}
}
