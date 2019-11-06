package relays

import (
  "fmt"
  "github.com/sergiorb/gotapult/scheduler"
  "github.com/sergiorb/gotapult/relays/data"
  log "github.com/sirupsen/logrus"
)

func Init() {

  for _, relayEvent := range data.ReadRelayEvents() {

    log.Debug(fmt.Sprintf("RelayEvent - ID: %v, LoadOnBoot: %v", relayEvent.Id, relayEvent.LoadOnBoot))

    if relayEvent.LoadOnBoot {

        scheduler.Store.AddFunction(relayEvent.Cron, relayEvent.Run)
    }
  }
}
