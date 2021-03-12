package daos

import (
	"cyberliver/go-alcochange-dtx/dbcon/mssqlcon"
	"cyberliver/go-alcochange-dtx/dtos"
	"fmt"

	"github.com/FenixAra/go-util/log"
)

type PatientEngagementReminder struct {
	l           *log.Logger
	dbConnMSSQL *mssqlcon.DBConn
}

func NewPatientEngagementReminder(l *log.Logger, dbConnMSSQL *mssqlcon.DBConn) *PatientEngagementReminder {
	return &PatientEngagementReminder{
		l:           l,
		dbConnMSSQL: dbConnMSSQL,
	}
}

type PatientEngagementReminderDao interface {
	InsertScheduledPatientEngagementReminder(scheduledInterventions dtos.ScheduledIntervention, userUuid string, userID int64) error
	InsertEngagedPatientEngagementReminder(engagedInterventions dtos.EngagedIntervention, userUuid string, userID int64) error
}

func (per *PatientEngagementReminder) InsertScheduledPatientEngagementReminder(req dtos.ScheduledIntervention, userUuid string, userID int64) error {

	query := fmt.Sprintf(`INSERT INTO ac_patient_engagement_reminder(user_id, user_uuid, intervention_type_id, notification_id, patient_engagement_time, message_shown, user_action) values('%v', '%v', '%v', '%v', '%v', '%v', '%v')`, userID, userUuid, req.InterventionTypeID, req.NotificationID, req.PatientEngagementTime, req.MessageShown, req.UserAction)
	_, err := per.dbConnMSSQL.GetQueryer().Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (per *PatientEngagementReminder) InsertEngagedPatientEngagementReminder(req dtos.EngagedIntervention, userUuid string, userID int64) error {

	query := fmt.Sprintf(`INSERT INTO ac_patient_engagement_reminder(user_id, user_uuid, intervention_type_id, notification_id, patient_engagement_time, message_shown, user_action) values('%v', '%v', '%v', '%v', '%v', '%v', '%v')`, userID, userUuid, req.InterventionTypeID, req.NotificationID, req.PatientEngagementTime, req.MessageShown, req.UserAction)

	_, err := per.dbConnMSSQL.GetQueryer().Exec(query)

	if err != nil {
		return err
	}

	return nil
}

// INSERT INTO cyberliver_platform.ac_patient_engagement_reminder
// (user_id, user_uuid, intervention_type_id, notification_id, patient_engagement_time, message_shown, user_action, created_at, updated_at)
// VALUES(0, '', 0, 0, '', '', '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
