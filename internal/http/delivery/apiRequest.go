package delivery

import (
	"net/http"
	"strconv"

	"backend/internal/auth"
	"backend/internal/model"

	"github.com/gin-gonic/gin"
)

// GetRequests godoc
// @Summary Получение списка заявок
// @Description Возвращает список всех не удаленных заявок
// @Tags Заявка
// @Produce json
// @Param startFormationDate query string false "Начало даты формирования" Format(email)
// @Param endFormationDate query string false "Конец даты формирования" Format(email)
// @Param requestStatus query string false "Статус заявки" Format(email)
// @Success 200 {object} model.GetRequests "Список заявок"
// @Failure 400 {object} model.ErrorResponse "Обработанная ошибка сервера"
// @Failure 401 {object} model.ErrorResponse "Пользователь не авторизован"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Security ApiKeyAuth
// @Router /request [get]
func (h *Handler) GetRequests(c *gin.Context) {
    authInstance := auth.GetAuthInstance()
	userID := uint(authInstance.UserID)

    startFormationDate := c.DefaultQuery("startFormationDate", "")
    endFormationDate := c.DefaultQuery("endFormationDate", "")
    requestStatus := c.DefaultQuery("requestStatus", "")

    var requests []model.GetRequests
    var err error

    if authInstance.Role == "модератор"{
        requests, err = h.UseCase.GetRequestsModerator(startFormationDate, endFormationDate, requestStatus)  
    } else {
        requests, err = h.UseCase.GetRequestsUser(startFormationDate, endFormationDate, requestStatus, userID)
    }
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"requests": requests})
}

// GetRequestByID godoc
// @Summary Получение заявки по идентификатору
// @Description Возвращает информацию о заявке по её идентификатору
// @Tags Заявка
// @Produce json
// @Param requestID path int true "Идентификатор заявки"
// @Success 200 {object} model.GetRequestByID "Информация о заявке"
// @Failure 400 {object} model.ErrorResponse "Обработанная ошибка сервера"
// @Failure 401 {object} model.ErrorResponse "Пользователь не авторизован"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Security ApiKeyAuth
// @Router /request/{requestID} [get]
func (h *Handler) GetRequestByID(c *gin.Context) {
    authInstance := auth.GetAuthInstance()
	userID := uint(authInstance.UserID)

    requestID, err := strconv.Atoi(c.Param("requestID"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "недопустимый ИД заявки"})
        return
    }

    var request model.GetRequestByID

    if authInstance.Role == "модератор"{
        request, err = h.UseCase.GetRequestByIDModerator(uint(requestID)) 
    } else {
        request, err = h.UseCase.GetRequestByIDUser(uint(requestID), userID)
    }
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"request" : request})
}

// DeleteRequest godoc
// @Summary Удаление заявки
// @Description Удаляет заявку по её идентификатору
// @Tags Заявка
// @Produce json
// @Param requestID path int true "Идентификатор заявки"
// @Param searchFlightNumber query string false "Номер рейса" Format(email)
// @Param startFormationDate query string false "Начало даты формирования" Format(email)
// @Param endFormationDate query string false "Конец даты формирования" Format(email)
// @Param requestStatus query string false "Статус заявки" Format(email)
// @Success 200 {object} model.GetRequests "Список заявок"
// @Failure 400 {object} model.ErrorResponse "Обработанная ошибка сервера"
// @Failure 401 {object} model.ErrorResponse "Пользователь не авторизован"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Security ApiKeyAuth
// @Router /request/{requestID} [delete]
func (h *Handler) DeleteRequest(c *gin.Context) {
    authInstance := auth.GetAuthInstance()
	userID := uint(authInstance.UserID)

    startFormationDate := c.DefaultQuery("startFormationDate", "")
    endFormationDate := c.DefaultQuery("endFormationDate", "")
    requestStatus := c.DefaultQuery("requestStatus", "")
    requestID, err := strconv.Atoi(c.Param("requestID"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "недопустимый ИД заявки"})
        return
    }

    if authInstance.Role == "модератор"{
        err = h.UseCase.DeleteRequestUser(uint(requestID), userID)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        requests, err := h.UseCase.GetRequestsModerator(startFormationDate, endFormationDate, requestStatus)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"requests": requests})
    } else {
        err = h.UseCase.DeleteRequestUser(uint(requestID), userID)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        requests, err := h.UseCase.GetRequestsUser(startFormationDate, endFormationDate, requestStatus, userID)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"requests": requests})
    }
}

// UpdateRequestStatusUser godoc
// @Summary Обновление статуса заявки для пользователя
// @Description Обновляет статус заявки для пользователя по идентификатору заявки
// @Tags Заявка
// @Produce json
// @Param requestID path int true "Идентификатор заявки"
// @Success 200 {object} model.GetRequestByID "Информация о заявке"
// @Failure 400 {object} model.ErrorResponse "Обработанная ошибка сервера"
// @Failure 401 {object} model.ErrorResponse "Пользователь не авторизован"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Security ApiKeyAuth
// @Router /request/{requestID}/status/user [put]
func (h *Handler) UpdateRequestStatusUser(c *gin.Context) {
    authInstance := auth.GetAuthInstance()
	userID := uint(authInstance.UserID)

    requestID, err := strconv.Atoi(c.Param("requestID"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "недоупстимый ИД заявки"})
        return
    }

    if authInstance.Role == "модератор" {
        err = h.UseCase.UpdateRequestStatusUser(uint(requestID), userID)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        request, err := h.UseCase.GetRequestByIDModerator(uint(requestID))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
        }

        c.JSON(http.StatusOK, gin.H{"request": request})
    
    } else {
        err = h.UseCase.UpdateRequestStatusUser(uint(requestID), userID)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        request, err := h.UseCase.GetRequestByIDUser(uint(requestID), userID)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
        }

        c.JSON(http.StatusOK, gin.H{"request": request})
    }
}

// UpdateRequestStatusModerator godoc
// @Summary Обновление статуса заявки для модератора
// @Description Обновляет статус заявки для модератора по идентификатору заявки
// @Tags Заявка
// @Produce json
// @Param requestID path int true "Идентификатор заявки"
// @Param requestStatus body model.RequestUpdateStatusRequest true "Новый статус заявки"
// @Success 200 {object} model.GetRequestByID "Информация о заявке"
// @Failure 400 {object} model.ErrorResponse "Обработанная ошибка сервера"
// @Failure 401 {object} model.ErrorResponse "Пользователь не авторизован"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Security ApiKeyAuth
// @Router /request/{requestID}/status/moderator [put]
func (h *Handler) UpdateRequestStatusModerator(c *gin.Context) {
    authInstance := auth.GetAuthInstance()
	userID := uint(authInstance.UserID)

    requestID, err := strconv.Atoi(c.Param("requestID"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "недопустимый ИД заявки"})
        return
    }

    var requestStatus model.StatusRequest
    if err := c.BindJSON(&requestStatus); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if authInstance.Role == "модератор" {
        err = h.UseCase.UpdateRequestStatusModerator(uint(requestID), userID, requestStatus)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        request, err := h.UseCase.GetRequestByIDModerator(uint(requestID))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
        }

        c.JSON(http.StatusOK, gin.H{"request": request})
    } else {
        c.JSON(http.StatusForbidden, gin.H{"error": "данный запрос доступен только модератору"})
        return
    }
}
