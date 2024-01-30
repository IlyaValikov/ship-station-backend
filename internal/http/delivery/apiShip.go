package delivery

import (
	"io"
	"net/http"
	"strconv"

	"backend/internal/auth"
	"backend/internal/model"

	"github.com/gin-gonic/gin"
)

// @Summary Получение списка судов
// @Description Возращает список всех активных судов
// @Tags Судно
// @Produce json
// @Param shipName query string false "Название судна" Format(email)
// @Success 200 {object} model.GetShips "Список судов"
// @Failure 400 {object} model.ErrorResponse "Обработанная ошибка сервера"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /ship [get]
func (h *Handler) GetShips(c *gin.Context) {
    authInstance := auth.GetAuthInstance()
	userID := uint(authInstance.UserID)
    shipName := c.DefaultQuery("shipName", "")

    ships, err := h.UseCase.GetShips(shipName,userID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"ships": ships.Ships, "requestID":ships.RequestID})
}

// @Summary Получение судна по ID
// @Description Возвращает информацию о суднe по его ID
// @Tags Судно
// @Produce json
// @Param shipID path int true "ID судна"
// @Success 200 {object} model.Ship "Информация о суднe"
// @Failure 400 {object} model.ErrorResponse "Обработанная ошибка сервера"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /ship/{shipID} [get]
func (h *Handler) GetShipByID(c *gin.Context) {
    authInstance := auth.GetAuthInstance()
	userID := uint(authInstance.UserID)

    shipID, err := strconv.Atoi(c.Param("shipID"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "недопустимый ИД судна"})
        return
    }

    ship, err := h.UseCase.GetShipByID(uint(shipID), userID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"ship": ship})
}

// @Summary Создание нового судна
// @Description Создает новое судно с предоставленными данными
// @Tags Судно
// @Accept json
// @Produce json
// @Param shipName query string false "Название судна" Format(email)
// @Param ship body model.ShipRequest true "Пользовательский объект в формате JSON"
// @Success 200 {object} model.GetShips "Список судов"
// @Failure 400 {object} model.ErrorResponse "Обработанная ошибка сервера"
// @Failure 401 {object} model.ErrorResponse "Пользователь не авторизован"
// @Failure 403 {object} model.ErrorResponse "У пользователя нет прав для этого запроса"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Security ApiKeyAuth
// @Router /ship [post]
func (h *Handler) CreateShip(c *gin.Context) {
    authInstance := auth.GetAuthInstance()
	userID := uint(authInstance.UserID)

    shipName := c.DefaultQuery("shipName", "")

	var ship model.ShipChange

	if err := c.BindJSON(&ship); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "не удалось прочитать JSON"})
		return
	}

    if authInstance.Role == "модератор"{
        err := h.UseCase.CreateShip(userID, ship)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
    
        ships, err := h.UseCase.GetShips(shipName, userID)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
    
        c.JSON(http.StatusOK, gin.H{"ships": ships.Ships, "requestID":ships.RequestID})
    } else {
        c.JSON(http.StatusForbidden, gin.H{"error": "данный запрос доступен только модератору"})
        return
    }
}

// @Summary Удаление судна
// @Description Удаляет судно по его ID
// @Tags Судно
// @Produce json
// @Param shipID path int true "ID судна"
// @Param shipName query string false "Название судна" Format(email)
// @Success 200 {object} model.GetShips "Список судов"
// @Failure 400 {object} model.ErrorResponse "Обработанная ошибка сервера"
// @Failure 401 {object} model.ErrorResponse "Пользователь не авторизован"
// @Failure 403 {object} model.ErrorResponse "У пользователя нет прав для этого запроса"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Security ApiKeyAuth
// @Router /ship/{shipID} [delete]
func (h *Handler) DeleteShip(c *gin.Context) {
    authInstance := auth.GetAuthInstance()
	userID := uint(authInstance.UserID)
	
    shipName := c.DefaultQuery("shipName", "")

	shipID, err := strconv.Atoi(c.Param("shipID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "недопустимый ИД судна"})
		return
	}

    if authInstance.Role == "модератор"{
        err = h.UseCase.DeleteShip(uint(shipID), userID)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
    
        ships, err := h.UseCase.GetShips(shipName, userID)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
    
        c.JSON(http.StatusOK, gin.H{"ships": ships.Ships, "requestID":ships.RequestID})
    } else {
        c.JSON(http.StatusForbidden, gin.H{"error": "данный запрос доступен только модератору"})
        return
    }
}

// @Summary Обновление информации о суднe
// @Description Обновляет информацию о суднe по его ID
// @Tags Судно
// @Accept json
// @Produce json
// @Param shipID path int true "ID судна"
// @Success 200 {object} model.Ship "Информация о суднe"
// @Failure 400 {object} model.ErrorResponse "Обработанная ошибка сервера"
// @Failure 401 {object} model.ErrorResponse "Пользователь не авторизован"
// @Failure 403 {object} model.ErrorResponse "У пользователя нет прав для этого запроса"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Security ApiKeyAuth
// @Router /ship/{shipID} [put]
func (h *Handler) UpdateShip(c *gin.Context) {
    authInstance := auth.GetAuthInstance()
	userID := uint(authInstance.UserID)

    shipID, err := strconv.Atoi(c.Param("shipID"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"error": "недопустимый ИД судна"}})
        return
    }

    var ship model.ShipChange
    if err := c.BindJSON(&ship); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "не удалось прочитать JSON"})
        return
    }
    
    if authInstance.Role == "модератор"{
        err = h.UseCase.UpdateShip(uint(shipID),uint(userID), ship)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        updatedShip, err := h.UseCase.GetShipByID(uint(shipID), uint(userID))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"ship": updatedShip})
    } else {
        c.JSON(http.StatusForbidden, gin.H{"error": "данный запрос дсотупен только модератору"})
        return
    }
}

