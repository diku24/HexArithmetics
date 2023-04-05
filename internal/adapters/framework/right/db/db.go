package db

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	//Connnect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		logrus.Fatalf("db connection failure: %v", err)
	}
	//test db connection
	err = db.Ping()
	if err != nil {
		logrus.Fatalf("db ping failure: %v", err)
	}
	return &Adapter{db: db}, nil
}

func (da Adapter) CloseDbConnection() {
	err := da.db.Close()
	if err != nil {
		logrus.Fatalf("db close failure: %v", err)
	}
}

func (da Adapter) AddToHistory(answer int32, operation string) error {
	queryString, args, err := sq.Insert("arith_history").Columns("date", "answer", "operation").Values(time.Now(), answer, operation).ToSql()
	if err != nil {
		return err
	}

	_, err = da.db.Exec(queryString, args...)
	if err != nil {
		return err
	}
	return nil
}
