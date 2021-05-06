package googleAnalyticsService

import (
	"context"
	"encoding/json"
	"go-alcochange-dtx-ga-ga/go-alcochange-dtx-ga/dbcon/mssqlcon"
	"go-alcochange-dtx-ga-ga/go-alcochange-dtx-ga/internals/adapter/googleAnalyticsAdapter"
	"go-alcochange-dtx-ga-ga/go-alcochange-dtx-ga/internals/daos"

	"github.com/FenixAra/go-util/log"
)

type GoogleAnalytics struct {
	l               *log.Logger
	dbConnMSSQL     *mssqlcon.DBConn
	googleAnalytics daos.GoogleAnalyticsDao
}

func New(l *log.Logger, dbConnMSSQL *mssqlcon.DBConn) *GoogleAnalytics {
	return &GoogleAnalytics{
		l:               l,
		dbConnMSSQL:     dbConnMSSQL,
		googleAnalytics: daos.NewGoogleAnalytics(l, dbConnMSSQL),
	}
}

func (ga *GoogleAnalytics) InsertGoogleAnalytics(date string, ctx context.Context) (*GoogleAnalyticsResponse, error) {

	response := GoogleAnalyticsResponse{}
	googleRes := googleAnalyticsAdapter.NewGoogleAnalyticsDtx(ga.l)
	resp, err := googleRes.GetGoogleAnalytics(date, ctx)
	if err != nil {
		ga.l.Error("InsertGoogleAnalytics Error - ", err)
		return nil, err
	}

	b, _ := json.Marshal(resp)
	ga.l.Debug("google analytics", string(b))

	errG := ga.googleAnalytics.InsertGoogleAnalyticsData(resp)
	if errG != nil {
		ga.l.Error("InsertGoogleAnalytics Error - ", errG)
		return nil, errG
	}

	response.Message = "Google Analytics data Inserted successfully"

	return &response, nil
}

type GoogleAnalyticsResponse struct {
	Message string `json:"message"`
}
