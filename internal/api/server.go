package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)


func StartServer() {
	log.Println("Server start up")
	//Чтение ship.json
	file, err := os.Open("resources/data/ship.json")
	if err != nil {
		log.Println("Ошибка при открытии JSON файла:", err)
		return
	}
	defer file.Close()
	//Декодирование JSON данных
	var ships []Ship
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&ships); err != nil {
		log.Println("Ошибка при декодировании JSON данных:", err)
		return
	}
	//Инициализация gin
	r := gin.Default()

	//Настройка статических ресурсов
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
	//запрос на поиск
	r.GET("/", func(c *gin.Context) {
		shipName := c.DefaultQuery("shipName", "")
		var foundShips []Ship
		for _, ship := range ships {
			if strings.HasPrefix(strings.ToLower(ship.ShipName), strings.ToLower(shipName)) {
				foundShips = append(foundShips, ship)
			}
		}
		data := gin.H{
			"ships": foundShips,
		}
		c.HTML(http.StatusOK, "index.tmpl", data)
	})
	//Запрос на получения багажа по id
	r.GET("/ship/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			// Обработка ошибки
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}

		ship := ships[id-1]
		c.HTML(http.StatusOK, "card.tmpl", ship)
	})

	r.Run()

	log.Println("Server down")
}
