package model

import "time"

type Request struct {
	RequestID      uint       `gorm:"type:serial;primarykey" json:"request_id"`
	CreationDate   time.Time  `json:"creation_date"`
	FormationDate  *time.Time `json:"formation_date"`
	CompletionDate *time.Time `json:"completion_date"`
	RequestStatus  string     `json:"request_status"`
	UserID         uint       `json:"user_id"`
	ModeratorID    *uint       `json:"moderator_id"`
}

type GetRequests struct {
	RequestID      uint      `json:"request_id"`
	CreationDate   time.Time `json:"creation_date"`
	FormationDate  *time.Time `json:"formation_date"`
	CompletionDate *time.Time `json:"completion_date"`
	RequestStatus  string    `json:"request_status"`
	FullName 	   string 	 `json:"full_name"`
}

type GetRequestByID struct{
	RequestID 	   uint       `json:"request_id"`
	CreationDate   time.Time  `json:"creation_date"`
	FormationDate  *time.Time `json:"formation_date"`
	CompletionDate *time.Time `json:"completion_date"`
	RequestStatus  string     `json:"request_status"`
	FullName 	   string 	  `json:"full_name"`
	Ships          []Ship    `gorm:"many2many:request_ships" json:"ships"`
}

type StatusRequest struct {
	RequestStatus string    `json:"request_status"`
}