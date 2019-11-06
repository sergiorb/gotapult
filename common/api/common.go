package api

import (
  "github.com/gin-gonic/gin"
)

type ApiFunction struct {
  Method  string
  Path    string
  Handler gin.HandlerFunc
}

func LoadApi(group *gin.RouterGroup, apiFunctions []ApiFunction) {

  for _, apiFunction := range apiFunctions {

    group.Handle(apiFunction.Method, apiFunction.Path, apiFunction.Handler)
  }
}
