package controller

import (
  "strconv"

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

func GetConfigInt(keyName string) (val int) {
  s := ConfigCache[keyName]
  if s == "" {
    s = cfg.Section("config").Key(keyName).String()
    ConfigCache[keyName] = s
  }
  val, _ = strconv.Atoi(ConfigCache[keyName])
  return
}

func GetConfigBool(keyName string) (val bool) {
  s := ConfigCache[keyName]
  if s == "" {
    s = cfg.Section("config").Key(keyName).String()
    ConfigCache[keyName] = s
  }
  if s == "1" || s == "true" || s == "TRUE" {
    val = true
  }
  return
}
