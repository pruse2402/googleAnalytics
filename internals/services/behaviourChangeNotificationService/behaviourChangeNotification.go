package behaviourChangeNotificationService

import (
	"cyberliver/go-alcochange-dtx/dbcon/mssqlcon"
	"cyberliver/go-alcochange-dtx/dtos"
	"cyberliver/go-alcochange-dtx/internals/daos"

	"github.com/FenixAra/go-util/log"
)

type BehaviourChangeNotification struct {
	l                           *log.Logger
	dbConnMSSQL                 *mssqlcon.DBConn
	behaviourChangeNotification daos.BehaviourChangeNotificationDao
}

func New(l *log.Logger, dbConnMSSQL *mssqlcon.DBConn) *BehaviourChangeNotification {
	return &BehaviourChangeNotification{
		l:                           l,
		dbConnMSSQL:                 dbConnMSSQL,
		behaviourChangeNotification: daos.NewBehaviourChangeNotification(l, dbConnMSSQL),
	}
}

func (bc *BehaviourChangeNotification) GetBehaviourChangeNotification() (*dtos.ACBehaviourChangeNotificationResponse, error) {

	bcn := dtos.ACBehaviourChangeNotificationResponse{}

	behaviourChangeNotificationIns, errL := bc.behaviourChangeNotification.GetBehaviourChangeNotificationList()
	if errL != nil {
		bc.l.Error("GetBehaviourChangeNotification Error - ", errL)
		return nil, errL
	}

	bcn.BehaviourChangeNotification = behaviourChangeNotificationIns

	return &bcn, nil
}
