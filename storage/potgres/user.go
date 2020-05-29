package potgres

import (
	"github.com/buy_event/storage/repo"
	"github.com/jmoiron/sqlx"
	"github.com/google/uuid"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) repo.UserStorageI {
	return &userRepo{db:db}
}

func (ur *userRepo) Create(user *repo.User) error {
	//generate new user_id
	id, err := uuid.NewRandom()

	if err != nil {
		return err
	}

	_, err = ur.db.Exec(`insert into users (id, email, phone_number) values ($1, $2, $3)`, id, user.Email, user.PhoneNumber)

	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepo) Get(id uuid.UUID) (*repo.User, error) {
	var user repo.User
	row := ur.db.QueryRow(`select id, email, phone_number from users where id=$1`, id)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PhoneNumber,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepo) GetAll() ([]*repo.User, error) {
	var users []*repo.User

	rows, err := ur.db.Queryx(`select id, email, phone_number from users`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user repo.User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.PhoneNumber,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (ur *userRepo) Delete(id uuid.UUID) error {
	_, err := ur.db.Exec(`delete from users where id=$1`, id)

	if err != nil {
		return err
	}
	return nil
}

