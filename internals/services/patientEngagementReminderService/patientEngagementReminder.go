package patientEngagementReminderService

import (
	"cyberliver/go-alcochange-dtx/dbcon/mssqlcon"
	"cyberliver/go-alcochange-dtx/dtos"
	"cyberliver/go-alcochange-dtx/internals/daos"

	"github.com/FenixAra/go-util/log"
)

type PatientEngagementReminder struct {
	l                         *log.Logger
	dbConnMSSQL               *mssqlcon.DBConn
	patientEngagementReminder daos.PatientEngagementReminderDao
}

func New(l *log.Logger, dbConnMSSQL *mssqlcon.DBConn) *PatientEngagementReminder {
	return &PatientEngagementReminder{
		l:                         l,
		dbConnMSSQL:               dbConnMSSQL,
		patientEngagementReminder: daos.NewPatientEngagementReminder(l, dbConnMSSQL),
	}
}

func (per *PatientEngagementReminder) InsertPatientEngagementReminder(patientEngagementReminder dtos.ACPatientEngagementReminder) (*dtos.ResponseMessage, error) {

	if patientEngagementReminder.ScheduledIntervention != nil {
		for _, scheduledInterventions := range patientEngagementReminder.ScheduledIntervention {
			scheduledInterventions.UserAction = "Scheduled"
			errL := per.patientEngagementReminder.InsertScheduledPatientEngagementReminder(scheduledInterventions, patientEngagementReminder.UserUuid, patientEngagementReminder.UserID)
			if errL != nil {
				per.l.Error("InsertPatientEngagementReminder Error - ", errL)
				return nil, errL
			}
		}
	}

	if patientEngagementReminder.EngagedIntervention != nil {
		for _, engagedIntervention := range patientEngagementReminder.EngagedIntervention {
			engagedIntervention.UserAction = "Responded"
			errL := per.patientEngagementReminder.InsertEngagedPatientEngagementReminder(engagedIntervention, patientEngagementReminder.UserUuid, patientEngagementReminder.UserID)
			if errL != nil {
				per.l.Error("InsertPatientEngagementReminder Error - ", errL)
				return nil, errL
			}
		}
	}

	responseMessage := dtos.ResponseMessage{}
	responseMessage.Message = "saved successfully"
	return &responseMessage, nil
}
