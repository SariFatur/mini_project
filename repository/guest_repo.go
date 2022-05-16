package repository

import "myproject/model"

type IGuestRepo interface {
	GetGuestController() ([]model.GuestCustom, error)
	GetGuestByIdController(id int) (model.Guest, error)
	CreateGuestController(guest model.Guest) error
}
