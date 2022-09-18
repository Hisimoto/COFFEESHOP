package sqlstore

import (
	"github.com/Hisimoto/COFFEESHOP/internal/app/model"
	"time"
)

type OrderRepository struct {
	store *Store
}

func (r *OrderRepository) CountCoffeByTypeAndUserId(id int, coffeType int) (int, error) {
	var u int
	if err := r.store.db.QueryRow("select count(*)  from orders o "+
		"join users u on o.userid = u.id"+
		"join membership m on u.membership_type = m.id "+
		"where u.id = $1 and o.coffee_type = $2 and o.order_date >= (current_timestamp - interval '1 hour' * m.timeframe)", id, coffeType).Scan(&u); err != nil {
		return -1, err
	}

	return u, nil
}

func (r *OrderRepository) CreateOrder(u *model.Order, coffeType int) error {
	return r.store.db.QueryRow(
		"INSERT INTO orders (userid,coffee_type, order_date) VALUES ($1,$2, current_timestamp)",
		u.ID, coffeType).Scan(&u.ID)
}

func (r *OrderRepository) CheckRemainingTime(u int, coffeType int) (timeLeft time.Time, err error) {
	u1 := timeLeft
	err = r.store.db.QueryRow(
		"select current_timestamp - min(order_date)  from orders o "+
			"join users u on o.userid = u.id"+
			"join membership m on u.membership_type = m.id "+
			"where u.id = $1 and o.coffee_type = $2 and o.order_date >= (current_timestamp - interval '1 hour' * m.timeframe)",
		u, coffeType).Scan(&u1)
	if err != nil {
		return u1, err
	}
	return u1, err
}
