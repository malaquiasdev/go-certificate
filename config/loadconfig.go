package config

import "ekoa-certificate-generator/internal/utils"

type AWS struct {
	Region          string
	AccessKey       string
	SecretAccessKey string
}

type Curseduca struct {
	ClassBaseUrl string
	ProfBaseUrl  string
	Username     string
	Password     string
	ApiKey       string
}

type Config struct {
	AWS       AWS
	Curseduca Curseduca
}

func LoadConfig(isLocal bool) Config {
	if isLocal {
		return Config{
			AWS: AWS{
				Region:          utils.GetEnvLocal("AWS_DEFAULT_REGION", ""),
				AccessKey:       utils.GetEnvLocal("AWS_ACCESS_KEY_ID", ""),
				SecretAccessKey: utils.GetEnvLocal("AWS_SECRET_ACCESS_KEY", ""),
			},
			Curseduca: Curseduca{
				ClassBaseUrl: utils.GetEnvLocal("CLASS_CURSEDUCA_BASE_URL", ""),
				ProfBaseUrl:  utils.GetEnvLocal("PROF_CURSEDUCA_BASE_URL", ""),
				Username:     utils.GetEnvLocal("PROF_CURSEDUCA_USERNAME", ""),
				Password:     utils.GetEnvLocal("PROF_CURSEDUCA_PASSWORD", ""),
				ApiKey:       utils.GetEnvLocal("CURSEDUCA_API_KEY", ""),
			},
		}
	}
	return Config{
		AWS: AWS{
			Region:          utils.GetEnv("AWS_DEFAULT_REGION", ""),
			AccessKey:       utils.GetEnv("AWS_ACCESS_KEY_ID", ""),
			SecretAccessKey: utils.GetEnv("AWS_SECRET_ACCESS_KEY", ""),
		},
		Curseduca: Curseduca{
			ClassBaseUrl: utils.GetEnv("CLASS_CURSEDUCA_BASE_URL", ""),
			ProfBaseUrl:  utils.GetEnv("PROF_CURSEDUCA_BASE_URL", ""),
			Username:     utils.GetEnv("PROF_CURSEDUCA_USERNAME", ""),
			Password:     utils.GetEnv("PROF_CURSEDUCA_PASSWORD", ""),
			ApiKey:       utils.GetEnv("CURSEDUCA_API_KEY", ""),
		},
	}
}
