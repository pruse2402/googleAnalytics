package daos

import (
	"cyberliver/go-alcochange-dtx/dbcon/mssqlcon"
	"cyberliver/go-alcochange-dtx/dtos"
	"fmt"

	"github.com/FenixAra/go-util/log"
)

type BehaviourChangeNotification struct {
	l           *log.Logger
	dbConnMSSQL *mssqlcon.DBConn
}

func NewBehaviourChangeNotification(l *log.Logger, dbConnMSSQL *mssqlcon.DBConn) *BehaviourChangeNotification {
	return &BehaviourChangeNotification{
		l:           l,
		dbConnMSSQL: dbConnMSSQL,
	}
}

type BehaviourChangeNotificationDao interface {
	GetBehaviourChangeNotificationList() (*[]dtos.ACBehaviourChangeNotification, error)
}

func (bc *BehaviourChangeNotification) GetBehaviourChangeNotificationList() (*[]dtos.ACBehaviourChangeNotification, error) {
	list := []dtos.ACBehaviourChangeNotification{}
	rows, err := bc.dbConnMSSQL.GetQueryer().Query(fmt.Sprintf("SELECT behaviour_notification_id, bcn_category, bcn_category_desc, bcn_group, bcn_group_description, bcn_trigger_event, app_route, bcn_id, bcn_message, bct_1, bct_2, bct_3, bct_4, alcochange_theme, frames_feedback, frames_responsibility, frames_advice, frames_menu, frames_empathy, frames_support_and_selfefficacy, develop_discrepancy, assessment from ac_behaviour_change_notifications"))
	if err != nil {
		bc.l.Error("Error GetBehaviourChangeNotificationList ", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		info := dtos.ACBehaviourChangeNotification{}
		err := rows.Scan(&info.BehaviourNotificationID,
			&info.BcnCategory,
			&info.BcnCategoryDesc,
			&info.BcnGroup,
			&info.BcnGroupDescription,
			&info.BcnTriggerEvent,
			&info.AppRoute,
			&info.BcnID,
			&info.BcnMessage,
			&info.Bct1,
			&info.Bct2,
			&info.Bct3,
			&info.Bct4,
			&info.AlcochangeTheme,
			&info.FramesFeedback,
			&info.FramesResponsibility,
			&info.FramesAdvice,
			&info.FramesMenu,
			&info.FramesEmpathy,
			&info.FramesSupportAndSelfefficacy,
			&info.DevelopDiscrepancy,
			&info.Assessment)

		if err != nil {
			bc.l.Error("Error GetBehaviourChangeNotificationList - ", err)
			return nil, err
		}
		list = append(list, info)
	}
	return &list, nil
}
