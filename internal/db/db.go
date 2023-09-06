package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewSqlDB() (*sqlx.DB, error) {
	var dsn string
	var err error
	var dbRaw *sql.DB
	dsn = "root:root@tcp(localhost:3306)/library?parseTime=true"
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	timeoutExceeded := time.After(time.Second * 5)

	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %d timeout %s", 5, err)
		case <-ticker.C:
			dbRaw, err = sql.Open("mysql", dsn)
			if err != nil {
				return nil, err
			}
			err = dbRaw.Ping()
			if err == nil {
				db := sqlx.NewDb(dbRaw, "mysql")
				db.SetMaxOpenConns(50)
				db.SetMaxIdleConns(50)
				return db, nil
			}
		}
	}
}
