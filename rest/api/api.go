package api

import (
  "net/http"
  common "github.com/sergiorb/gotapult/common/api"
)

var ApiFunctions []common.ApiFunction

func init() {

  ApiFunctions = []common.ApiFunction{
    common.ApiFunction{
      Method:  http.MethodGet,
      Path:    "/all",
      Handler:  GetAllRelays,
    },
  }
}
