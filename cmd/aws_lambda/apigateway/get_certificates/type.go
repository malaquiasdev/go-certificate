package main

type CertificatesDTO struct {
	Count         int              `json:"count"`
	Items         []CertificateDTO `json:"items"`
	NextPageToken string           `json:"nextPageToken"`
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
