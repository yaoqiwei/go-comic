package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
	q                 query
	tx                *sql.Tx
	prefix            string
	suffix            func(interface{}) string
	errHandler        func(error, *Orm)
	afterQueryHandler func(*Orm)
}

func (db *DB) Table(s ...interface{}) *Orm {
	return db.Default().Table(s...)
}

func (db *DB) Default() *Orm {
	return &Orm{db: db}
}

func (db *DB) WithPrefix(p string) *DB {
	db.prefix = p
	return db
}

func (db *DB) WithSuffix(p func(interface{}) string) *DB {
	db.suffix = p
	return db
}

func (db *DB) WithErrHandler(p func(error, *Orm)) *DB {
	db.errHandler = p
	return db
}

func (db *DB) WithAfterQueryHandler(p func(*Orm)) *DB {
	db.afterQueryHandler = p
	return db
}

func (db *DB) Start() *DB {

	q := db.q
	tx, err := db.Begin()
	if err == nil {
		q = tx
	}

	return &DB{
		DB:                db.DB,
		q:                 q,
		tx:                tx,
		prefix:            db.prefix,
		suffix:            db.suffix,
		errHandler:        db.errHandler,
		afterQueryHandler: db.afterQueryHandler,
	}
}

func (db *DB) Rollback() (err error) {
	err = db.tx.Rollback()
	return
}

func (db *DB) Commit() (err error) {
	err = db.tx.Commit()
	return
}

func Open(driverName, dataSourceName string) (*DB, error) {
	d, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{DB: d, q: d}, nil
}
