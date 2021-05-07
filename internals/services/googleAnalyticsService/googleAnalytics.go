package googleAnalyticsService

import (
	"context"
	"encoding/json"
	"fmt"
	"go-alcochange-dtx-ga-ga/go-alcochange-dtx-ga/dbcon/mssqlcon"
	"go-alcochange-dtx-ga-ga/go-alcochange-dtx-ga/dtos"
	"go-alcochange-dtx-ga-ga/go-alcochange-dtx-ga/internals/adapter/googleAnalyticsAdapter"
	"go-alcochange-dtx-ga-ga/go-alcochange-dtx-ga/internals/daos"
	"strings"

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
	if date != "" {
		date = strings.Replace(date, "-", "", 2)
	}

	ga.l.Debug("google Analytics Initiated ", date)
	resp, err := googleRes.GetGoogleAnalytics(date, ctx)
	if err != nil {
		ga.l.Error("InsertGoogleAnalytics Error - ", err)
		return nil, err
	}

	data, _ := json.Marshal(resp)
	var googleAnalytics []dtos.GA_GoogleAnalytics
	json.Unmarshal(data, &googleAnalytics)

	ga.l.Debug("googleAnalytics ########## -Date, total records taken from GA ", date, len(*resp))

	recordSize := len(*resp)

	for _, gaRow := range googleAnalytics {

		dateStr := ""
		if gaRow.EventDate != "" {
			year := gaRow.EventDate[0:4]
			month := gaRow.EventDate[4:6]
			day := gaRow.EventDate[6:8]
			dateStr = fmt.Sprintf("%v-%v-%v", year, month, day)
		}
		//ga.l.Debug("googleAnalytics", "date", dateStr, gaRow.EventDate, gaRow.Device)
		errG := ga.googleAnalytics.InsertGoogleAnalyticsData(gaRow, dateStr)
		if errG != nil {
			ga.l.Error("InsertGoogleAnalytics Error - ", gaRow.EventTimestamp, gaRow.EventName, errG)
			rowError := GAInsertError{}
			rowError.Error = errG.Error()
			response.Error = append(response.Error, rowError)
			recordSize--
			continue
		}
	}

	resmsg := fmt.Sprintf("%d row Google Analytics data Inserted successfully ", recordSize)
	response.Message = resmsg

	ga.l.Debug("-", resmsg)

	return &response, nil
}

type GoogleAnalyticsResponse struct {
	Message string          `json:"message"`
	Error   []GAInsertError `json:"error"`
}
type GAInsertError struct {
	Error string `json:"error"`
}

func (ga *GoogleAnalytics) DeleteGARecords(date string, ctx context.Context) (*GoogleAnalyticsResponse, error) {
	response := GoogleAnalyticsResponse{}
	// if date != "" {
	// 	date = strings.Replace(date, "-", "", 2)
	// }
	count, errG := ga.googleAnalytics.DeleteGoogleAnalyticsData(date)
	if errG != nil {
		ga.l.Error("InsertGoogleAnalytics Error - ", errG)
		rowError := GAInsertError{}
		rowError.Error = errG.Error()
		response.Error = append(response.Error, rowError)
	}
	resmsg := fmt.Sprintf("%d Rows Deleted successfully ", count)
	response.Message = resmsg
	return &response, nil
}
