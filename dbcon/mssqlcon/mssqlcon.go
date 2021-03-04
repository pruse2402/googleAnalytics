package mssqlcon

import (
	"context"
	"database/sql"
	lg "log"
	"sync"

	"github.com/FenixAra/go-util/log"
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/hgfischer/mysql"
)

//Pool of database connection
var once sync.Once
var ConnPool *sql.DB

type DBConn struct {
	conn          *sql.DB
	tx            *sql.Tx
	isTransaction bool
	l             *log.Logger
}
type IDBConn interface {
	Init(*log.Logger)
	GetQueryer() Queryer
	ExecuteInTransaction(func() error) error
	// rollbackTransaction(tx *sql.Tx)
}

// Interface to abstract the queryer(dbconnection or transaction)
type Queryer interface {
	Exec(sql string, arguments ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, sql string, args ...interface{}) (sql.Result, error)
	Query(sql string, args ...interface{}) (*sql.Rows, error)
	QueryRow(sql string, args ...interface{}) *sql.Row
	Prepare(sql string) (*sql.Stmt, error)
}

func New(l *log.Logger) *DBConn {
	return &DBConn{
		l:    l,
		conn: ConnPool,
	}
}

func NewDBConn(l *log.Logger, conn *sql.DB) *DBConn {
	return &DBConn{
		l:    l,
		conn: conn,
	}
}

// Initialize the DB connection and assign the existing db connection
func (db *DBConn) Init(l *log.Logger) {
	db.conn = ConnPool
	db.l = l
}

func (db *DBConn) GetQueryer() Queryer {
	if db.isTransaction {
		return db.tx
	} else {
		return db.conn
	}
}

// ExecuteInTransaction executes the given function in DB transaction, i.e. It commits
// only if there is not error otherwise it is rolledback.
func (db *DBConn) ExecuteInTransaction(f func() error) (err error) {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	db.tx = tx
	db.isTransaction = true

	defer func() {
		if r := recover(); r != nil {
			db.l.Fatal("Recovered in function ", r)
			db.rollbackTransaction(tx)
		}
		db.isTransaction = false
	}()

	err = f()
	if err != nil {
		db.rollbackTransaction(tx)
		return err
	}
	err = tx.Commit()
	if err != nil {
		db.rollbackTransaction(tx)
		return err
	}
	return nil
}

func (db *DBConn) rollbackTransaction(tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil {
		db.l.Error("Error While rollback, Err: ", err)
	}
}

func MSSqlInit(url string) {
	if ConnPool == nil {
		once.Do(func() {
			var err error
			// Create connection pool
			ConnPool, err = sql.Open("mysql", url)
			if err != nil {
				lg.Println("Error creating connection pool: %+v", err)
			}
			ConnPool.SetMaxOpenConns(15)
			ctx := context.Background()
			err = ConnPool.PingContext(ctx)
			if err != nil {
				lg.Println("Unable to ping to DB. Err: %+v", err)
				return
			}
			lg.Println("Connected to database successfully!")
		})
	}
}

//ConnPool.Close()

func MSSqlConnClose() {
	if ConnPool != nil {
		ConnPool.Close()
	}
}

// func Init(url string) {
// 	if ConnPool == nil {
// 		once.Do(func() {
// 			var err error
// 			// Create connection pool
// 			ConnPool, err = sql.Open("sqlserver", url)
// 			if err != nil {
// 				globals.Logger.Fatalf("Error creating connection pool: %+v", err)
// 			}

// 			ConnPool.SetMaxOpenConns(config.MAX_DB_CONNECTIONS)
// 			ctx := context.Background()
// 			err = ConnPool.PingContext(ctx)
// 			if err != nil {
// 				globals.Logger.Errorf("Unable to ping to DB. Err: %+v", err)
// 				return
// 			}
// 			globals.Logger.Infof("Connected to database successfully!")
// 		})
// 	}
// }
