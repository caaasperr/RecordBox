package utils

import (
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis/v3"
)

var SessionManager *session.Store

func InitializeRedis() {
	Logger.Info("Initializing redis...")
	redisStorage := redis.New(redis.Config{
		Host:     "127.0.0.1",
		Port:     6379,
		Password: "",
		Database: 0,
		PoolSize: 10 * runtime.GOMAXPROCS(0),
	})

	SessionManager = session.New(session.Config{
		Storage:        redisStorage,
		CookieHTTPOnly: true,
		CookieSecure:   true,
		CookieSameSite: "Strict",
		Expiration:     24 * 30 * time.Hour,
	})
	Logger.Info("Initialized redis.")
}
