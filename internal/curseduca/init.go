package curseduca

import (
	"ekoa-certificate-generator/config"
)

func NewClient(c config.Curseduca) (ICurseduca, error) {
	auth, err := login(c)
	if err != nil {
		return nil, err
	}

	return &Curseduca{
		httpConfig: requestConfig{
			auth:         auth,
			apiKey:       c.ApiKey,
			classBaseUrl: c.ClassBaseUrl,
			profBaseUrl:  c.ProfBaseUrl,
		},
	}, nil
}
