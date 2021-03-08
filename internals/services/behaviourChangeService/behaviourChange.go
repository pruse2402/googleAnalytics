package behaviourChangeService

import (
	"cyberliver/go-alcochange-dtx/dbcon/mssqlcon"
	"cyberliver/go-alcochange-dtx/dtos"
	"cyberliver/go-alcochange-dtx/internals/daos"

	"github.com/FenixAra/go-util/log"
)

type BehaviourChange struct {
	l               *log.Logger
	dbConnMSSQL     *mssqlcon.DBConn
	behaviourChange daos.BehaviourChangeDao
}

func New(l *log.Logger, dbConnMSSQL *mssqlcon.DBConn) *BehaviourChange {
	return &BehaviourChange{
		l:               l,
		dbConnMSSQL:     dbConnMSSQL,
		behaviourChange: daos.NewBehaviourChange(l, dbConnMSSQL),
	}
}

func (bc *BehaviourChange) GetBehaviourChange() (*dtos.ACBehaviourChangeResponse, error) {

	bct := dtos.ACBehaviourChangeResponse{}

	behaviourChangeIns, errL := bc.behaviourChange.GetBehaviourChangeList()
	if errL != nil {
		bc.l.Error("GetBehaviourChange Error - ", errL)
		return nil, errL
	}

	bct.BehaviourChange = behaviourChangeIns

	return &bct, nil
}
