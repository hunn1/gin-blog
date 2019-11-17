package config

import (
	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	config := Config{
		Name: cfg,
	}
	if err := config.initConfig(); err != nil {
		return err
	}
	config.initLog()
	config.watchConfig()
	return nil
}

/**
初始化配置文件
*/
func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("SITE_CONF")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// 监控配置文件变更
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		//log.Infof("Config file Changed: %s", in.Name)
	})
}

func (c *Config) initLog() {
	//passLagerCfg := log.PassLagerCfg{
	//	Writers:        viper.GetString("log.writers"),
	//	LoggerLevel:    viper.GetString("log.logger_level"),
	//	LoggerFile:     viper.GetString("log.logger_file"),
	//	LogFormatText:  viper.GetBool("log.log_format_text"),
	//	RollingPolicy:  viper.GetString("log.rollingPolicy"),
	//	LogRotateDate:  viper.GetInt("log.log_rotate_date"),
	//	LogRotateSize:  viper.GetInt("log.log_rotate_size"),
	//	LogBackupCount: viper.GetInt("log.log_backup_count"),
	//}
	//log.InitWithConfig(&passLagerCfg)
}
