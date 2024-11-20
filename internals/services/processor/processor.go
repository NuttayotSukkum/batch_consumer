package processor

import (
	"github.com/NuttayotSukkum/batch_consumer/internals/pkg/constants"
	logger "github.com/labstack/gommon/log"
)

func Processor() error {
	logger.Infof("Starting to process files from directory: %s", constants.DirPath)
	products, err := ReadFileInDirectory(constants.DirPath, 2)
	if err != nil {
		logger.Errorf("Invalid to read file csv err :%s", err)
		return err
	}
	logger.Infof("products:%+v", products)
	return nil
}
