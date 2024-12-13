package repository

import (
	"context"
	"database/sql"
	"log"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type DataStore interface {
	StartTransaction(ctx context.Context, fn func(ds DataStore) (any, error)) (any, error)
	GetUserRepository() UserRepository
}

type dataStore struct {
	conn *sql.DB
	db   DBTX
}

func NewDataStore(db *sql.DB) dataStore {
	return dataStore{
		conn: db,
		db:   db,
	}
}

func (ds dataStore) GetUserRepository() UserRepository {
	return NewUserRepository(ds.db)
}

func (ds dataStore) StartTransaction(ctx context.Context, fn func(ds DataStore) (any, error)) (any, error) {
	tx, err := ds.conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}

		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("Error on rollback: %v", rbErr)
				return
			}
			log.Printf("Error on transaction: %v", err)
			return
		}
		tx.Commit()
	}()

	dataStore := &dataStore{
		conn: ds.conn,
		db:   tx,
	}

	result, err := fn(dataStore)
	if err != nil {
		return nil, err
	}

	return result, nil
}
