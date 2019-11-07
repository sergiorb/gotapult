package api

import (
  common "github.com/sergiorb/gotapult/common/api"
)

var ApiFunctions []common.ApiFunction

func init() {

  ApiFunctions = []common.ApiFunction{
    common.ApiFunction{
      Path:    "/http/*proxyPath",
      Handler:  HttpCatapultLaunch,
    },
  }
}
