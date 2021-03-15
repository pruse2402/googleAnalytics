package benefitTherapyService

import (
	"cyberliver/go-alcochange-dtx/dbcon/mssqlcon"
	"cyberliver/go-alcochange-dtx/dtos"
	"cyberliver/go-alcochange-dtx/internals/daos"

	"github.com/FenixAra/go-util/log"
)

type BenefitTherapy struct {
	l              *log.Logger
	dbConnMSSQL    *mssqlcon.DBConn
	benefitTherapy daos.BenefitTherapyDao
}

func New(l *log.Logger, dbConnMSSQL *mssqlcon.DBConn) *BenefitTherapy {
	return &BenefitTherapy{
		l:              l,
		dbConnMSSQL:    dbConnMSSQL,
		benefitTherapy: daos.NewBenefitTherapy(l, dbConnMSSQL),
	}
}

func (bt *BenefitTherapy) GetBenefitTherapy() (*dtos.BenefitTherapyResponse, error) {

	btIns := dtos.BenefitTherapyResponse{}

	benefitTherapyIns, errL := bt.benefitTherapy.GetBenefitTherapyList()
	if errL != nil {
		bt.l.Error("GetBenefitTherapy Error - ", errL)
		return nil, errL
	}

	btIns.Title = "Benefits of AlcoChange Therapy"
	btIns.Header = "By abstaining from alcohol you will…"
	btIns.ButtonText = "I commit to abstaining from alcohol"
	btIns.BenefitTherapy = benefitTherapyIns

	return &btIns, nil
}
