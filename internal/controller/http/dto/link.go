package dto

type CreateLink struct {
	URL        string `json:"url,omitempty"`
	Hash       string `json:"hash,omitempty"`
	DomainName string `json:"domain_name,omitempty"`
}

type GetLink struct {
	ID string `json:"id,omitempty"`
}
