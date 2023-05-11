package dao

import (
	"github.com/quniverse26/miniproject/model"
)

type BuyerDaoInterface interface {
	GetBuyer(id int) (model.Buyer, error)
}