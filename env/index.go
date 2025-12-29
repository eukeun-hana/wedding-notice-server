package env

import (
	"os"

	"github.com/joho/godotenv"
)

var AdminPassword string
var AllowOrigin string

func init() {
	// .env는 "있으면 로드", 없어도 죽지 않음
	_ = godotenv.Load()

	AdminPassword = os.Getenv("ADMIN_PASSWORD")
	AllowOrigin = os.Getenv("ALLOW_ORIGIN")
}
