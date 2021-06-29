package entity

import "time"

type ResponseUsers struct {
	Status string   `json:"status"`
	Data   []*Users `json:"data"`
}

type Users struct {
	Id        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	CreatedOn time.Time `json:"createdOn"`
	UpdatedOn time.Time `json:"updatedOn"`
	DeletedOn time.Time `json:"deletedOn"`
}
