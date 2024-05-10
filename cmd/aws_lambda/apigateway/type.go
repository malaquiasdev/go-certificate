package main

type CertificatesDTO struct {
	NextPageKey string           `json:"nextPageKey"`
	Count       *int64           `json:"count"`
	Items       []CertificateDTO `json:"items"`
}

type CertificateDTO struct {
	ID         string `json:"id"`
	ContentId  int    `json:"contentId"`
	StudentId  int    `json:"studentId"`
	CreatedAt  string `json:"createdAt"`
	ExpiresAt  string `json:"expiresAt,omitempty"`
	FinishedAt string `json:"finishedAt,omitempty"`
	URL        string `json:"url"`
}
