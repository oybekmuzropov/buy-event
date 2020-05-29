package potgres

import (
	"github.com/buy_event/storage/repo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type logRepo struct {
	db *sqlx.DB
}

func NewLogRepo(db *sqlx.DB) repo.LogStorageI {
	return &logRepo{db:db}
}

func (lr *logRepo) Create(log *repo.Log) error {
	//generate id
	id, err := uuid.NewRandom()
	
	if err != nil {
		return err
	}
	
	_, err = lr.db.Exec(`insert into logs(id, error) values ($1, $2)`, id, log.Error)
	
	return err
}

func (lr *logRepo) GetAll() ([]*repo.Log, error) {
	var logs []*repo.Log
	
	rows, err := lr.db.Queryx(`select id, error from logs`)
	
	if err != nil {
		return nil, err
	}
	
	for rows.Next() {
		var log repo.Log
		
		err := rows.Scan(
			&log.ID,
			&log.Error,
		)
		
		if err != nil {
			return nil, err
		}
		
		logs = append(logs, &log)
	}

	return logs, nil
}