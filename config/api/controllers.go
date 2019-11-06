package api

import (
  "github.com/gin-gonic/gin"
  theConfig "github.com/sergiorb/gotapult/config"
)

func GetStore(c *gin.Context) {

  c.JSON(200, gin.H{
    "object": theConfig.Store,
  })
}
