package googleAnalyticsAdapter

import (
	"context"
	"go-alcochange-dtx-ga-ga/go-alcochange-dtx-ga/conf"
	"go-alcochange-dtx-ga-ga/go-alcochange-dtx-ga/dtos"

	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/FenixAra/go-util/log"
	"google.golang.org/api/iterator"
)

type GoogleAnalyticsDtx struct {
	l *log.Logger
}

func NewGoogleAnalyticsDtx(ll *log.Logger) GoogleAnalyticsDtxDao {
	return &GoogleAnalyticsDtx{
		l: ll,
	}
}

type GoogleAnalyticsDtxDao interface {
	GetGoogleAnalytics(date string, ctx context.Context) (*[]dtos.GoogleAnalytics, error)
}

func (ga GoogleAnalyticsDtx) GetGoogleAnalytics(date string, ctx context.Context) (*[]dtos.GoogleAnalytics, error) {

	googleAnalyticsIns := []dtos.GoogleAnalytics{}

	client, err := bigquery.NewClient(ctx, conf.Cfg.PROJECT_ID)
	if err != nil {
		ga.l.Error("bigquery.NewClient Error ", err.Error())
		return nil, fmt.Errorf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	//getGoogleAnalyticsQuery := fmt.Sprintf("SELECT * FROM alcochange-dtx-dev.analytics_269133399.events_%v", date)
	getGoogleAnalyticsQuery := "SELECT * FROM `alcochange-dtx-dev.analytics_269133399.events_20210416` order by event_timestamp ASC limit 1 offset 1"

	q := client.Query(getGoogleAnalyticsQuery)

	// Location must match that of the dataset(s) referenced in the query.
	q.Location = "US"

	// Run the query and print results when the query job is completed.
	job, err := q.Run(ctx)
	if err != nil {
		ga.l.Error("Run Error ", err.Error())
		return nil, fmt.Errorf("bigquery.job: %v", err)
	}

	status, err := job.Wait(ctx)
	if err != nil {
		ga.l.Error("bigquery.status Error ", err.Error())
		return nil, fmt.Errorf("bigquery.status: %v", err)
	}
	if err := status.Err(); err != nil {
		ga.l.Error("status.Err Error", err.Error())
		return nil, fmt.Errorf("bigquery.status: %v", err)
	}

	it, err := job.Read(ctx)

	for {
		var row dtos.GoogleAnalytics
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			ga.l.Error("Error in fetching data ", err.Error())
			return nil, fmt.Errorf("error in fetching data: %v", err)
		}

		ga.l.Debug("row.EventDate", row.EventName)

		googleAnalyticsIns = append(googleAnalyticsIns, row)

	}

	return &googleAnalyticsIns, nil
}
