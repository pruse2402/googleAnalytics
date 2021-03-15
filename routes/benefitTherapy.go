package routes

import (
	"cyberliver/go-alcochange-dtx/internals/services/benefitTherapyService"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func BenefitTherapy(router *httprouter.Router) {
	router.GET("/ac-benefit-therapy", GetBenefitTherapyList)
}

func GetBenefitTherapyList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	p := benefitTherapyService.New(rd.l, rd.dbConnMSSQL)
	res, err := p.GetBenefitTherapy()
	if err != nil {
		writeJSONMessage(err.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}
