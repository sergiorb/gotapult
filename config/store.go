package config

import (
  "os"
  "io/ioutil"
  "encoding/json"
  catapults "github.com/sergiorb/gotapult/catapults/config"
  data "github.com/sergiorb/gotapult/data/config"
)

var Store ConfigStore

type ConfigStore struct {
  Log       ConfigLogStore  `json:"log"`
  Catapults catapults.Conf  `json:"catapults`
  Data      data.Conf       `json:"data"`
}

type ConfigLogStore struct {
  Level       string  `json:"level"`
  Development bool    `json:"development"`
}

func Init(path string) {

  Store = ConfigStore{}

  if _, err := os.Stat(path); os.IsNotExist(err) {

    panic("err")

  } else {

    Store.load(path)
  }
}

func (c *ConfigStore) load(path string) {

  b, err := ioutil.ReadFile(path)

  if err != nil {
    panic(err)
  }

  if err := json.Unmarshal(b, &Store); err != nil {
    panic(err)
  }
}
