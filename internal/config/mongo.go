package config

import (
	"github.com/LycsHub/crawlab-sdk/internal/constants"
	"github.com/LycsHub/crawlab-sdk/internal/interfaces"
	"os"
)

type WithMongoConfig struct {
	interfaces.WithConfig
}

func (my *WithMongoConfig) GetConfigMap() map[string]string {
	keys := []string{
		constants.ENV_MONGO_HOST,
		constants.ENV_MONGO_PORT,
		constants.ENV_MONGO_USERNAME,
		constants.ENV_MONGO_PASSWORD,
		constants.ENV_MONGO_AUTH_SOURCE,
		constants.ENV_MONGO_AUTH_MECHANISM,
	}

	configMap := make(map[string]string)
	for _, key := range keys {
		configMap[key] = os.Getenv(key)
	}

	return configMap
}
