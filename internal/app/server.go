package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"backend/internal/app/config"
	"backend/internal/app/dsn"
	"backend/internal/app/repository"

	"github.com/gin-gonic/gin"
)

// Application представляет основное приложение.
type Application struct {
	Config       *config.Config
	Repository   *repository.Repository
	RequestLimit int
}

// New создает новый объект Application и настраивает его.
func New(ctx context.Context) (*Application, error) {
	// Инициализируйте конфигурацию
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		return nil, err
	}

	// Инициализируйте подключение к базе данных (DB)
	repo, err := repository.New(dsn.FromEnv())
	if err != nil {
		return nil, err
	}

	// Инициализируйте и настройте объект Application
	app := &Application{
		Config:     cfg,
		Repository: repo,
		// Установите другие параметры вашего приложения, если необходимо
	}

	return app, nil
}

// Run запускает приложение.
func (app *Application) Run() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/css", "./resources/css")
	r.Static("/data", "./resources/data")
	r.Static("/images", "./resources/images")
	r.Static("/fonts", "./resources/fonts")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {

		shipName := c.DefaultQuery("shipName", "")

		ships, err := app.Repository.GetShips(shipName)
		if err != nil {
			log.Println("Error Repository method GetAll:", err)
			return
		}
		data := gin.H{
			"ships":    ships,
			"shipName": shipName,
		}
		c.HTML(http.StatusOK, "index.tmpl", data)
	})

	r.GET("/ship/:id", func(c *gin.Context) {
		shipID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			// Обработка ошибки
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}

		ship, err := app.Repository.GetShipByID(uint(shipID))
		if err != nil {
			// Обработка ошибки
			c.JSON(http.StatusBadRequest, gin.H{"error": "GetShipByID"})
			return
		}

		c.HTML(http.StatusOK, "card.tmpl", ship)
	})

	r.POST("/ship/:id", func(c *gin.Context) {

		shipID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			// Обработка ошибки
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}
		app.Repository.DeleteShip(uint(shipID))
		c.Redirect(http.StatusMovedPermanently, "/")
	})
	r.GET("/request", func(c *gin.Context) {

		ships, err := app.Repository.GetShipsFromRequest(29)
		if err != nil {
			log.Println("Error Repository method GetAll:", err)
			return
		}
		data := gin.H{
			"ships": ships,
		}
		c.HTML(http.StatusOK, "constructor.tmpl", data)
	})
	addr := fmt.Sprintf("%s:%d", app.Config.ServiceHost, app.Config.ServicePort)
	r.Run(addr)
	log.Println("Server down")
}
