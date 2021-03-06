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
	GetBehaviourChangeList() (*[]dtos.ACBehaviourChangeResponse, error)
}

func (bc *BehaviourChange) GetBehaviourChangeList() (*[]dtos.ACBehaviourChangeResponse, error) {
	list := []dtos.ACBehaviourChangeResponse{}
	rows, err := bc.dbConnMSSQL.GetQueryer().Query(fmt.Sprintf("SELECT behaviour_change_id,bct_taxonomy_id, bct_taxonomy, bct_id,bct_description, created_at, updated_at from ac_behaviour_change_techniques"))
	if err != nil {
		bc.l.Error("Error GetBehaviourChangeList ", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		info := dtos.ACBehaviourChangeResponse{}
		err := rows.Scan(&info.BehaviourChangeID,
			&info.BctTaxonomyID,
			&info.BctTaxonomy,
			&info.BctID,
			&info.BctDescription,
			&info.CreatedAt,
			&info.CreatedAt)
		if err != nil {
			bc.l.Error("Error GetBehaviourChangeList - ", err)
			return nil, err
		}
		list = append(list, info)
	}
	return &list, nil
}
