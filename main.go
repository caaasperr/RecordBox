package main

import (
	"myvinyl/routes"
	"myvinyl/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	utils.InitializeLogger()
	utils.InitializeDB()
	utils.InitializeRedis()
	utils.SetLastFmEnv()
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://recordbox.org, https://m.recordbox.org", // Allow the frontend's origin
		AllowMethods:     "GET,POST,PUT,DELETE",                            // Allowed HTTP methods
		AllowHeaders:     "Content-Type,Authorization",                     // Allowed headers
		AllowCredentials: true,                                             // Allow credentials (cookies, headers)
	}))
	//app.Use(csrf.New(csrf.Config{
	//	KeyLookup:      "header:X-CSRF-Token",
	//	CookieHTTPOnly: true,
	//}))
	routes.SetupRoute(app)
	certFile := "/etc/letsencrypt/live/api.recordbox.org/fullchain.pem"
	keyFile := "/etc/letsencrypt/live/api.recordbox.org/privkey.pem"

	// HTTPS 서버 시작
	utils.Logger.Trace("Starting HTTPS server on https://recordbox.org:443")
	if err := app.ListenTLS(":443", certFile, keyFile); err != nil {
		utils.Logger.Fatal(err)
	}
}
