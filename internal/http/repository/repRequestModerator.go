package repository

import (
	"errors"
	"time"

	"github.com/markgregr/RIP/internal/model"
)


func (r *Repository) GetRequestsModerator(startFormationDate, endFormationDate, requestStatus string) ([]model.GetRequests, error) {
    query := r.db.Table("requests").
        Select("requests.request_id, requests.creation_date, requests.formation_date, equests.completion_date, requests.request_status, creator.full_name, moderator.full_name as moderator_name").
        Joins("JOIN users creator ON creator.user_id = requests.user_id").
        Joins("LEFT JOIN users moderator ON moderator.user_id = requests.moderator_id").
        Where("requests.request_status LIKE ? AND requests.flight_number LIKE ? AND requests.request_status != ? AND requests.request_status != ?", requestStatus,  model.REQUEST_STATUS_DELETED, model.REQUEST_STATUS_DRAFT)
    
    if startFormationDate != "" && endFormationDate != "" {
        query = query.Where("requests.formation_date BETWEEN ? AND ?", startFormationDate, endFormationDate)
    }
// Сортировка по дате формирования в порядке убывания
    query = query.Order("requests.formation_date DESC")

    var requests []model.GetRequests
    if err := query.Find(&requests).Error; err != nil {
    return nil, errors.New("ошибка получения заявок")
    }

    return requests, nil
}

func (r *Repository) GetRequestByIDModerator(requestID uint) (model.GetRequestByID, error) {
    var request model.GetRequestByID

    if err := r.db.
        Table("requests").
        Select("requests.request_id, requests.creation_date, requests.formation_date, requests.completion_date, requests.request_status, users.full_name").
        Joins("JOIN users ON users.user_id = requests.user_id").
        Where("requests.request_status != ? AND requests.request_id = ?", model.REQUEST_STATUS_DELETED, requestID).
        First(&request).Error; err != nil {
        return model.GetRequestByID{}, errors.New("ошибка получения заявки по ИД")
    }

    var ships []model.Ship
    if err := r.db.
        Table("ships").
        Joins("JOIN request_ships ON ships.ship_id = request_ships.ship_id").
        Where("request_ships.request_id = ?", request.RequestID).
        Scan(&ships).Error; err != nil {
        return model.GetRequestByID{}, errors.New("ошибка получения судов заявки")
    }

    request.Ships = ships

    return request, nil
}

func (r *Repository) UpdateRequestStatusModerator(requestID, moderatorID uint, requestStatus model.StatusRequest) error {
    var request model.Request
    if err := r.db.Table("requests").
        Where("request_id = ? AND request_status = ?", requestID, model.REQUEST_STATUS_WORK).
        First(&request).
        Error; err != nil {
        return errors.New("заявка не найдена или не принадлежит указанному модератору")
    }

    request.RequestStatus = requestStatus.RequestStatus
    request.ModeratorID = &moderatorID
    currentTime := time.Now()
	request.CompletionDate = &currentTime

    if err := r.db.Save(&request).Error; err != nil {
        return errors.New("ошибка обновления статуса заявки в БД")
    }

    return nil
}