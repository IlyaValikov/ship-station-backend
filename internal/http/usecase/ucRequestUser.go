package usecase

import (
	"errors"
	"strings"

	"backend/internal/model"
)

func (uc *UseCase) GetRequestsUser(startFormationDate, endFormationDate, requestStatus string, userID uint) ([]model.GetRequests, error) {
	requestStatus = strings.ToLower(requestStatus + "%")

	if userID <= 0 {
		return nil, errors.New("недопустимый ИД пользователя")
	}

	requests, err := uc.Repository.GetRequestsUser(startFormationDate, endFormationDate, requestStatus, userID)
	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (uc *UseCase) GetRequestByIDUser(requestID, userID uint) (model.GetRequestByID, error) {
	if requestID < 0 {
		return model.GetRequestByID{}, errors.New("недопустимый ИД заявки")
	}
	if userID < 0 {
		return model.GetRequestByID{}, errors.New("недопустимый ИД пользователя")
	}

	requests, err := uc.Repository.GetRequestByIDUser(requestID, userID)
	if err != nil {
		return model.GetRequestByID{}, err
	}

	return requests, nil
}

func (uc *UseCase) DeleteRequestUser(requestID, userID uint) error{
	if requestID < 0 {
		return errors.New("недопустимый ИД заявки")
	}
	if userID < 0 {
		return errors.New("недопустимый ИД пользователя")
	}

	err := uc.Repository.DeleteRequestUser(requestID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) UpdateRequestStatusUser(requestID, userID uint, check bool) error{
	if requestID < 0 {
		return errors.New("недопустимый ИД заявки")
	}
	if userID < 0 {
		return errors.New("недопустимый ИД пользователя")
	}

	err := uc.Repository.UpdateRequestStatusUser(requestID, userID, check)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) CheckRequestUser(requestID uint, token string) error{
	err := uc.Repository.CheckRequestUser(requestID, token)
	if err != nil {
		return err
	}

	return nil
}


