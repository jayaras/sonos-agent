package mqttclient

import (
	"log"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type MQTTClient struct {
	broker  string
	topic   string
	handler func(client MQTT.Client, msg MQTT.Message)
}

func (m *MQTTClient) onLost(client MQTT.Client, err error) {
	log.Print("MQTT Connection Lost")

}

func (m *MQTTClient) onConnect(client MQTT.Client) {
	log.Print("MQTT Connected.")
	if token := client.Subscribe(m.topic, 0, m.handler); token.Wait() && token.Error() != nil {
		log.Print(token.Error())
		log.Fatal("MQTT token error onConnect")
	}
}

func (m *MQTTClient) Run() {

	opts := MQTT.NewClientOptions().AddBroker(m.broker)
	opts.SetClientID("sonos-controller")
	opts.SetAutoReconnect(true)
	opts.SetConnectTimeout(30 * time.Second)
	//opts.SetKeepAlive(10 * time.Second)
	opts.SetConnectionLostHandler(m.onLost)
	opts.SetOnConnectHandler(m.onConnect)
	//opts.SetMaxReconnectInterval(1 * time.Second)

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

}

func NewMQTTClient(broker string, baseTopic string, device string, handler func(client MQTT.Client, msg MQTT.Message)) *MQTTClient {

	topic := baseTopic + "/" + device + "/rfid/uid"
	x := &MQTTClient{broker: broker, handler: handler, topic: topic}
	return x
}
