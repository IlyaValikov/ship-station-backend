package model

type Ship struct {
	ShipID            uint    `gorm:"type:serial;primarykey" json:"ship_id"` //есть
	ShipName          string  `json:"ship_name"`                             //yes
	ShipType          string  `json:"ship_type"`                             //yes
	ShipStatus        string  `json:"ship_status"`                           //yes
	CargoCapacity     float64 `json:"cargo_capacity"`                        //yes
	MaxDepth          float64 `json:"max_depth"`
	MaxLength         float64 `json:"max_length"`
	YearBuilt         int     `json:"year_built"`
	Flag              string  `json:"flag"`               //yes
	Classification    string  `json:"classification"`     //yes
	CrewCapacity      int     `json:"crew_capacity"`      //yes
	PassengerCapacity int     `json:"passenger_capacity"` //yes
	Photo             string  `json:"photo"`              //yes
}

type ShipChange struct {
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
}

type GetShips struct {
	Ships     []Ship `json:"ships"`
	RequestID uint   `json:"request_id"`
}
