package routes

import (
	"cyberliver/go-alcochange-dtx/internals/services/behaviourChangeService"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func BehaviourChange(router *httprouter.Router) {
	router.GET("/ac-behaviourChange", GetBehaviourChange)
}

func GetBehaviourChange(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	p := behaviourChangeService.New(rd.l, rd.dbConnMSSQL)
	res, err := p.GetBehaviourChange()
	if err != nil {
		writeJSONMessage(err.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}
