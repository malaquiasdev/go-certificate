package curseduca

type ICurseduca interface {
	GetReportEnrollment(limit int) (ReportEnrollment, error)
	GetMemberById(memberId int) (Member, error)
}

type Curseduca struct {
	httpConfig requestConfig
}

type requestConfig struct {
	auth         auth
	apiKey       string
	classBaseUrl string
	profBaseUrl  string
}

type auth struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	RedirectUrl  string `json:"redirectUrl"`
	ExpiresAt    string `json:"expiresAt"`
}

type Metadata struct {
	TotalCount int  `json:"totalCount"`
	HasMore    bool `json:"hasmore"`
	Limit      int  `json:"limit"`
	Offset     int  `json:"offset"`
}

type Content struct {
	ID    int    `json:"id"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

type EnrollmentsMember struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Email    string `json:"email"`
	GroupIds []int  `json:"groupIds"`
}

type Report struct {
	ID                int               `json:"id"`
	Content           Content           `json:"content"`
	StartedAt         *string           `json:"startedAt"`  // Pointer for handling null value
	FinishedAt        *string           `json:"finishedAt"` // Pointer for handling null value
	Member            EnrollmentsMember `json:"member"`
	SituationID       int               `json:"situationId"`
	Progress          int               `json:"progress"`
	ExpiresAt         string            `json:"expiresAt,omitempty"`
	ExpirationEnabled bool              `json:"expirationEnabled"`
	Integration       string            `json:"integration"`
}

type ReportEnrollment struct {
	Metadata Metadata `json:"metadata"`
	Reports  []Report `json:"data"`
}

type Member struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	UUID     string `json:"uuid"`
	Email    string `json:"email"`
	Slug     string `json:"slug"`
	Document string `json:"document"`
}
