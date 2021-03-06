package main

import (
	"log"
	"time"

	"github.com/ianr0bkny/go-sonos"
	"github.com/jayaras/sonos-agent/config"
	"github.com/jayaras/sonos-agent/mqttclient"
	"github.com/jayaras/sonos-agent/songdb"
	"github.com/jayaras/sonos-agent/sonosclient"
)

var player *sonos.Sonos
var lastSong string

func uidHandler(blockID string) {

	db := songdb.NewSongDB("songblocks.json")

	if blockID == "" {
		// stop the song
		log.Print("Pausing Song")
		player.Pause(0)
	} else {

		log.Print("Looking up Block's Song: " + blockID)

		foundSong, err := db.Lookup(blockID)

		if err != nil {
			log.Println("Could not find song for tag" + blockID)

			sinfo, err := player.GetTransportInfo(0)
			if err != nil {
				log.Print("Error fetching current Transport Info:" + err.Error())
			}

			if sinfo.CurrentTransportState == "PLAYING" {
				pinfo, err := player.GetPositionInfo(0)
				if err != nil {
					log.Print("Error Fetching Current Position info: ")
				} else {
					db.Save(blockID, pinfo.TrackURI)
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
			// log.Println("Block was removed and placed back... resuming song.")
			player.Play(0, "1")
			lastSong = foundSong
		}

	}
}

func main() {

	log.Print("Loading Configuration...")

	c := config.Config{}

	netif := c.GetString("interface", "wlp58s0")
	playerName := c.GetString("player", "Living Room")
	mqttConn := c.GetTCPConnection("mqtt_server", "hass.local:1883")
	mqttBaseTopic := c.GetString("mqtt_base_topic", "homie")
	nodeName := c.GetString("node_name", "song-block")
	retryCount := c.GetInt("sonos_retry_count", 4)

	sonosClient := sonosclient.NewSonosClient(playerName, netif, retryCount)
	sonosClient.Connect()
	player = sonosClient.GetSonosPlayer()

	mqttclient := mqttclient.NewMQTTClient(mqttConn, mqttBaseTopic, nodeName, uidHandler)
	mqttclient.Run()

	for {
		time.Sleep(1 * time.Second)
	}

}
