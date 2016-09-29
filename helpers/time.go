package helpers

import (
	"time"
)

func StrToLocalTime(StrDateTime string) int64 {
	set_location, _ := time.LoadLocation("Local")
	my_date_formate := "2006-01-02 15:04:05"
	t, _ := time.ParseInLocation(my_date_formate, StrDateTime, set_location)
	return t.Unix()
}

func StrToUtcTime(StrDateTime string) int64 {
	my_date_formate := "2006-01-02 15:04:05"
	t, _ := time.Parse(my_date_formate, StrDateTime)
	return t.Unix()
}

func DiffUnxiTime(StartDateTime, EndDateTime string) int64 {
	return StrToLocalTime(EndDateTime) - StrToLocalTime(StartDateTime)
}

func StrToFormateDate(StrDateTime string) time.Time {
	set_location, _ := time.LoadLocation("Local")
	my_date_formate := "2006-01-02 15:04:05"
	t, _ := time.ParseInLocation(my_date_formate, StrDateTime, set_location)

	return t
}

func MyNowDate(format ...string) string {
	if len(format) == 0 {
		format = append(format, "2006-01-02 15:04:05")
	}
	return time.Now().Format(format[0])
}
