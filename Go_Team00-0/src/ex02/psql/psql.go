package psql

import (
	"database/sql"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Model struct {
	DB *sql.DB
}

func NewDbConnection(host, port, user, password, dbname string) (*Model, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Model{DB: db}, nil
}

func (m *Model) Insert(sessionId string, frequency float64, UTC timestamppb.Timestamp) error {
	stmt := `INSERT INTO logger (sessionid, frequency, utc)
	VALUES ($1, $2, $3);`
	_, err := m.DB.Exec(stmt, sessionId, frequency, UTC.AsTime().String())
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) CreateTable() error {
	stmt := `CREATE TABLE IF NOT EXISTS logger (
    sessionId varchar(20) NOT NULL ,
    frequency decimal not null,
    UTC varchar(50) );`
	_, err := m.DB.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}
