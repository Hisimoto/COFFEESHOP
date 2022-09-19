package sqlstore

import (
	"github.com/Hisimoto/COFFEESHOP/internal/app/model"
)

type OrderRepository struct {
	store *Store
}

func (r *OrderRepository) CountCoffeByTypeAndUserId(id int, coffeType int) (int, error) {
	var u int
	if err := r.store.db.QueryRow("select count(*)  from orders o "+
		"join users u on o.userid = u.id "+
		"join membership m on u.membership_type = m.id "+
		"where u.id = $1 and o.coffee_type = $2 and o.order_date >= (current_timestamp - interval '1 hour' * m.timeframe)", id, coffeType).Scan(&u); err != nil {
		return -1, err
	}

	return u, nil
}

func (r *OrderRepository) CreateOrder(u *model.Order, coffeType int) error {
	return r.store.db.QueryRow(
		"INSERT INTO orders (userid,coffee_type, order_date) VALUES ($1,$2, current_timestamp) RETURNING ID",
		u.ID, coffeType).Scan(&u.ID)
}

func (r *OrderRepository) CheckRemainingTime(u int, coffeType int) (string, error) {
	var u1 string
	err := r.store.db.QueryRow(
		"select (order_date+interval '1 hour' * m.timeframe) - current_timestamp from orders o "+
			"join users u on o.userid = u.id "+
			"join membership m on u.membership_type = m.id "+
			"where u.id = $1 and o.coffee_type = $2 and o.order_date >= (current_timestamp - interval '1 hour' * m.timeframe) "+
			"order by order_date asc "+
			"limit 1",
		u, coffeType).Scan(&u1)
	if err != nil {
		return u1, err
	}
	return u1, err
}
