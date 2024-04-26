package config

import "ekoa-certificate-generator/internal/utils"

type AWS struct {
	Region            string
	AccessKey         string
	SecretAccessKey   string
	GeneretorQueueUrl string
	BucketName        string
	DynamoTableName   string
	IndexerQueueUrl   string
}

type Curseduca struct {
	ClassBaseUrl string
	ProfBaseUrl  string
	Username     string
	Password     string
	ApiKey       string
	BlockList    string
}

type Config struct {
	AWS       AWS
	Curseduca Curseduca
}

func LoadConfig(isLocal bool) Config {
	if isLocal {
		return Config{
			AWS: AWS{
				Region:            utils.GetEnvLocal("AWS_DEFAULT_REGION", ""),
				AccessKey:         utils.GetEnvLocal("AWS_ACCESS_KEY_ID", ""),
				SecretAccessKey:   utils.GetEnvLocal("AWS_SECRET_ACCESS_KEY", ""),
				GeneretorQueueUrl: utils.GetEnvLocal("AWS_GENERATOR_QUEUE_URL", ""),
				BucketName:        utils.GetEnvLocal("AWS_BUCKET_NAME", ""),
				DynamoTableName:   utils.GetEnvLocal("AWS_DYNAMO_TABLE_NAME", ""),
				IndexerQueueUrl:   utils.GetEnvLocal("AWS_INDEXER_QUEUE_URL", ""),
			},
			Curseduca: Curseduca{
				ClassBaseUrl: utils.GetEnvLocal("CLASS_CURSEDUCA_BASE_URL", ""),
				ProfBaseUrl:  utils.GetEnvLocal("PROF_CURSEDUCA_BASE_URL", ""),
				Username:     utils.GetEnvLocal("PROF_CURSEDUCA_USERNAME", ""),
				Password:     utils.GetEnvLocal("PROF_CURSEDUCA_PASSWORD", ""),
				ApiKey:       utils.GetEnvLocal("CURSEDUCA_API_KEY", ""),
				BlockList:    utils.GetEnvLocal("CURSEDUCA_BLOCK_LIST", ""),
			},
		}
	}
	return Config{
		AWS: AWS{
			Region:            utils.GetEnv("AWS_DEFAULT_REGION", ""),
			AccessKey:         utils.GetEnv("AWS_ACCESS_KEY_ID", ""),
			SecretAccessKey:   utils.GetEnv("AWS_SECRET_ACCESS_KEY", ""),
			GeneretorQueueUrl: utils.GetEnv("AWS_GENERATOR_QUEUE_URL", ""),
			BucketName:        utils.GetEnv("AWS_BUCKET_NAME", ""),
			DynamoTableName:   utils.GetEnv("AWS_DYNAMO_TABLE_NAME", ""),
			IndexerQueueUrl:   utils.GetEnv("AWS_INDEXER_QUEUE_URL", ""),
		},
		Curseduca: Curseduca{
			ClassBaseUrl: utils.GetEnv("CLASS_CURSEDUCA_BASE_URL", ""),
			ProfBaseUrl:  utils.GetEnv("PROF_CURSEDUCA_BASE_URL", ""),
			Username:     utils.GetEnv("PROF_CURSEDUCA_USERNAME", ""),
			Password:     utils.GetEnv("PROF_CURSEDUCA_PASSWORD", ""),
			ApiKey:       utils.GetEnv("CURSEDUCA_API_KEY", ""),
			BlockList:    utils.GetEnv("CURSEDUCA_BLOCK_LIST", ""),
		},
	}
}
