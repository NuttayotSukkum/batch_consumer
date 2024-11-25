package configs

import (
	"context"
	logger "github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"os"
	"strings"
	"sync"
)

type Config struct {
	App           AppConfig     `mapstructure:"app"`
	Log           Log           `mapstructure:"log"`
	Secrets       Secrets       `mapstructure:"secrets"`
	KafkaProducer KafkaProducer `mapstructure:"kafka_producer"`
}

type AppConfig struct {
	AppName         string `mapstructure:"name"`
	Port            string `mapstructure:"port"`
	Version         string `mapstructure:"version"`
	env             string `mapstructure:"env"`
	ChunkSize       int    `mapstructure:"chunk-size"`
	ChunkSizeReader int    `mapstructure:"chunk-size-reader"`
}

type Log struct {
	Env   string `mapstructure:"env"`
	Level string `mapstructure:"level"`
}

type Secrets struct {
	Host      string     `mapstructure:"cloud-sql-gormhost"`
	Port      string     `mapstructure:"cloud-sql-port"`
	Username  string     `mapstructure:"cloud-sql-username"`
	Password  string     `mapstructure:"cloud-sql-password"`
	DBName    string     `mapstructure:"cloud-sql-dbname"`
	AWSSecret AWSSecrets `mapstructure:"aws"`
}

type AWSSecrets struct {
	AccessKey string `mapstructure:"access-key"`
	SecretKey string `mapstructure:"secret-key"`
	S3        S3     `mapstructure:"S3"`
}

type S3 struct {
	S3Bucket  string `mapstructure:"bucket-name"`
	BucketArn string `mapstructure:"bucket-arn"`
	Region    string `mapstructure:"region"`
}

type KafkaProducer struct {
	Version    string              `mapstructure:"version"`
	KafkaSASAL KafkaSASAL          `mapstructure:"sasl"`
	KafkaTLS   KafkaTLS            `mapstructure:"tls"`
	Producer   ProducerConfig      `mapstructure:"producer"`
	Topics     KafkaProducerTopics `mapstructure:"topics"`
}

type KafkaSASAL struct {
	Enable    bool   `mapstructure:"enable"`
	Mechanism string `mapstructure:"mechanism"`
}

type KafkaTLS struct {
	Enable bool `mapstructure:"enable"`
}

type ProducerConfig struct {
	Partitioner string `mapstructure:"partitioner"`
}

type KafkaProducerTopics struct {
	Topic string `mapstructure:"topic"`
}

var (
	config     *Config
	configOnce sync.Once
)

func InitConfig(ctx context.Context) *Config {
	configOnce.Do(func() {
		configPath, ok := os.LookupEnv("CONFIG_PATH")
		if !ok {
			logger.Errorf("%v :API_CONFIG_PATH not found, using default config", ctx)
			configPath = "../configs"
		}
		configName, ok := os.LookupEnv("CONFIG_NAME")
		if !ok {
			logger.Infof("%v: API_CONFIG_NAME not found, using default config", ctx)
			configName = "config"
		}
		viper.AddConfigPath(configPath)
		viper.SetConfigName(configName)
		viper.SetConfigType("yaml")

		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		if err := viper.ReadInConfig(); err != nil {
			logger.Infof("%v: config file not found. using default/env config: %s", ctx, err)
		}
		viper.AutomaticEnv()
		config = &Config{}
		if err := viper.Unmarshal(&config); err != nil {
			logger.Errorf("%v :unable to decode struct: %s", ctx, err)
		}
	})
	return config
}
