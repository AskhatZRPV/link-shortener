package get_link

import (
	"link-shortener/internal/application/link/find"
)

type requestBody struct {
	Alias string `json:"alias"`
}

func (r *requestBody) toUsecasePayload() *find.Payload {
	return &find.Payload{
		Alias: r.Alias,
	}
}
