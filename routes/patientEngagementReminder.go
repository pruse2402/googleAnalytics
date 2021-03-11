package routes

import (
	"cyberliver/go-alcochange-dtx/dtos"
	"cyberliver/go-alcochange-dtx/internals/services/patientEngagementReminderService"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func PatientEngagementReminder(router *httprouter.Router) {
	router.POST("/patient-engagement-reminder", SavePatientEngagementReminder)
}

func SavePatientEngagementReminder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)

	patientEngRem := dtos.ACPatientEngagementReminder{}
	if !parseJSON(w, r.Body, &patientEngRem) {
		return
	}

	p := patientEngagementReminderService.New(rd.l, rd.dbConnMSSQL)
	res, err := p.InsertPatientEngagementReminder(patientEngRem)
	if err != nil {
		writeJSONMessage(err.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}