// @Summary Добавление судна к доставке
// @Description Добавляет судно к доставке по его ID
// @Tags Судно
// @Produce json
// @Param shipID path int true "ID судна"
// @Param shipName query string false "Название судна" Format(email)
// @Success 200 {object} model.GetShips  "Список судов"
// @Failure 400 {object} model.ErrorResponse "Обработанная ошибка сервера"
// @Failure 401 {object} model.ErrorResponse "Пользователь не авторизован"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Security ApiKeyAuth
// @Router /ship/{shipID}/request [post]
func (h *Handler) AddShipToRequest(c *gin.Context) {
	authInstance := auth.GetAuthInstance()
	userID := uint(authInstance.UserID)

    shipName := c.DefaultQuery("shipName", "")

    shipID, err := strconv.Atoi(c.Param("shipID"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "недопустимый ИД судна"})
        return
    }

    err = h.UseCase.AddShipToRequest(uint(shipID), uint(userID))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	ships, err := h.UseCase.GetShips(shipName, uint(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    c.JSON(http.StatusOK, gin.H{"ships": ships.Ships, "requestID":ships.RequestID})
}

// @Summary Удаление судна из доставки
// @Description Удаляет судно из доставки по его ID
// @Tags Судно
// @Produce json
// @Param shipID path int true "ID судна"
// @Param shipName query string false "Название судна" Format(email)
// @Success 200 {object} model.GetShips "Список судов"
// @Failure 400 {object} model.ErrorResponse "Обработанная ошибка сервера"
// @Failure 401 {object} model.ErrorResponse "Пользователь не авторизован"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Security ApiKeyAuth
// @Router /ship/{shipID}/request [delete]
func (h *Handler) RemoveShipFromRequest(c *gin.Context) {
    authInstance := auth.GetAuthInstance()
	userID := uint(authInstance.UserID)

    shipName := c.DefaultQuery("shipName", "")

    shipID, err := strconv.Atoi(c.Param("shipID"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "недопустимый ИД судна"})
        return
    }
   
    err = h.UseCase.RemoveShipFromRequest(uint(shipID), uint(userID))  
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ships, err := h.UseCase.GetShips(shipName, uint(userID))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"ships": ships.Ships, "requestID":ships.RequestID})
}

// @Summary Добавление изображения к судноу
// @Description Добавляет изображение к судноу по его ID
// @Tags Судно
// @Accept mpfd
// @Produce json
// @Param shipID path int true "ID судна"
// @Param image formData file true "Изображение судна"
// @Success 200 {object} model.Ship "Информация о суднe с изображением"
// @Failure 400 {object} model.ErrorResponse "Обработанная ошибка сервера"
// @Failure 401 {object} model.ErrorResponse "Пользователь не авторизован"
// @Failure 403 {object} model.ErrorResponse "У пользователя нет прав для этого запроса"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Security ApiKeyAuth
// @Router /ship/{shipID}/image [post]
func (h* Handler) AddShipImage(c* gin.Context) {
    authInstance := auth.GetAuthInstance()
	userID := uint(authInstance.UserID)

    shipID, err := strconv.Atoi(c.Param("shipID"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "недопустимый ИД судна"})
        return
    }

    image, err := c.FormFile("image")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "недопустимое изображение"})
        return
    }

    file, err := image.Open()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "не удалось открыть изображение"})
        return
    }
    defer file.Close()

    imageBytes, err := io.ReadAll(file)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "не удалось прочитать изображение в байтах"})
        return
    }

	contentType := image.Header.Get("Content-Type")
    
    if authInstance.Role == "модератор"{
        err = h.UseCase.AddShipImage(uint(shipID), uint(userID),imageBytes, contentType)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        ship, err := h.UseCase.GetShipByID(uint(shipID),uint(userID))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"ship": ship})
    } else {
        c.JSON(http.StatusForbidden, gin.H{"error": "данный запрос доступен только модератору"})
        return
    }
}



