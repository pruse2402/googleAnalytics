package daos

import (
	"go-alcochange-dtx-ga-ga/go-alcochange-dtx-ga/dbcon/mssqlcon"

	"github.com/FenixAra/go-util/log"
)

type PingCheckStruct struct {
	Count int32 `json:"count"`
}

type Ping struct {
	l           *log.Logger
	dbConnMSSQL *mssqlcon.DBConn
}

func NewPing(l *log.Logger, dbConnMSSQL *mssqlcon.DBConn) *Ping {
	return &Ping{
		l:           l,
		dbConnMSSQL: dbConnMSSQL,
	}
}

func (p *Ping) Ping() (bool, error) {
	pingModel := PingCheckStruct{}
	errMs := p.dbConnMSSQL.GetQueryer().QueryRow("SELECT 1").Scan(&pingModel.Count)
	if errMs != nil {
		p.l.Error("Ping error - MSSQL-", errMs)
		return false, errMs
	}
	return (pingModel.Count == 1), errMs
}
