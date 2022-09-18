package store

import "github.com/Hisimoto/COFFEESHOP/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	CountCoffeByTypeAndUserId(int) (int, error)
	GetNumbersOfCoffeeByTypeAndUserId(int) (int, error)
}
