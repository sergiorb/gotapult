package http

import (
	"fmt"
	"net/url"
	_ "strings"
	_ "net/http"
	"net/http/httputil"
	"github.com/gin-gonic/gin"
	"github.com/sergiorb/gotapult/catapults/http/config"
	_ "github.com/sirupsen/logrus"
)

type Catapult struct {
	Targets map[string]Target
}

type Target struct {
	Id		string
	Schema	string
	Host	string
	Port	uint
	Client	*httputil.ReverseProxy
}

func (c* Catapult) GetTargetOrDefault(id string) *Target {
	
	if target, ok := c.Targets[id]; ok {

    	return &target;
	}

	defaultTaget := c.Targets["default"]

	return &defaultTaget;
}

func (c *Catapult) Launch(gc *gin.Context) {

	id := gc.GetHeader("x-gotapult-id")

	path, err := url.Parse(gc.Param("proxyPath"))
	if err != nil { panic(err) }

	gc.Request.URL.Path = path.Path

	target := c.GetTargetOrDefault(id)

	target.Client.ServeHTTP(gc.Writer, gc.Request)
}

func Build(config config.Conf) *Catapult {

	targets := make(map[string]Target)
 
	for k, v := range config.Targets {

		url, err := url.Parse(fmt.Sprintf("%v://%v:%d", v.Schema, v.Host, v.Port))
		if err != nil { panic(err) }

		targets[k] = Target{
			Id: k,
			Schema: v.Schema,
			Host: v.Host,
			Port: v.Port,
			Client: httputil.NewSingleHostReverseProxy(url),
		}
	}
	
	return &Catapult{
		Targets: targets,
	}
}
