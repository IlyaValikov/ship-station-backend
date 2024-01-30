package usecase

import (
	"errors"
	"strings"

	"github.com/markgregr/RIP/internal/model"
)

func (uc *UseCase) GetShips(shipName string, userID uint) (model.GetShips, error) {
	if userID < 0 {
		return model.GetShips{}, errors.New("недопустимый ИД пользователя")
	}

	shipName = strings.ToLower(shipName + "%")

	ships, err := uc.Repository.GetShips(shipName, userID)
	if err != nil {
		return model.GetShips{}, err
	}

	return ships, nil
}

func (uc *UseCase) GetShipByID(shipID, userID uint) (model.Ship, error) {
	if shipID < 0 {
		return model.Ship{}, errors.New("недопустимый ИД судна")
	}
	if userID < 0 {
		return model.Ship{}, errors.New("недопустимый ИД пользователя")
	}

	ship, err := uc.Repository.GetShipByID(shipID, userID)
	if err != nil {
		return model.Ship{}, err
	}

	return ship, nil
}

func (uc *UseCase) CreateShip(userID uint, shipUpdate model.ShipChange) error {
	if userID <= 0 {
		return errors.New("недопустимый ИД пользователя")
	}
	if shipUpdate.ShipName == "" {
		return errors.New("название судна должно быть заполнено")
	}
	if shipUpdate.CargoCapacity == 0 {
		return errors.New("грузоподъемность судна должна быть заполнена")
	}
	if shipUpdate.MaxDepth == 0 {
		return errors.New("максимальная глубина судна должна быть заполнена")
	}
	if shipUpdate.MaxLength == 0 {
		return errors.New("максимальная длина судна должна быть заполнена")
	}
	if shipUpdate.ShipType == "" {
		return errors.New("тип судна должен быть заполнен")
	}
	if shipUpdate.Flag == "" {
		return errors.New("флаг судна должен быть заполнен")
	}
	if shipUpdate.Classification == "" {
		return errors.New("классификация судна должна быть заполнена")
	}

	ship := model.Ship{
		ShipName:          shipUpdate.ShipName,
		ShipType:          shipUpdate.ShipType,
		CargoCapacity:     shipUpdate.CargoCapacity,
		MaxDepth:          shipUpdate.MaxDepth,
		MaxLength:         shipUpdate.MaxLength,
		YearBuilt:         shipUpdate.YearBuilt,
		Flag:              shipUpdate.Flag,
		Classification:    shipUpdate.Classification,
		CrewCapacity:      shipUpdate.CrewCapacity,
		PassengerCapacity: shipUpdate.PassengerCapacity,
		ShipStatus:        model.SHIP_STATUS_ACTIVE,
	}

	err := uc.Repository.CreateShip(userID, ship)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) DeleteShip(shipID, userID uint) error {
	if shipID < 0 {
		return errors.New("недопустимый ИД судна")
	}
	if userID < 0 {
		return errors.New("недопустимый ИД пользователя")
	}

	err := uc.Repository.DeleteShip(shipID, userID)
	if err != nil {
		return err
	}

	err = uc.Repository.RemoveServiceImage(shipID, userID)
	if err != nil {
		return err
	}
	
	return nil
}

func (uc *UseCase) UpdateShip(shipID, userID uint, shipUpdate model.ShipChange) error {
	if shipID < 0 {
		return errors.New("недопустимый ИД судна")
	}
	if userID < 0 {
		return errors.New("недопустимый ИД пользователя")
	}

	ship := model.Ship{
		ShipName:          shipUpdate.ShipName,
		ShipType:          shipUpdate.ShipType,
		CargoCapacity:     shipUpdate.CargoCapacity,
		MaxDepth:          shipUpdate.MaxDepth,
		MaxLength:         shipUpdate.MaxLength,
		YearBuilt:         shipUpdate.YearBuilt,
		Flag:              shipUpdate.Flag,
		Classification:    shipUpdate.Classification,
		CrewCapacity:      shipUpdate.CrewCapacity,
		PassengerCapacity: shipUpdate.PassengerCapacity,
	}

	err := uc.Repository.UpdateShip(shipID, userID, ship)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) AddShipToRequest(shipID, userID uint) error {
	if shipID < 0 {
		return errors.New("недопустимый ИД судна")
	}
	if userID < 0 {
		return errors.New("недопустимый ИД пользователя")
	}

	err := uc.Repository.AddShipToRequest(shipID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) RemoveShipFromRequest(shipID, userID uint) error {
	if shipID < 0 {
		return errors.New("недопустимый ИД судна")
	}
	if userID < 0 {
		return errors.New("недопустимый ИД пользователя")
	}

	err := uc.Repository.RemoveShipFromRequest(shipID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) AddShipImage(shipID, userID uint, imageBytes []byte, ContentType string) error {
	if shipID < 0 {
		return errors.New("недопустимый ИД судна")
	}
	if userID < 0 {
		return errors.New("недопустимый ИД пользователя")
	}
	if imageBytes == nil {
		return errors.New("недопустимый imageBytes изображения")
	}
	if ContentType == "" {
		return errors.New("недопустимый ContentType изображения")
	}

	imageURL, err := uc.Repository.UploadServiceImage(shipID, userID, imageBytes, ContentType)
	if err != nil {
		return err
	}

	err = uc.Repository.AddShipImage(shipID, userID, imageURL)
	if err != nil {
		return err
	}

	return nil
}



