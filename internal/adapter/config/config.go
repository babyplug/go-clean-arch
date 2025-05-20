package config

import (
	"log"
	"os"
	"reflect"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var (
	_config = Config{
		MongoURI: "mongodb://localhost:27017",

		Port:          "8080",
		JWTExpiration: "86400s", // 24 hours

		AllowedOrigins: "*",
		AllowedMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowedHeaders: "Origin,Content-Type,Accept,Authorization",
	}
	_configOnce sync.Once
)

type Config struct {
	MongoURI       string `mapstructure:"MONGO_URI"`
	Port           string `mapstructure:"PORT"`
	Env            string `mapstructure:"ENV"`
	AllowedOrigins string `mapstructure:"ALLOWED_ORIGINS"`
	AllowedMethods string `mapstructure:"ALLOWED_METHODS"`
	AllowedHeaders string `mapstructure:"ALLOWED_HEADERS"`

	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpiration string `mapstructure:"JWT_EXPIRATION"`
}

func Load() *Config {
	_configOnce.Do(func() {
		envFilePath, ok := os.LookupEnv("ENV_FILE")
		if len(envFilePath) > 0 && ok {
			log.Printf("env file is set reading config from file %s\n", envFilePath)
			viper.SetConfigFile(envFilePath)
			if err := viper.ReadInConfig(); err != nil {
				log.Printf("read config from file %s failed: %v\tcontinue reading from `env`\n", envFilePath, err)
				viper.AutomaticEnv()
			}
		} else {
			log.Println("env file not set reading config from `env")
			viper.AutomaticEnv()
		}

		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		bindEnv(_config)

		err := viper.Unmarshal(&_config)
		if err != nil {
			log.Fatalf("Failed to unmarshal config: %v", err)
		}
		log.Printf("Config: %+v\n", _config)
	})
	return &_config
}

func Reset() {
	_configOnce = sync.Once{}
}

func bindEnv(dest any, parts ...string) {
	ifv := reflect.ValueOf(dest)
	ift := reflect.TypeOf(dest)

	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)

		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}

		switch v.Kind() {
		case reflect.Struct:
			bindEnv(v.Interface(), append(parts, tv)...)
		default:
			envKey := strings.Join(append(parts, tv), ".")
			err := viper.BindEnv(envKey)
			if err != nil {
				log.Printf("bind env key %s failed: %v\n", envKey, err)
			}
		}
	}
}
