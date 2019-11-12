package http

import (
	"fmt"
	"time"
	"net/url"
	_ "strings"
	"net/http"
	_ "context"
	"net/http/httputil"
	"github.com/gin-gonic/gin"
	"github.com/sergiorb/gotapult/catapults/http/config"
	"github.com/sergiorb/gotapult/catapults/http/event"
	"github.com/sergiorb/gotapult/events"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/sergiorb/gotapult/data"
	_ "github.com/sirupsen/logrus"
)

type Catapult struct {
	Targets 		map[string]Target
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

func ErrHandle(res http.ResponseWriter, req *http.Request, err error) {

	_, err = data.Store.HttpEvents.InsertOne(data.Store.Ctx(), bson.M{
		"timestamp": time.Now().UTC(), 
		"err": err.Error(),
	})

  	if err != nil {
    	fmt.Println(err)
  	}
}

func Build(config config.Conf) *Catapult {

	targets := make(map[string]Target)
 
	for k, v := range config.Targets {

		url, err := url.Parse(fmt.Sprintf("%v://%v:%d", v.Schema, v.Host, v.Port))
		if err != nil { panic(err) }

		client := httputil.NewSingleHostReverseProxy(url)

		client.ErrorHandler = func (res http.ResponseWriter, req *http.Request, err error) {

			id := req.Header.Get("x-gotapult-id")

			if id == "" {
				id = "default"
			}

			_, err = data.Store.HttpEvents.InsertOne(data.Store.Ctx(), event.HttpEvent{
				Timestamp:	time.Now().UTC(),
				Err:		err.Error(),
				Target:		events.Target{
					Id:		id,
				},
			})
		
			if err != nil {
				fmt.Println(err)
			}
		}

		targets[k] = Target{
			Id: k,
			Schema: v.Schema,
			Host: v.Host,
			Port: v.Port,
			Client: client,
		}
	}
	
	return &Catapult{
		Targets: targets,
	}
}
