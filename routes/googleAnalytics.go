package routes

import (
	"context"
	"go-alcochange-dtx-ga-ga/go-alcochange-dtx-ga/internals/services/googleAnalyticsService"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GoogleAnalytics(router *httprouter.Router) {
	router.POST("/saveGoogleAnalytics", SaveGoogleAnalytics)
}

func SaveGoogleAnalytics(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rd := logAndGetContext(w, r)
	ctx := context.Background()

	keys := r.URL.Query()
	requestDate := keys.Get("date")
	g := googleAnalyticsService.New(rd.l, rd.dbConnMSSQL)
	res, err := g.InsertGoogleAnalytics(requestDate, ctx)
	if err != nil {
		writeJSONMessage(err.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}

	writeJSONStruct(res, http.StatusOK, rd)
}
