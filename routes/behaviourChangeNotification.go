package routes

import (
	"cyberliver/go-alcochange-dtx/internals/services/behaviourChangeNotificationService"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func BehaviourChangeNotification(router *httprouter.Router) {
	router.GET("/ac-behaviour-change-notification", GetBehaviourChangeNotification)
}

func GetBehaviourChangeNotification(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	p := behaviourChangeNotificationService.New(rd.l, rd.dbConnMSSQL)
	res, err := p.GetBehaviourChangeNotification()
	if err != nil {
		writeJSONMessage(err.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}
