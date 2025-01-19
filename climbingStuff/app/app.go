package app

import (
	"climbingStuff/handlers"
	"climbingStuff/services"
	"database/sql"

	"github.com/labstack/echo/v4"

	_ "climbingStuff/docs" // Import Swagger generated docs

	echoSwagger "github.com/swaggo/echo-swagger"
)

type App struct {
	DB          *sql.DB
	GymService  services.ClimbingGymService
	ShoeService services.ClimbingShoeService
	Echo        *echo.Echo
}

func NewApp(
	db *sql.DB,
	gymService services.ClimbingGymService,
	shoeService services.ClimbingShoeService,
	gymHandler handlers.ClimbingGymHandler,
	shoeHandler handlers.ClimbingShoeHandler,
) *App {
	e := echo.New()

	// Setup routes
	e.GET("/climbing-gyms", gymHandler.GetAll)
	e.GET("/climbing-gyms/:id", gymHandler.GetById)
	e.GET("/climbing-gyms/city/:city", gymHandler.GetByCity)
	e.POST("/climbing-gyms", gymHandler.Insert)
	e.DELETE("/climbing-gyms/:id", gymHandler.DeleteById)

	e.GET("/climbing-shoes", shoeHandler.GetAll)
	e.POST("/climbing-shoes", shoeHandler.Insert)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return &App{
		Echo:        e,
		GymService:  gymService,
		ShoeService: shoeService,
		DB:          db,
	}
}

func (app *App) Run(address string) error {
	return app.Echo.Start(address)
}

func (app *App) Close() error {
	if app.DB != nil {
		return app.DB.Close()
	}
	return nil
}
