package daos

import (
	"encoding/json"
	"fmt"
	"go-alcochange-dtx-ga-ga/go-alcochange-dtx-ga/dbcon/mssqlcon"
	"go-alcochange-dtx-ga-ga/go-alcochange-dtx-ga/dtos"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/FenixAra/go-util/log"
)

type GoogleAnalytics struct {
	l           *log.Logger
	dbConnMSSQL *mssqlcon.DBConn
}

func NewGoogleAnalytics(l *log.Logger, dbConnMSSQL *mssqlcon.DBConn) *GoogleAnalytics {
	return &GoogleAnalytics{
		l:           l,
		dbConnMSSQL: dbConnMSSQL,
	}
}

type GoogleAnalyticsDao interface {
	InsertGoogleAnalyticsData(ga dtos.GA_GoogleAnalytics, dateStr string) error
	DeleteGoogleAnalyticsData(dateStr string) (int64, error)
}

func (ga *GoogleAnalytics) InsertGoogleAnalyticsData(googleAnalytics dtos.GA_GoogleAnalytics, dateStr string) error {

	eventParamsBytes, _ := json.Marshal(googleAnalytics.EventParams)
	userProperties, _ := json.Marshal(googleAnalytics.UserProperties)
	deviceBytes, _ := json.Marshal(googleAnalytics.Device)
	geoBytes, _ := json.Marshal(googleAnalytics.Geo)
	userLtvBytes, _ := json.Marshal(googleAnalytics.UserLtv)
	appInfoBytes, _ := json.Marshal(googleAnalytics.AppInfo)
	trafficSourceBytes, _ := json.Marshal(googleAnalytics.TrafficSource)
	eventDimensionsBytes, _ := json.Marshal(googleAnalytics.EventDimensions)
	ecommerceBytes, _ := json.Marshal(googleAnalytics.Ecommerce)
	itemsBytes, _ := json.Marshal(googleAnalytics.Items)

	// ga.l.Debug("eventParamsBytes---", string(eventParamsBytes))
	// ga.l.Debug("userProperties---", string(userProperties))
	// ga.l.Debug("deviceBytes---", string(deviceBytes))
	// ga.l.Debug("geoBytes---", string(geoBytes))
	// ga.l.Debug("userLtvBytes---", string(userLtvBytes))
	// ga.l.Debug("appInfoBytes---", string(appInfoBytes))
	// ga.l.Debug("trafficSourceBytes---", string(trafficSourceBytes))
	// ga.l.Debug("eventDimensionsBytes---", string(eventDimensionsBytes))
	// ga.l.Debug("ecommerceBytes---", string(ecommerceBytes))
	// ga.l.Debug("itemsBytes---", string(itemsBytes))

	gaIns, err := ga.dbConnMSSQL.GetQueryer().Exec(`INSERT INTO google_analytics (
	event_date_ref_str,
	event_date,
	event_timestamp,
	event_name,
	event_params,
	event_previous_timestamp,
	event_value_in_usd,
	event_bundle_sequence_id,
	event_server_timestamp_offset,
	user_id,
	user_pseudo_id,
	user_properties,
	user_first_touch_timestamp,
	user_ltv,
	device,
	geo,
	app_info,
	traffic_source,
	stream_id,
	platform,
	event_dimensions,
	ecommerce,
	items,
	created_at, 
	updated_at)
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		dateStr,
		googleAnalytics.EventDate,
		googleAnalytics.EventPreviousTimestamp,
		googleAnalytics.EventName,
		string(eventParamsBytes),
		googleAnalytics.EventPreviousTimestamp,
		googleAnalytics.EventValueInUsd,
		googleAnalytics.EventBundleSequenceID,
		googleAnalytics.EventServerTimestampOffset,
		googleAnalytics.UserID,
		googleAnalytics.UserPseudoID,
		string(userProperties),
		googleAnalytics.UserFirstTouchTimestamp,
		string(userLtvBytes),
		string(deviceBytes),
		string(geoBytes),
		string(appInfoBytes),
		string(trafficSourceBytes),
		googleAnalytics.StreamID,
		googleAnalytics.Platform,
		string(eventDimensionsBytes),
		string(ecommerceBytes),
		string(itemsBytes),
		time.Now(),
		time.Now())
	if err != nil {
		fmt.Println("gaIns -Error:", err)
		return err
	}

	count, err := gaIns.RowsAffected()
	//ga.l.Debug("RowsAffected count", count)
	if count == 0 {
		fmt.Println("gaInas-No rows returned")
		return err
	}
	return nil

}

func BigqueryToStr(value bigquery.NullString) string {
	if !value.Valid {
		return ""
	}
	return value.StringVal
}

func BigqueryToInt64(value bigquery.NullInt64) int64 {
	if !value.Valid {
		return 0
	}
	return value.Int64
}

func (ga *GoogleAnalytics) DeleteGoogleAnalyticsData(dateStr string) (int64, error) {

	gaIns, err := ga.dbConnMSSQL.GetQueryer().Exec(`delete from google_analytics where event_date_ref_str = ? `, dateStr)
	if err != nil {
		ga.l.Debug("gaIns -Error:", err)
		return 0, err
	}

	count, err := gaIns.RowsAffected()
	//ga.l.Debug("RowsAffected count", count)
	if count == 0 {
		fmt.Println("gaInas-No rows returned")
		return 0, err
	}
	return count, nil

}
