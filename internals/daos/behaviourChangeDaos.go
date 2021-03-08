package daos

import (
	"cyberliver/go-alcochange-dtx/dbcon/mssqlcon"
	"cyberliver/go-alcochange-dtx/dtos"
	"fmt"

	"github.com/FenixAra/go-util/log"
)

type BehaviourChange struct {
	l           *log.Logger
	dbConnMSSQL *mssqlcon.DBConn
}

func NewBehaviourChange(l *log.Logger, dbConnMSSQL *mssqlcon.DBConn) *BehaviourChange {
	return &BehaviourChange{
		l:           l,
		dbConnMSSQL: dbConnMSSQL,
	}
}

type BehaviourChangeDao interface {
	GetBehaviourChangeList() (*[]dtos.ACBehaviourChange, error)
}

func (bc *BehaviourChange) GetBehaviourChangeList() (*[]dtos.ACBehaviourChange, error) {
	list := []dtos.ACBehaviourChange{}
	rows, err := bc.dbConnMSSQL.GetQueryer().Query(fmt.Sprintf("SELECT behaviour_change_id,bct_taxonomy_id, bct_taxonomy, bct_id,bct_description from ac_behaviour_change_techniques"))
	if err != nil {
		bc.l.Error("Error GetBehaviourChangeList ", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		info := dtos.ACBehaviourChange{}
		err := rows.Scan(&info.BehaviourChangeID,
			&info.BctTaxonomyID,
			&info.BctTaxonomy,
			&info.BctID,
			&info.BctDescription)
		if err != nil {
			bc.l.Error("Error GetBehaviourChangeList - ", err)
			return nil, err
		}
		list = append(list, info)
	}
	return &list, nil
}
