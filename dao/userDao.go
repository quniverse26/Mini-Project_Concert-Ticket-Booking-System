package dao

import (
	"github.com/quniverse26/miniproject/model"
)

type UserDaoInterface interface {
	GetUser(id int) (model.User, error)
}