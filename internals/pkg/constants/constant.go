package constants

import "time"

var (
	DirPath         = "../input"
	TIME_LOCAL, err = time.LoadLocation(TIME_ZONE)
)

const (
	BatchStatusProcessing = "processing"
	BatchStatusSuccess    = "success"
	BatchStatusFailed     = "failed"
	FILE_NAME             = "test"
	DATE_FORMAT           = "20060102"
	DATE_FORMATTER        = "20060102150405"
	TIME_ZONE             = "Asia/Bangkok"
	SUFFIX                = ".csv"
	DATE_TIME_FORMATTER   = "2006-01-02 15:04:05"
	REGEX_FILE_NAME       = `^test\d{14}\.csv$`
	CSV_FILE              = "*.csv"
)
