package utils

import (
	"gopkg.in/ini.v1"
	"log"
)

var configFile string = "app/conf/development.ini"
/**
 * @Description: 获取配置项信息
 * @param section
 * @param key
 * @return string
 */
func GetConfig(section string, key string) string {
	cfg, err := ini.Load(configFile)
	if err != nil {
		log.Fatal(err)
	}
	return cfg.Section(section).Key(key).String()
}
