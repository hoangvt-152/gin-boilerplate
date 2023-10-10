package config

import (
	"gin-boilerplate/infra/logger"

	"github.com/spf13/viper"
)

type Configuration struct {
	Server       ServerConfiguration
	Database     DatabaseConfiguration
	Qrcodefolder string
}

var QrCodeImageFolder string

// SetupConfig configuration
func SetupConfig() error {
	var configuration *Configuration

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("Error to reading config file, %s", err)
		return err
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		logger.Errorf("error to decode, %v", err)
		return err
	}
	QrCodeImageFolder = viper.GetString("QR_CODE_IMAGE_FOLDER")

	return nil
}
