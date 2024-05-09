package main

type CertificatesDTO struct {
	NextPageKey string           `json:"nextPageKey"`
	Count       *int64           `json:"count"`
	Items       []CertificateDTO `json:"items"`
}

type CertificateDTO struct {
	ID        string `json:"id"`
	ContentId int    `json:"contentId"`
	StudentId int    `json:"studentId"`
	URL       string `json:"url"`
}
