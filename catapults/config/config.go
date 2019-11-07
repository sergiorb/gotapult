package config

import (
	rabbitmqConf	"github.com/sergiorb/gotapult/catapults/rabbitmq/config"
	httpConf		"github.com/sergiorb/gotapult/catapults/http/config"
)

type Conf struct {
	Rabbitmq	rabbitmqConf.Conf	`json:"rabbitmq"`
	Http		httpConf.Conf		`json:"http"`
}
