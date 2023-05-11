package dao

import (
	"github.com/quniverse26/miniproject/model"
)

type AdminDaoInterface interface {
	GetAdmin(id int) (model.Admin, error)
}