package api

import (
  "github.com/gin-gonic/gin"
)

func GetAllRelays(c *gin.Context) {

  c.JSON(200, gin.H{
    "objects": "",
  })
}
