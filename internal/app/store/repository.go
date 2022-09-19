package store

import (
	"github.com/Hisimoto/COFFEESHOP/internal/app/model"
)

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	GetNumbersOfCoffeeByTypeAndUserId(int, int) (int, error)
}

type OrderRepository interface {
	CreateOrder(*model.Order, int) error
	CountCoffeByTypeAndUserId(int, int) (int, error)
	CheckRemainingTime(int, int) (string, error)
}
