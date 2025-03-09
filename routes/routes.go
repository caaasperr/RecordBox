package routes

import (
	shelf "myvinyl/modules/shelf/handler"
	user "myvinyl/modules/user/handler"
	vinyl "myvinyl/modules/vinyl/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	//User Module Routing
	userR := v1.Group("/auth")
	userR.Post("/signup", user.SignUpHandler)
	userR.Post("/login", user.LogInHandler)
	userR.Post("/logout", user.LogOutHandler) //Except
	userR.Get("/session", user.SessionCheckHandelr)

	//Methods for authorized users
	authedUserR := v1.Group("/user", user.UserMiddleware)
	authedUserR.Get("/", user.GetProfileHandler)
	authedUserR.Put("/", user.UpdateProfileHandler)
	authedUserR.Delete("/", user.DeleteUserHandler)

	vinylR := v1.Group("/vinyls", user.UserMiddleware)
	vinylR.Post("/covers", vinyl.GetAlbumCoversFromLastFmByNameHandler)
	vinylR.Get("/", vinyl.GetVinylsHandler)
	vinylR.Post("/", vinyl.CreateVinyl)
	vinylR.Get("/:id", vinyl.GetVinylHandler)
	vinylR.Delete("/:id", vinyl.DeleteVinylHandler)
	vinylR.Put("/:id", vinyl.UpdateVinylHandler)
	vinylR.Put("/:id/slot", vinyl.UpdateVinylSlotHandler)
	vinylR.Get("/genre/:id", vinyl.GetVinylsByGenreHandler)

	genreR := v1.Group("/genres")
	genreR.Get("/", vinyl.GetGenresHandler)

	shelfR := v1.Group("/shelves", user.UserMiddleware)
	shelfR.Post("/", shelf.CreateShelfHandler)
	shelfR.Get("/", shelf.GetShelvesHandler)
	shelfR.Get("/:id", shelf.GetShelfHandler)
	shelfR.Delete("/:id", shelf.DeleteShelfHandler)
	shelfR.Put("/:id", shelf.UpdateBookshelfHandler)
	shelfR.Get("/:id/slots/:value", vinyl.GetVinylsByShelfSlotHandler)
	shelfR.Put("/:id/slots/state", shelf.UpdateShelfSlotStateHandler)

	slotR := v1.Group("/slots", user.UserMiddleware)
	slotR.Get("/:id", shelf.GetShelfslotByIdHandler)

	//Methods for admin
	adminUserR := v1.Group("/users", user.AdminMiddleware)
	adminUserR.Get("/", user.GetAllUsersHandler)
	adminGenreR := v1.Group("/genres", user.AdminMiddleware)
	adminGenreR.Post("/", vinyl.CreateGenreHandler)
	adminGenreR.Delete("/:id", vinyl.DeleteGenreHandler)
}
