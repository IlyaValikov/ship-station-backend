package repository

import (
	"backend/internal/app/ds"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetShipByID(shipID uint) (*ds.Ship, error) {
	ship := &ds.Ship{}
	err := r.db.First(ship, "ship_id = ? AND ship_status = ?", shipID, ds.SHIP_STATUS_ACTIVE).Error
	if err != nil {
		return nil, err
	}
	return ship, nil
}

	func (r *Repository) GetShips(shipName string) ([]ds.Ship,error) {
		shipName = strings.ToLower(shipName+"%")
		var ships []ds.Ship
		if err := r.db.Find(&ships, "ship_status = ? AND LOWER(ship_name) LIKE ?", ds.SHIP_STATUS_ACTIVE, shipName).Error; err != nil {
			return nil, err
		}
		return ships, nil
	}

func (r *Repository) DeleteShip(shipID uint) error {
	return r.db.Exec("UPDATE ships SET ship_status = ? WHERE ship_id = ?", ds.SHIP_STATUS_DELETED, shipID).Error
}
