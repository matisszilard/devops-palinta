package main

import (
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/matisszilard/devops-palinta/palinta/pkg/model"
	"github.com/matisszilard/devops-palinta/palinta/pkg/mqtt"
	"github.com/matisszilard/devops-palinta/palinta/pkg/store"
	"github.com/matisszilard/devops-palinta/palinta/pkg/store/influxdb"
)

var dbStore store.Store

func main() {
	log.Info("Connect to Influx database...")
	dbStore = influxdb.New(model.InfluxDBHost, model.InfluxDBPort)
	log.Info("Connection created to Influx database!")

	var mqttChannel chan string
	mqttChannel = make(chan string, 2)

	log.Info("Connect to the MQTT broker")
	client, err := mqtt.ConnectToMqttBroker(model.MqttBroker, mqttChannel)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Unable to connect to the MQTT broker")
		os.Exit(1)
	}
	err = mqtt.Subscribe(client, model.MqttPrometheusTemperatureTopic)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Unable to subscribe for tokens")
		os.Exit(1)
	}

	for {
		topic := <-mqttChannel
		payload := <-mqttChannel

		switch topic {
		case model.MqttPrometheusTemperatureTopic:
			{
				saveTemperature(payload)
			}
	}
}

func saveTemperature(temperature string) {
	var temp model.Temperature
	lt := time.Now()
	temp.Time = lt.String()

	d := strings.SplitN(temperature, "/", -1)

	t, err := strconv.ParseFloat(d[1], 64)
	if err != nil {
		t = 0.0
	}
	temp.Temperature = t
	temp.PhotonID = d[0]

	dbStore.Temperatures().Save(temp)
}
