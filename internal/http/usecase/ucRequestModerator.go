package usecase

import (
	"errors"
	"strings"

	"github.com/markgregr/RIP/internal/model"
)

func (uc *UseCase) GetRequestsModerator(startFormationDate, endFormationDate, requestsStatus string) ([]model.GetRequests, error) {
	requestsStatus = strings.ToLower(requestsStatus + "%")

	requests, err := uc.Repository.GetRequestsModerator(startFormationDate, endFormationDate, requestsStatus)
	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (uc *UseCase) GetRequestByIDModerator(requestsID uint) (model.GetRequestByID, error) {
	if requestsID < 0 {
		return model.GetRequestByID{}, errors.New("недопустимый ИД заявки")
	}

	requests, err := uc.Repository.GetRequestByIDModerator(requestsID)
	if err != nil {
		return model.GetRequestByID{}, err
	}

	return requests, nil
}

func (uc *UseCase) UpdateRequestStatusModerator(requestsID, moderatorID uint, requestsStatus model.StatusRequest) error{
	if requestsID < 0 {
		return errors.New("недопустимый ИД зявки")
	}
	if requestsStatus.RequestStatus != model.REQUEST_STATUS_COMPLETED && requestsStatus.RequestStatus != model.REQUEST_STATUS_REJECTED {
		return errors.New("текущий статус заявки уже завершен или отклонен")
	}

	err := uc.Repository.UpdateRequestStatusModerator(requestsID, moderatorID, requestsStatus)
	if err != nil {
		return err
	}

	return nil
}


