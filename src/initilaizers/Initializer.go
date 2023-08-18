package initilaizers

import "github.com/Valgard/godotenv"

func Init() {
	dotenv := godotenv.New()
	if err := dotenv.Load(".env"); err != nil {
		panic(err)
	}
}
