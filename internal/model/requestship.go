package model

type RequestShip struct {
	RequestID uint `gorm:"type:serial;primaryKey;index" json:"request_id"`
	ShipID    uint `gorm:"type:serial;primaryKey;index" json:"ship_id"`
}