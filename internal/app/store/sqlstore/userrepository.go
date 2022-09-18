package sqlstore

import "github.com/Hisimoto/COFFEESHOP/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	return r.store.db.QueryRow(
		"INSERT INTO users (email,membership_type) VALUES ($1,$2) RETURNING ID",
		u.Email, u.MembershipType,
	).Scan(&u.ID)
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, email, membership_type "+
		"from users where email = $1", email).Scan(&u.ID, &u.Email, &u.MembershipType); err != nil {
		return nil, err
	}

	return u, nil
}
func (r *UserRepository) GetNumbersOfCoffeeByTypeAndUserId(id int, coffeType int) (int, error) {
	var u int
	if err := r.store.db.QueryRow("select coffee_amount from membership m "+
		"join users u on m.membership_id=u.membership_type where u.id = $1 and m.coffee_type = $2", id, coffeType).Scan(&u); err != nil {
		return -1, err
	}

	return u, nil
}
