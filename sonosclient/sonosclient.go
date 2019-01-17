package sonosclient

import (
	"errors"
	"log"

	sonos "github.com/ianr0bkny/go-sonos"
)

type SonosClient struct {
	name       string
	player     *sonos.Sonos
	retryCount int
	netif      string
}

func (s *SonosClient) Connect() error {

	log.Print("Starting Sonos Discovery...")
	for x := 0; x < s.retryCount; x++ {
		mgr, err := sonos.Discover(s.netif, "11223")
		if err != nil {
			log.Print("Error With Discovery: ")
			log.Fatal(err)
		}
		ss := sonos.ConnectAny(mgr, nil, sonos.SVC_ALL)

		for _, y := range ss {
			a, _, err := y.GetZoneAttributes()

			if err != nil {
				panic(err)
			}

			if a == s.name {
				log.Print("Found Player: " + s.name)
				s.player = y
				return nil
			}

		}
		mgr.Close()
	}

	if s.player == nil {
		log.Fatal("No Play Found.")
		return errors.New("Could not find Sonos")
	}

	return nil

}

func (s *SonosClient) GetSonosPlayer() *sonos.Sonos {
	return s.player
}

func NewSonosClient(name string, netif string, retryCount int) *SonosClient {
	return &SonosClient{name: name, netif: netif, retryCount: retryCount}

}
