package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// DB는 데이터베이스 커넥션 풀을 가집니다.
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenConns = 10
const maxIdleConns = 5
const maxDbLifetime = 5 * time.Minute

// ConnectSQL은 데이터베이스에 연결하는 주요 메서드입니다.
func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenConns)
	d.SetMaxIdleConns(maxIdleConns)
	d.SetConnMaxLifetime(maxDbLifetime)

	dbConn.SQL = d

	err = testDB(d)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

// testDB는 데이터베이스 연결을 테스트합니다.
func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}

// NewDatabase는 새로운 데이터베이스 연결을 생성합니다.
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
