# sonos-agent
## description
Basic agent manager that listens for RFID or NFC uids over MQTT and either plays a song or assoicates the current playing song with a uid not in the database.  Inspired by  https://shawnrk.github.io/songblocks/ and requires an NFC or RFID reader that talks MQTT (such as: https://github.com/jayaras/esp8266-nfc-reader).


## Running
The following Environment variables need to be set.

```bash

INTERFACE # the network interface to listen on for Sonos Discovery.
PLAYER # the name of the player eg "Living Room"
MQTT_SERVER # host:port for MQTT Broker
MQTT_BASE_TOPIC # assumes homie spec for the full topic so 'homie' by default
NODE_NAME # what you named your homie node when you configured it
RETRY_COUNT # how many times to retry to find your sonos.
```

Example:
```
$ INTERFACE=wlan0 PLAYER="Kitchen" MQTT_SERVER=localhost:1883 NODE_NAME=song-block ./sonos-agent
2019/01/16 16:01:25 Loading Configuration...
2019/01/16 16:01:25 Starting Sonos Discovery...
2019/01/16 16:01:25 Invalid server description `Synology/DSM/192.168.5.17'
2019/01/16 16:01:28 Invalid server description `Synology/DSM/192.168.5.17'
2019/01/16 16:01:31 Loading http://192.168.5.13:1400/xml/device_description.xml
2019/01/16 16:01:31 Loading http://192.168.5.13:1400/xml/ZoneGroupTopology1.xml
2019/01/16 16:01:31 Loading http://192.168.5.13:1400/xml/ContentDirectory1.xml
2019/01/16 16:01:31 Loading http://192.168.5.13:1400/xml/ConnectionManager1.xml
2019/01/16 16:01:31 Loading http://192.168.5.13:1400/xml/SystemProperties1.xml
2019/01/16 16:01:31 Loading http://192.168.5.13:1400/xml/MusicServices1.xml
2019/01/16 16:01:31 Loading http://192.168.5.13:1400/xml/DeviceProperties1.xml
2019/01/16 16:01:31 Loading http://192.168.5.13:1400/xml/GroupManagement1.xml
2019/01/16 16:01:31 Loading http://192.168.5.13:1400/xml/RenderingControl1.xml
2019/01/16 16:01:31 Loading http://192.168.5.13:1400/xml/AVTransport1.xml
2019/01/16 16:01:31 Loading http://192.168.5.13:1400/xml/AlarmClock1.xml
2019/01/16 16:01:31 Invalid server description `Synology/DSM/192.168.5.17'
2019/01/16 16:01:34 Invalid server description `Synology/DSM/192.168.5.17'
2019/01/16 16:01:37 Loading http://192.168.5.14:1400/xml/device_description.xml
2019/01/16 16:01:37 Loading http://192.168.5.14:1400/xml/DeviceProperties1.xml
2019/01/16 16:01:37 Loading http://192.168.5.14:1400/xml/SystemProperties1.xml
2019/01/16 16:01:37 Loading http://192.168.5.14:1400/xml/ZoneGroupTopology1.xml
2019/01/16 16:01:37 Loading http://192.168.5.14:1400/xml/GroupManagement1.xml
2019/01/16 16:01:37 Loading http://192.168.5.14:1400/xml/ContentDirectory1.xml
2019/01/16 16:01:37 Loading http://192.168.5.14:1400/xml/ConnectionManager1.xml
2019/01/16 16:01:37 Loading http://192.168.5.14:1400/xml/AlarmClock1.xml
2019/01/16 16:01:37 Loading http://192.168.5.14:1400/xml/MusicServices1.xml
2019/01/16 16:01:37 Loading http://192.168.5.14:1400/xml/RenderingControl1.xml
2019/01/16 16:01:37 Loading http://192.168.5.14:1400/xml/AVTransport1.xml
2019/01/16 16:01:37 Found Player: 
2019/01/16 16:01:38 MQTT Connected.

```

