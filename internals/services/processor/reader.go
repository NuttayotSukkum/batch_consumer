package processor

import (
	"encoding/csv"
	"github.com/NuttayotSukkum/batch_consumer/internals/models/dao"
	"github.com/NuttayotSukkum/batch_consumer/internals/pkg/constants"
	"github.com/NuttayotSukkum/batch_consumer/internals/pkg/utils"
	logger "github.com/labstack/gommon/log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

func ReadFileInDirectory(directory string, chunkSize int) ([][]dao.Product, error) {
	var allChunks [][]dao.Product
	files, err := filepath.Glob(filepath.Join(directory, constants.CSV_FILE))
	if err != nil {
		return nil, err
	}
	regex := regexp.MustCompile(constants.REGEX_FILE_NAME)

	for _, file := range files {
		_, fileName := filepath.Split(file)
		if !regex.MatchString(fileName) {
			continue
		}
		products, err := readCSV(file, fileName)
		if err != nil {
			return nil, err
		}

		chunks := chunkProduct(*products, chunkSize)
		allChunks = append(allChunks, chunks...)
	}
	return allChunks, nil
}

func readCSV(filePath, fileName string) (*[]dao.Product, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '|'
	records, err := reader.ReadAll()
	if err != nil {
		logger.Errorf("Error to read All file:%s", err)
		return nil, err
	}

	var products []dao.Product
	for _, record := range records {
		if len(record[0]) == 0 || len(record[0]) == 200 {
			logger.Warnf("File: %s Invalid ID format: %s", fileName, record[0])
			continue
		}
		amount, err := strconv.Atoi(record[2])
		if err != nil {
			logger.Warnf("File: %s Invalid amount format: %s", fileName, record[2])
			continue
		}
		price, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			logger.Warnf("File: %s Invalid price format: %s", fileName, record[3])
			continue
		}

		var created time.Time
		createdStr := record[4]
		if len(createdStr) == 14 {
			created, err = time.ParseInLocation(constants.DATE_FORMATTER, createdStr, constants.TIME_LOCAL)
			if err != nil {
				logger.Warnf("Invalid date format: %s error:%s", createdStr, err)
				continue
			}
		} else {
			logger.Warnf("Invalid date lenght: %s", fileName)
			continue
		}

		product := dao.Product{ID: record[0], Name: record[1], Amount: amount, Price: price, CreatedAt: created, UpdatedAt: utils.TimeLocal(time.Now())}
		products = append(products, product)
	}
	logger.Infof("all product:%+v", products)
	return &products, nil
}

func chunkProduct(products []dao.Product, chunkSize int) [][]dao.Product {
	var chunks [][]dao.Product
	for i := 0; i < len(products); i += chunkSize {
		end := i + chunkSize
		if end > len(products) {
			end = len(products)
		}
		chunks = append(chunks, products[i:end])
	}
	return chunks
}
