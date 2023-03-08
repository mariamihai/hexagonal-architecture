package db

import (
	"database/sql"
	"fmt"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		return nil, fmt.Errorf("could not connect to the db: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("couldn't ping the db: %w", err)
	}

	return &Adapter{
		db: db,
	}, nil
}

func (a Adapter) CloseDBConnection() error {
	return a.db.Close()
}

func (a Adapter) AddToHistory(answer int, operation string) error {
	_, err := a.db.Exec(`INSERT INTO history (answer, operation) VALUES($1, $2)`, answer, operation)

	return err
}
