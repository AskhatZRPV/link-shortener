package dto

type Link struct {
	URL        string `json:"url,omitempty"`
	Hash       string `json:"hash,omitempty"`
	DomainName string `json:"domain_name,omitempty"`
}

type LinkID struct {
	ID string `json:"id,omitempty"`
}
