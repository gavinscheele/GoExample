package psql

import (
	"database/sql"

	"github.com/gdscheele/Example/src/model"
)

type PsqlRepository struct {
	conn *sql.DB
}

func NewPsqlRepository(dbHost, dbName, username, password string) (*PsqlRepository, error) {
	db := NewDatabase(dbHost, dbName, username, password)
	conn, err := db.Open()
	if err != nil {
		return nil, err
	}

	return &PsqlRepository{conn}, nil
}

func (repo *PsqlRepository) Close() error {
	return repo.conn.Close()
}

func (repo *PsqlRepository) GetUser(ID int64) (*model.User, error) {
	user := model.User{}
	err := repo.conn.QueryRow("SELECT id, username FROM users WHERE id=$1",
		ID).Scan(&user.ID, &user.Username)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, err
}

func (repo *PsqlRepository) DeleteUser(ID int64) error {
	_, err := repo.conn.Exec("DELETE FROM users WHERE id=$1", ID)
	return err
}

func (repo *PsqlRepository) CreateUser(username string) (int64, error) {
	var id int64
	err := repo.conn.QueryRow("INSERT INTO users(username) VALUES($1) RETURNING id", username).Scan(&id)
	return id, err
}
