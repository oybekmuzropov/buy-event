package potgres

import (
	"github.com/buy_event/storage/repo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type purchaseRepo struct {
	db *sqlx.DB
}

func NewPurchaseRepo(db *sqlx.DB) repo.PurchaseStorageI {
	return &purchaseRepo{db: db}
}

func (gr *purchaseRepo) Create(purchase *repo.Purchase) (string, error) {
	//generate purchase id
	guid, err := uuid.NewRandom()

	if err != nil {
		return "", err
	}

	_, err = gr.db.Exec(`insert into purchases(id, user_id, goods, total_price) values ($1, $2, $3, $4)`,
		guid,
		purchase.UserID,
		purchase.Goods,
		purchase.Price,
	)

	if err != nil {
		return "", err
	}

	return guid.String(), nil
}

func (gr *purchaseRepo) Get(id uuid.UUID) (*repo.Purchase, error) {
	var purchase repo.Purchase

	row := gr.db.QueryRow(`select id, goods, total_price, user_id from purchases where id=$1`, id)

	err := row.Scan(
		&purchase.ID,
		&purchase.Goods,
		&purchase.Price,
		&purchase.UserID,
	)

	if err != nil {
		return nil, err
	}

	return &purchase, nil
}

func (gr *purchaseRepo) GetAll() ([]*repo.Purchase, error) {
	var purchases []*repo.Purchase

	rows, err := gr.db.Queryx(`select id, goods, total_price, user_id from purchases`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var purchase repo.Purchase

		err := rows.Scan(
			&purchase.ID,
			&purchase.Goods,
			&purchase.Price,
			&purchase.UserID,
		)

		if err != nil {
			return nil, err
		}

		purchases = append(purchases, &purchase)
	}
	return purchases, nil
}

func (gr *purchaseRepo) Delete(id uuid.UUID) error {
	_, err := gr.db.Exec(`delete from purchases where id=$1`, id)

	return err
}
