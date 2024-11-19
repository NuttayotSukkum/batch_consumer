package configs

var (
	AppName string
	Version string
)

func LoadAppConfig(cfg AppConfig) {
	AppName = cfg.AppName
	Version = cfg.Version
}
