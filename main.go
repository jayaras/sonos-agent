package main

import (
	"log"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/ianr0bkny/go-sonos"
	"github.com/jayaras/sonos-agent/mqttclient"
	"github.com/jayaras/sonos-agent/songdb"
)

var player *sonos.Sonos
var lastSong string

func uidHandler(client MQTT.Client, msg MQTT.Message) {

	db := songdb.NewSongDB("songblocks.json")

	if string(msg.Payload()) == "" {
		// stop the song
		log.Print("Pausing Song")
		player.Pause(0)
	} else {

		log.Print("Looking up Block's Song: " + string(msg.Payload()))

		foundSong, err := db.Lookup(string(msg.Payload()))

		if err != nil {
			log.Println("Could not find song for tag" + string(msg.Payload()))

			sinfo, err := player.GetTransportInfo(0)
			if err != nil {
				log.Print("Error fetching current Transport Info:" + err.Error())
			}

			if sinfo.CurrentTransportState == "PLAYING" {
				pinfo, err := player.GetPositionInfo(0)
				if err != nil {
					log.Print("Error Fetching Current Position info: ")
				} else {
					db.Save(string(msg.Payload()), pinfo.TrackURI)
				}
			}

		} else {
			log.Println("Found Song:" + foundSong)

			pinfo, err := player.GetPositionInfo(0)
			if err != nil {
				log.Print("Error Fetching Current Position info: ")
			}
			log.Print("Currently on the sonos: " + pinfo.TrackURI)
			if lastSong != foundSong || lastSong != pinfo.TrackURI {

				log.Println("New Block, updating current song on sonos.")
				player.SetAVTransportURI(0, foundSong, "")
			}
			// resume current song or start it over if the song went away and you want an 'again'
			//log.Println("Block was removed and placed back... resuming song.")
			player.Play(0, "1")
			lastSong = foundSong
		}

	}
}

func main() {

	log.Print("Starting Sonos Discovery...")
	mgr, err := sonos.Discover("wlp58s0", "11223")
	if err != nil {
		log.Print("Error With Discovery: ")
		log.Fatal(err)
	}

	ss := sonos.ConnectAny(mgr, nil, sonos.SVC_ALL)

	for _, s := range ss {
		a, _, err := s.GetZoneAttributes()

		if err != nil {
			panic(err)
		}

		if a == "Living Room" {
			log.Print("Found Player: ")
			player = s
		}

	}

	if player == nil {
		log.Fatal("No Play Found.")
	}

	client := mqttclient.NewMQTTClient("tcp://hass.local:1883", "homie", "ABCD", uidHandler)

	client.Run()

	for {
		time.Sleep(1 * time.Second)
	}

}