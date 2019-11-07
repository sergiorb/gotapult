package catapults

import (
	"github.com/sergiorb/gotapult/catapults/config"
	"github.com/sergiorb/gotapult/catapults/http"
)

var Store Catapults

func Init(config config.Conf) {

	Store = Catapults{
		Http: http.Build(config.Http),
	}
}

type Catapults struct {

	Http	*http.Catapult
}
