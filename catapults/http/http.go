package http

import (
	"fmt"
	"net/http"
	"net/url"
	"net/http/httputil"
	"github.com/gin-gonic/gin"
	"github.com/sergiorb/gotapult/catapults/http/config"
)

func Build(config config.Conf) *Catapult {

	targets := make(map[string]Target)
 
	for k, v := range config.Targets {

		targets[k] = Target{
			Id: k,
			Schema: v.Schema,
			Host: v.Host,
			Port: v.Port,
			Url: fmt.Sprintf("%v://%v:%d", v.Schema, v.Host, v.Port),
			Client: http.Client{},
		}
	}
	
	return &Catapult{
		Targets: targets,
	}
}

type Catapult struct {
	Targets map[string]Target
}

type Target struct {
	Id			string
	Schema	string
	Host		string
	Port		uint
	Url			string
	Client 	http.Client
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
	proxyPath := gc.Param("proxyPath")

	target := c.GetTargetOrDefault(id)

	fmt.Printf("%v%v\n", target.Url, proxyPath)

	u, err := url.Parse(target.Url)
	if err != nil { panic(err) }

	_, err = url.Parse(proxyPath)
	if err != nil { panic(err) }

	proxy := httputil.NewSingleHostReverseProxy(u)

	fmt.Printf("%v\n", gc.Request.URL)

	// gc.Request.URL = e

	proxy.ServeHTTP(gc.Writer, gc.Request)
}
