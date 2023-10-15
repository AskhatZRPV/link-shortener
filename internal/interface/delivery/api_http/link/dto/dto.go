package dto

type CreateLinkDto struct {
	URL        string `json:"url,omitempty"`
	Hash       string `json:"hash,omitempty"`
	DomainName string `json:"domain_name,omitempty"`
}

type GetLinkDto struct {
	ID string `json:"id,omitempty"`
}
