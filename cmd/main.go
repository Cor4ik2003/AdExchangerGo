package main

import (
	"AdExchangerGo/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func main() {

	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "post",
		})
	})

	r.Run("8000")

	repository.Init()

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("error")
	}
}
