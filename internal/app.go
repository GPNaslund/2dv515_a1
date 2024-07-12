package internal

import "github.com/gofiber/fiber/v2"

type App struct {
  port string

}

func NewApp(port string) *App {
	return &App{
    port: port,
  }
}

func (a *App) Run() {
	app := fiber.New()

  api := app.Group("/api")

  v1 := api.Group("/v1")

  v1.Get("/movie-recommendations", )
  v1.Get("/similar-users", )

  app.Listen(a.port)
}
