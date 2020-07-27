package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	model "github.com/matisszilard/devops-palinta/pkg/model"
	mqtt "github.com/matisszilard/devops-palinta/pkg/mqtt"
	util "github.com/matisszilard/devops-palinta/pkg/util"
)

var mqttChannel chan string

func main() {
	mqttChannel = make(chan string, 2)
	rand.Seed(time.Now().UTC().UnixNano())

	log.Info("Connect to the MQTT broker")
	client, err := mqtt.ConnectToMqttBroker(model.MqttBroker, mqttChannel)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Unable to connect to the MQTT broker")
		os.Exit(1)
	}

	id := util.RandStringBytes(10)

	for {
		s := strconv.FormatInt(rand.Int63n(35), 10)

		mqtt.Publish(client, model.MqttPrometheusTemperatureTopic, id+"/"+s)

		log.WithFields(log.Fields{
			"data": s,
		}).Info("Send data to client")

		time.Sleep(120 * time.Second)
	}
}
