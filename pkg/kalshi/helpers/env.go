package helpers

import (
	"log"

	"github.com/spf13/viper"
)

func ReadEnvFile() (string, string) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error in ReadEnvFile: '%v'\n", err)
	}
	kalshiUsername, ok := viper.Get("KALSHI_USERNAME").(string)
	if !ok {
		log.Fatalf("Error: KALSHI_USERNAME was not a string\n")
	}
	kalshiPassword, ok := viper.Get("KALSHI_PASSWORD").(string)
	if !ok {
		log.Fatalf("Error: KALSHI_PASSWORD was not a string\n")
	}
	return kalshiUsername, kalshiPassword
}
