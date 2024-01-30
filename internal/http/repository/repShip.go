package repository

import (
	"errors"
	"time"

	"backend/internal/model"
)

func (r *Repository) GetShips(shipName string, userID uint) (model.GetShips, error) {
    var request model.Request
    if err := r.db.Table("requests").
        Where("user_id = ? AND request_status = ?", userID, model.REQUEST_STATUS_DRAFT).
        Take(&request).Error; 
        err != nil {}

    var ships []model.Ship
    if err := r.db.Table("ships").
        Where("ships.ship_status = ? AND ships.ship_name LIKE ?", model.SHIP_STATUS_ACTIVE, shipName).
        Scan(&ships).Error; err != nil {
        return model.GetShips{}, errors.New("ошибка нахождения списка судов")
    }

    shipResponse := model.GetShips{
        Ships:   ships,
        RequestID: request.RequestID,
    }

    return shipResponse, nil
}

func (r *Repository) GetShipByID(shipID, userID uint) (model.Ship, error) {
    var ship model.Ship

	if err := r.db.Table("ships").
    Where("ship_status = ? AND ship_id = ?", model.SHIP_STATUS_ACTIVE, shipID).
    First(&ship).Error; 
    err != nil {
		return model.Ship{}, errors.New("ошибка при получении судна из БД")
	}

    return ship, nil
}

func (r *Repository) CreateShip(userID uint, ship model.Ship) error {
	if err := r.db.Create(&ship).Error; err != nil {
		return errors.New("ошибка создания судна")
	}

	return nil
}

func (r *Repository) DeleteShip(shipID, userID uint) error {
    var ship model.Ship

	if err := r.db.Table("ships").Where("ship_id = ? AND ship_status = ?", shipID, model.SHIP_STATUS_ACTIVE).First(&ship).Error; 
    err != nil {
		return errors.New("судно не найдено или уже удалено")
	}

	ship.ShipStatus = model.SHIP_STATUS_DELETED

	if err := r.db.Table("ships").
    Model(&model.Ship{}).
    Where("ship_id = ?", shipID).
    Updates(ship).Error; 
    err != nil {
		return errors.New("ошибка при обновлении статуса судна в БД")
	}
    return nil
}

func (r *Repository) UpdateShip(shipID, userID uint, ship model.Ship) error {
    if err := r.db.Table("ships").
    Model(&model.Ship{}).
    Where("ship_id = ? AND ship_status = ?", shipID, model.SHIP_STATUS_ACTIVE).
    Updates(ship).Error; 
    err != nil {
		return errors.New("ошибка при обновлении информации о судне в БД")
	}

	return nil
}

func (r *Repository) AddShipToRequest(shipID, userID uint) error {
    var ship model.Ship

	if err := r.db.Table("ships").
    Where("ship_id = ? AND ship_status = ?", shipID, model.SHIP_STATUS_ACTIVE).
    First(&ship).Error; 
    err != nil {
		return errors.New("судно не найдено или удалено")
	}

    var request model.Request

    if err := r.db.Table("requests").
    Where("request_status = ? AND user_id = ?", model.REQUEST_STATUS_DRAFT, userID).
    Last(&request).Error; 
    err != nil {
        request = model.Request{
            RequestStatus: model.REQUEST_STATUS_DRAFT,
            CreationDate:   time.Now(),
            UserID:         userID, 
        }

        if err := r.db.Table("requests").
        Create(&request).Error;
        err != nil {
            return errors.New("ошибка создания заявки со статусом черновик")
        }
    }

    requestShip := model.RequestShip{
        ShipID:  ship.ShipID,
        RequestID: request.RequestID,
    }

    if err := r.db.Table("request_ships").
    Create(&requestShip).Error;
    err != nil {
		return errors.New("ошибка при создании связи между заявкой и судном")
	}

    return nil
}

func (r *Repository) RemoveShipFromRequest(shipID, userID uint) error {
    var requestShip model.RequestShip

    if err := r.db.Joins("JOIN requests ON request_ships.request_id = requests.request_id").
        Where("request_ships.ship_id = ? AND requests.user_id = ? AND requests.request_status = ?", shipID, userID, model.REQUEST_STATUS_DRAFT).
        First(&requestShip).Error; 
        err != nil {
        return errors.New("судно не принадлежит пользователю или находится не в статусе черновик")
    }

    if err := r.db.Table("request_ships").
    Delete(&requestShip).Error; 
    err != nil {
        return errors.New("ошибка удаления связи между заявкой и судном")
    }
 
    return nil
}

func (r *Repository) AddShipImage(shipID, userID uint, imageURL string) error {
    err := r.db.Table("ships").Where("ship_id = ?", shipID).Update("photo", imageURL).Error
    if err != nil {
        return errors.New("ошибка обновления url изображения в БД")
    }

    return nil
}




