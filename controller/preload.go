package controller

import (
  "github.com/Cyvadra/binance-local/api"
  "github.com/go-ini/ini"
)

var (
  filepath    string = "config/config.ini"
  cfg         *ini.File
  ConfigCache map[string]string
)

func init() {
  var err error
  cfg, err = ini.Load(filepath)
  ConfigCache = make(map[string]string)
  if err != nil {
    panic("找不到配置文件")
  }
  api.SetAPIKey(GetConfig("API_KEY"))
  api.SetSecretKey(GetConfig("SECRET_KEY"))
  api.Init()
  return
}

func GetConfig(keyName string) (val string) {
  if ConfigCache[keyName] == "" {
    val = cfg.Section("config").Key(keyName).String()
    ConfigCache[keyName] = val
  } else {
    val = ConfigCache[keyName]
  }
  return
}
