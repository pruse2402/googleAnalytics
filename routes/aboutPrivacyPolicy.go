package routes

import (
	"net/http"

	"cyberliver/go-alcochange-dtx/internals/services/aboutPrivacyPolicy"

	"github.com/julienschmidt/httprouter"
)

func GetAboutPrivacyPolicy(router *httprouter.Router) {
	router.GET("/ac-about/privacy-policy", AboutPrivacyPolicy)
}

func AboutPrivacyPolicy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	p := aboutPrivacyPolicy.New(rd.l, rd.dbConnMSSQL)
	res, err := p.AboutPrivacyPolicy()
	if err != nil {
		writeJSONMessage(err.Error(), ERR_MSG, http.StatusInternalServerError, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}
