package db

import (
	"LibraryService/config"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// NewSqlDB initializes a new SQL database connection based on the provided configuration.
//
// It takes a *config.DB parameter which contains the necessary database configuration.
// It returns a *sqlx.DB object and an error if there was an issue establishing the connection.
func NewSqlDB(cfg *config.DB) (*sqlx.DB, error) {
	var dsn string
	var err error
	var dbRaw *sql.DB
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Db)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	timeoutExceeded := time.After(time.Second * time.Duration(cfg.Timeout))

	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %d timeout %s", cfg.Timeout, err)
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
