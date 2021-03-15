package daos

import (
	"cyberliver/go-alcochange-dtx/dbcon/mssqlcon"
	"cyberliver/go-alcochange-dtx/dtos"
	"fmt"

	"github.com/FenixAra/go-util/log"
)

type BenefitTherapy struct {
	l           *log.Logger
	dbConnMSSQL *mssqlcon.DBConn
}

func NewBenefitTherapy(l *log.Logger, dbConnMSSQL *mssqlcon.DBConn) *BenefitTherapy {
	return &BenefitTherapy{
		l:           l,
		dbConnMSSQL: dbConnMSSQL,
	}
}

type BenefitTherapyDao interface {
	GetBenefitTherapyList() (*[]dtos.BenefitTherapy, error)
}

func (bt *BenefitTherapy) GetBenefitTherapyList() (*[]dtos.BenefitTherapy, error) {
	list := []dtos.BenefitTherapy{}
	rows, err := bt.dbConnMSSQL.GetQueryer().Query(fmt.Sprintf("SELECT id, message from ac_benefit_therapy"))
	if err != nil {
		bt.l.Error("Error GetBenefitTherapyList ", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		info := dtos.BenefitTherapy{}
		err := rows.Scan(&info.ID,
			&info.Message)

		if err != nil {
			bt.l.Error("Error GetBenefitTherapyList - ", err)
			return nil, err
		}
		list = append(list, info)
	}
	return &list, nil
}
