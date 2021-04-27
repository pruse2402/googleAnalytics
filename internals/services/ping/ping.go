package ping

import (
	"errors"
	"go-alcochange-dtx-ga-ga/go-alcochange-dtx-ga/dbcon/mssqlcon"
	"go-alcochange-dtx-ga-ga/go-alcochange-dtx-ga/internals/daos"

	"github.com/FenixAra/go-util/log"
)

type Ping struct {
	l           *log.Logger
	ping        *daos.Ping
	dbConnMSSQL *mssqlcon.DBConn
}

var (
	ErrUnableToPingDB = errors.New("Unable to ping database")
)

func New(l *log.Logger, dbConnMSSQL *mssqlcon.DBConn) *Ping {
	return &Ping{
		l:           l,
		dbConnMSSQL: dbConnMSSQL,
		ping:        daos.NewPing(l, dbConnMSSQL),
	}
}

func (p *Ping) Ping() error {
	ok, err := p.ping.Ping()
	if err != nil {
		return err
	}
	if !ok {
		return ErrUnableToPingDB
	}
	return nil
}
