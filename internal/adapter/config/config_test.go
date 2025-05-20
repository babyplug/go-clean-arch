package config_test

import (
	"os"
	"testing"

	"clean-arch/internal/adapter/config"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {

	tests := []struct {
		name        string
		expected    config.Config
		isClearEnv  bool
		envFilePath string
	}{
		{
			name:       "if_no_env_file_path_should_get_default_config",
			isClearEnv: true,
			expected: config.Config{
				Port:           "8080",
				MongoURI:       "mongodb://localhost:27017",
				JWTExpiration:  "86400s", // 24 hours
				AllowedOrigins: "*",
				AllowedMethods: "GET,POST,PUT,DELETE,OPTIONS",
				AllowedHeaders: "Origin,Content-Type,Accept,Authorization",
			},
		},
		{
			name:        "if_env_file_path_set_but_not_found_should_read_from_env",
			envFilePath: ".env.example",
			expected: config.Config{
				Port:           "8080",
				MongoURI:       "mongodb://localhost:27017",
				JWTExpiration:  "86400s", // 24 hours
				AllowedOrigins: "*",
				AllowedMethods: "GET,POST,PUT,DELETE,OPTIONS",
				AllowedHeaders: "Origin,Content-Type,Accept,Authorization",
			},
		},
		{
			name:        "if_env_file_path_set_corrected_should_read_from_env_file",
			envFilePath: "../../../.env.mock.yaml",
			expected: config.Config{
				MongoURI: "mongodb://mongo:27017",

				Port: "9999",

				JWTSecret:     "your_jwt_secret",
				JWTExpiration: "7d", // 24 hours

				Env:            "production",
				AllowedOrigins: "http://localhost:3000",
				AllowedMethods: "GET,POST,PUT,DELETE,OPTIONS",
				AllowedHeaders: "Origin,Content-Type, Authorization, Accept",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.isClearEnv {
				_ = os.Setenv("ENV_FILE", "")
			} else {
				_ = os.Setenv("ENV_FILE", test.envFilePath)
			}

			// test logic
			cfg := config.Load()
			defer config.Reset()
			assert.Equal(t, test.expected, *cfg)
		})
	}
}
