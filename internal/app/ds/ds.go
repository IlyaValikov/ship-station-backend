package api

import "time"

const (
	SHIP_STATUS_ACTIVE  = "активен"
	SHIP_STATUS_DELETED = "удален"
)

type Ship struct {
	ShipID            uint    `gorm:"primarykey" json:"ship_id"`
	ShipName          string  `json:"ship_name"`
	ShipType          string  `json:"ship_type"`
	CargoCapacity     float64 `json:"cargo_capacity"`
	MaxDepth          float64 `json:"max_depth"`
	MaxLength         float64 `json:"max_length"`
	YearBuilt         int     `json:"year_built"`
	Flag              string  `json:"flag"`
	Classification    string  `json:"classification"`
	CrewCapacity      int     `json:"crew_capacity"`
	PassengerCapacity int     `json:"passenger_capacity"`
	Photo             string  `json:"photo"`
	Href              string  `json:"href"`
}

type Request struct {
	RequestID       uint      `gorm:"primarykey" json:"request_id"`
	CreationDate    time.Time `json:"creation_date"`
	FormationDate   *time.Time `json:"formation_date"`
	CompletionDate  *time.Time `json:"completion_date"`
	RequestStatus   string    `json:"request_status"`
	UserID          uint      `json:"user_id"`
	ModeratorID     *uint      `json:"moderator_id"`
}

type RequestShip struct {
	RequestShipID uint `gorm:"primarykey" json:"request_ship_id"`
	RequestID     uint `json:"request_id"`
	ShipID        uint `json:"ship_id"`
}

type User struct {
	UserID   uint   `gorm:"primarykey" json:"user_id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
