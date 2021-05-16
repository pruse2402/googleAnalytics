package jobs

import (
	"context"
	"fmt"
	"go-alcochange-dtx-ga-ga/go-alcochange-dtx-ga/dbcon/mssqlcon"
	"go-alcochange-dtx-ga-ga/go-alcochange-dtx-ga/internals/services/googleAnalyticsService"
	"go-alcochange-dtx-ga-ga/go-alcochange-dtx-ga/utils"
	"time"

	"github.com/FenixAra/go-util/log"
	"github.com/robfig/cron"
)

//ScheduleCronForAttendanceCheckOutEntryLog func
func ScheduleCronGoogleAnalytics() {
	fmt.Println("Init", "ScheduleCronGoogleAnalytics")
	c := cron.New()
	cronExp1 := fmt.Sprintf("0 0 23 * * ?")
	c.AddFunc(cronExp1, func() {
		ctx := context.Background()
		previousDay := time.Now().Add(-24 * time.Hour)
		dateStr := utils.DateToStringDFyyyyMMdd(previousDay)
		cfg := log.NewConfig("alcochange-dtx")
		cfg.SetLevelStr("Debug")
		cfg.SetFilePathSizeStr("Full")
		l := log.New(cfg)
		dbConn := new(mssqlcon.DBConn)
		dbConn.Init(l)

		l.Debug("Job Started--", "working")

		g := googleAnalyticsService.New(l, dbConn)
		res, err := g.InsertGoogleAnalytics(dateStr, ctx)
		if err != nil {
			l.Error("Error-------", "Something wrong", err)
		}
		l.Debug("&&&&&&&&&&&", "working", res)
	})
	c.Start()
}
