package api

import (
  "github.com/gin-gonic/gin"
  "github.com/sergiorb/gotapult/catapults"
)

func HttpCatapultLaunch(c *gin.Context) {

  catapults.Store.Http.Launch(c)
}
