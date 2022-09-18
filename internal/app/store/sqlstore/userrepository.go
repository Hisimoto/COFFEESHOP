package sqlstore

import "github.com/Hisimoto/COFFEESHOP/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	return r.store.db.QueryRow(
		"INSERT INTO users (email,encrypted_password) VALUES ($1,$2) RETURNING ID",
		u.Email, u.EncryptedPassword,
	).Scan(&u.ID)
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, email, encrypted_password from users where email = $1", email).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) CountCoffeByTypeAndUserId(id int) (int, error) {
	u := id
	if err := r.store.db.QueryRow("select count(*)  from orders o "+
		"join users u on o.userid = u.id"+
		"join membership m on u.membership_type = m.id "+
		"where u.id = $1 and o.coffee_type = $2 and o.order_date >= (current_timestamp - interval '24 hour')", id).Scan(&u); err != nil {
		return -1, err
	}

	return u, nil
}

func (r *UserRepository) GetNumbersOfCoffeeByTypeAndUserId(id int) (int, error) {
	u := id
	if err := r.store.db.QueryRow("select coffee_amount from membership m"+
		"join users u on m.membership_id=u.membership_type where u.email = 'cristi@gmail.com' and m.coffee_type = 1", id).Scan(&u); err != nil {
		return -1, err
	}

	return u, nil
}
