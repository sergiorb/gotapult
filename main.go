package main

import (
  "flag"
  "github.com/gin-gonic/gin"
  "github.com/sergiorb/gotapult/config"
  "github.com/sergiorb/gotapult/logger"
  "github.com/sergiorb/gotapult/catapults"

  log         "github.com/sirupsen/logrus"
  apiCommon   "github.com/sergiorb/gotapult/common/api"
  apiConfig   "github.com/sergiorb/gotapult/config/api"
  apiRest     "github.com/sergiorb/gotapult/rest/api"
)

const (
  release = "v0.1-alpha"
)

func init() {

  configPath := flag.String("configPath", "./config.json", "A string defining where the config file is located.")

  flag.Parse()
  config.Init(*configPath)
  logger.Init()
  catapults.Init(config.Store.Catapults)
}

func main() {

  log.WithFields(log.Fields{
    "release": release,
  }).Info("Starting Gotapult...")

  router  := gin.Default()
  base    := router.Group("/v1")

  configRouter  := base.Group("/config")
  apiRestRouter := base.Group("")
  
  apiCommon.LoadApi(configRouter,   apiConfig.ApiFunctions)
  apiCommon.LoadApi(apiRestRouter,  apiRest.ApiFunctions)

  log.Info("Runing...")

  // router.RunUnix(unixSocketPath)
  router.Run()
}
