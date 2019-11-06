package logger

import (
  "github.com/sergiorb/gotapult/config"
  log "github.com/sirupsen/logrus"
)

func Init() {

  if config.Store.Log.Development {

    log.SetFormatter(&log.TextFormatter{})

  } else {

    log.SetFormatter(&log.JSONFormatter{})
  }

  switch config.Store.Log.Level {
  case "PANIC":
    log.SetLevel(log.FatalLevel)
  case "FATAL":
    log.SetLevel(log.FatalLevel)
  case "ERROR":
    log.SetLevel(log.ErrorLevel)
  case "WARNING":
    log.SetLevel(log.WarnLevel)
  case "INFO":
    log.SetLevel(log.InfoLevel)
  default:
    log.SetLevel(log.DebugLevel)
    config.Store.Log.Level = "DEBUG"
  }

}
