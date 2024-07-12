package internal

import (
	"context"
	"gn222gq/rec-sys/internal/db"
	movierecommendations "gn222gq/rec-sys/internal/endpoints/movie-recommendations"
	similarusers "gn222gq/rec-sys/internal/endpoints/similar-users"
	"gn222gq/rec-sys/internal/repository"
	"log"

	"github.com/gofiber/fiber/v2"
)

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

  ctx := context.Background()
  dbPool := db.NewSqliteDb("/home/gpnaslund/GitHub/2dv515_a1/recsys.db")
  dbConn, err := dbPool.GetConnection(ctx)
  if err != nil {
    log.Fatalf("Failed to get database connection")
  }
  repo := repository.NewRepository(dbConn)

  movieRecHandler := movierecommendations.NewHandler()
  v1.Get("/movie-recommendations", movieRecHandler.Handle)

  
  suService := similarusers.NewService(repo)
  suHandler := similarusers.NewHandler(suService)
  v1.Get("/similar-users", suHandler.Handle)

  app.Listen(a.port)
}
