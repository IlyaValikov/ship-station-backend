package repository

import (
	"errors"
	"time"

	"github.com/markgregr/RIP/internal/model"
)

func (r *Repository) GetRequestsUser(startFormationDate, endFormationDate, requestStatus string, userID uint) ([]model.GetRequests, error) {
    query := r.db.Table("requests").
        Select("requests.request_id, requests.creation_date, requests.formation_date, requests.completion_date, requests.request_status, users.full_name").
        Joins("JOIN users ON users.user_id = requests.user_id").
        Where("requests.request_status LIKE ? AND requests.user_id = ? AND requests.request_status != ? AND requests.request_status != ?", requestStatus, userID, model.REQUEST_STATUS_DELETED, model.REQUEST_STATUS_DRAFT)
    
    if startFormationDate != "" && endFormationDate != "" {
        query = query.Where("requests.formation_date BETWEEN ? AND ?", startFormationDate, endFormationDate)
    }

   query = query.Order("requests.formation_date DESC")

    var requests []model.GetRequests
    if err := query.Find(&requests).Error; err != nil {
        return nil, errors.New("ошибка получения заявок")
    }

    return requests, nil
}

func (r *Repository) GetRequestByIDUser(requestID, userID uint) (model.GetRequestByID, error) {
    var request model.GetRequestByID

    if err := r.db.
        Table("requests").
        Select("requests.request_id, requests.creation_date, requests.formation_date, requests.completion_date, requests.request_status, users.full_name").
        Joins("JOIN users ON users.user_id = requests.user_id").
        Where("requests.request_status != ? AND requests.request_id = ? AND requests.user_id = ?", model.REQUEST_STATUS_DELETED, requestID, userID).
        First(&request).Error; err != nil {
        return model.GetRequestByID{}, errors.New("ошибка получения заявки по ИД")
    }
    

    var ships []model.Ship
    if err := r.db.
        Table("ships").
        Joins("JOIN request_ships ON ships.ship_id = request_ships.ship_id").
        Where("request_ships.request_id = ? AND ships.ship_status != ?", request.RequestID, model.SHIP_STATUS_DELETED).
        Scan(&ships).Error; err != nil {
        return model.GetRequestByID{}, errors.New("ошибка получения судов заявки")
    }

    request.Ships = ships

    return request, nil
}

func (r *Repository) DeleteRequestUser(requestID, userID uint) error {
    var request model.Request
    if err := r.db.Table("requests").
        Where("request_id = ? AND user_id = ? AND request_status = ?", requestID, userID, model.REQUEST_STATUS_DRAFT).
        First(&request).
        Error; err != nil {
        return errors.New("заявка не найдена, или не принадлежит указанному пользователю, или не находится в статусе черновик")
    }

    err := r.db.Model(&model.Request{}).Where("request_id = ?", requestID).Update("request_status", model.REQUEST_STATUS_DELETED).Error
    if err != nil {
        return errors.New("ошибка обновления статуса на удален")
    }
     
    return nil
}

func (r *Repository) UpdateRequestStatusUser(requestID, userID uint) error {
    var request model.Request
    if err := r.db.Table("requests").
        Where("requests.request_id = ? AND requests.user_id = ? AND requests.request_status = ?", requestID, userID, model.REQUEST_STATUS_DRAFT).
        First(&request).
        Error; err != nil {
        return errors.New("заявка не найдена, или не принадлежит указанному пользователю, или не имеет статус черновик")
    }

    request.RequestStatus = model.REQUEST_STATUS_WORK
    currentTime := time.Now()
	request.FormationDate = &currentTime

    if err := r.db.Save(&request).Error; err != nil {
        return errors.New("ошибка обновления статуса заявки в БД")
    }

    return nil
}







