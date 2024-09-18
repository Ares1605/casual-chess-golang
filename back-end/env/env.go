package env

import (
	"os"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}
func Get(key string) (string) {
  return os.Getenv(key)
}
