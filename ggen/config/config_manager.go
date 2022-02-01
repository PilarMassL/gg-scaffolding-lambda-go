package config

import (
	"log"

	"github.com/spf13/viper"
)

const CONFIG_NOT_FOUND = `No se encuentra el archivo de configuración. Recuerde que para agregar funciones debe
inicializar primero el proyecto usando el comando: 'ggen init -b projectbase'. Mas info: `

func LoadConfig() {
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatal(CONFIG_NOT_FOUND, err)
	}
}

func InitConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("ggen")
	viper.SetConfigType("yaml")
}
