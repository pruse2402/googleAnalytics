package utils

import (
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	//TFhhmm time format hour:minute
	TFhhmm = "15:04"
)

// DateToHHMMFormat convert date to hh:mm
func DateToHHMMFormat(time time.Time) string {
	return time.Format(TFhhmm)
}

//CheckTimeFormatHHMM function is use to check given time is valid or not
func CheckTimeFormatHHMM(timeStr string) bool {
	_, err := time.ParseInLocation(TFhhmm, timeStr, time.UTC)
	if err != nil {
		return false
	}
	return true
}

// HHMMToTime function is use to convert hhmm to date format
func HHMMToTime(date time.Time, timeStr string) time.Time {

	strArr := strings.Split(timeStr, ":")

	hour, _ := strconv.Atoi(strArr[0])
	min, _ := strconv.Atoi(strArr[1])
	// t := time.Now()

	t1 := time.Date(date.Year(), date.Month(), date.Day(), hour, min, 00, 0, time.UTC)

	return t1.UTC()
}

func StringToDate(dateString interface{}) (date time.Time) {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		log.Printf("Error in load UTC location: %s", err.Error())
		return
	}

	dateStr, _ := dateString.(string)
	date, _ = time.ParseInLocation("01-02-2006 15:04:05", dateStr, loc)
	return
}

func DateToStringFormat(date time.Time) string {
	return date.Format("2006-01-02 15:04")
}
