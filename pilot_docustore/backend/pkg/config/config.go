package config

import (
	"errors"

	"github.com/spf13/viper"
)

const (
	configType       = "properties"
	configName       = "config"
	secretConfigName = "secret"
)

// Key constants for settings
const (
	CONFIG_NAME            = "CONFIG_NAME" // env var
	CONFIG_PATH            = "CONFIG_PATH" // env var
	ENABLE_TLS             = "ENABLE_TLS"
	MONGODB_AUTH_MECHANISM = "MONGODB_AUTH-MECHANISM"
	MONGODB_DATABASE       = "MONGODB_DATABASE"
	MONGODB_COLLECTION     = "MONGODB_COLLECTION"
	SECRET_CONFIG_PATH     = "SECRET_CONFIG_PATH"
	SECRET_CONFIG_NAME     = "SECRET_CONFIG_NAME"
	MONGODB_URL            = "MONGODB_URL"
	MONGODB_USER           = "MONGODB_USER"
	MONGODB_USER_PASSWORD  = "MONGODB_USER_PASSWORD"
	TLS_CA_BUNDLE_PATH     = "TLS_CA_BUNDLE-PATH"
)

func Init() error {
	SetDefaults()
	viper.AutomaticEnv()
	// Add config
	viper.AddConfigPath(viper.GetString(CONFIG_PATH))
	viper.AddConfigPath(viper.GetString(SECRET_CONFIG_PATH))
	viper.SetConfigType(configType)
	viper.SetConfigName(viper.GetString(CONFIG_NAME))

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.SetConfigName(secretConfigName)
	if err := viper.MergeInConfig(); err != nil {
		//logger.Error(fmt.Sprintf("Failed to read config: %s", err))
		return err
	}

	err := verifyConfig()
	if err != nil {
		return err
	}
	watch()
	return nil
}

func verifyConfig() error {
	if viper.GetString(MONGODB_URL) == "" {
		//logger.Info("Database URL is not defined .")
		return errors.New("Database URL is not defined!")
	}
	return nil
}

func SetDefaults() {
	viper.SetDefault(CONFIG_NAME, "config")
	viper.SetDefault(CONFIG_PATH, "/workspace/config")
	viper.SetDefault(SECRET_CONFIG_PATH, "/secret/database-secret")
	viper.SetDefault(SECRET_CONFIG_NAME, "secret")
	viper.SetDefault(ENABLE_TLS, false)
	viper.SetDefault(MONGODB_AUTH_MECHANISM, "")
	viper.SetDefault(MONGODB_DATABASE, "docustore")
	viper.SetDefault(MONGODB_URL, "mongodb://localhost:27017")
	viper.SetDefault(MONGODB_USER_PASSWORD, "")
	viper.SetDefault(MONGODB_USER, "")
}

func watch() {
	viper.WatchConfig()
}
