package storage

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"time"
)

type DataBaseStorage struct {
	db *sql.DB
}

func Connection(strConnection string) (*DataBaseStorage, error) {
	db, err := sql.Open("pgx", strConnection)
	if err != nil {
		return nil, err
	}

	log.Println("Открыто подключение к базе", db)

	return &DataBaseStorage{db: db}, nil
}

func (stge *DataBaseStorage) ConnectionCheck() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := stge.db.PingContext(ctx); err != nil {
		return false, err
	}

	return true, nil
}

func (stge *DataBaseStorage) SaveLinkToStge(hash string, original []byte) error {
	createdAt := time.Now()
	_, err := stge.db.Exec("INSERT INTO pairs (original, short, created_at) VALUES ($1, $2, $3)", string(original), hash, createdAt.Format(time.DateTime))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (stge *DataBaseStorage) GetLinkFromStge(hash []byte) ([]byte, error) {
	links, err := stge.db.Query("SELECT original FROM pairs WHERE short = $1", hash)
	if err != nil {
		return nil, err
	}
	defer links.Close()

	var original string
	for links.Next() {
		if err = links.Scan(&original); err != nil {
			return nil, err
		}
	}

	err = links.Err()
	if err != nil {
		return nil, err
	}

	return []byte(original), nil
}