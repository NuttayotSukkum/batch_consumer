package utils

import (
	"github.com/NuttayotSukkum/batch_consumer/internals/pkg/constants"
	logger "github.com/labstack/gommon/log"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func DeleteDirectory(path string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, f := range files {
		fPath := filepath.Join(path, f.Name())
		if f.IsDir() {
			if err := DeleteDirectory(fPath); err != nil {
				return err
			}
		}
		if err := os.Remove(fPath); err != nil {
			return err
		}
	}
	return nil
}

func EmptyInputDirectory() {
	dirPath := constants.DirPath

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
			logger.Errorf("Failed to create dir %s: %s", dirPath, err)
		} else {
			logger.Infof("Successfully dir %s", dirPath)
		}
	} else {
		if err := DeleteDirectory(dirPath); err != nil {
			logger.Errorf("Failed to delete dir %s: %s", dirPath, err)
		}
		if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
			logger.Errorf("Failed to create dir %s: %s", dirPath, err)
		} else {
			DeleteDirectory(dirPath)
			logger.Infof("Successfully dir %s", dirPath)
		}
	}
}

func IsFileInRange(key string) bool {
	if !strings.HasSuffix(key, constants.SUFFIX) {
		logger.Infof("File %s does not have .csv suffix. Skipping.", key)
		return false
	}
	location, err := time.LoadLocation(constants.TIME_ZONE)
	if err != nil {
		log.Fatalf("Failed to load timezone Asia/Bangkok: %v", err)
	}
	time.Local = location

	startDate := time.Now().In(location).AddDate(0, 0, -1).Truncate(24 * time.Hour)
	endDate := time.Now().In(location).Truncate(24 * time.Hour).Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	key = strings.TrimPrefix(key, constants.FILE_NAME)
	key = strings.TrimSuffix(key, constants.SUFFIX)
	fileDate, err := time.ParseInLocation(constants.DATE_FORMATTER, key, location)
	if err != nil {
		logger.Errorf("Failed to load timezone Asia/Bangkok: %v", err)
		return false
	}
	return !fileDate.Before(startDate) && !fileDate.After(endDate)
}

func TimeLocal(timer time.Time) time.Time {
	location, err := time.LoadLocation(constants.TIME_ZONE)
	if err != nil {
		log.Fatalf("Failed to load timezone %s: %v", constants.TIME_ZONE, err)
	}

	if timer.IsZero() {
		log.Fatalf("Provided time is zero time: %v", timer)
	}
	localTime := timer.In(location)
	return localTime
}

func SubString(value string, number int) string {
	if len(value) < number {
		return value
	}
	return value[:number]
}
