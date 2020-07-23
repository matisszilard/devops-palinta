package influxdb

import (
	"time"

	log "github.com/sirupsen/logrus"

	client "github.com/influxdata/influxdb/client/v2"
	"github.com/matisszilard/devops-palinta/palinta/pkg/model"
)

type influxdbTemperatureStore struct {
	client client.Client
}

func (db *influxdbTemperatureStore) Save(temperature model.Temperature) {
	// Create a new point batch
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "demeter",
		Precision: "s",
	})

	// Create a point and add to batch
	tags := map[string]string{"photonId": temperature.PhotonID}
	fields := map[string]interface{}{
		"data":         temperature.Temperature,
		"published_at": temperature.Time,
		"coreId":       temperature.PhotonID,
	}
	log.Info("TemperatureStore - Save - ", fields)
	pt, err := client.NewPoint("temp", tags, fields, time.Now())
	if err != nil {
		log.Error("Error while creating a new point: ", err.Error())
	}
	bp.AddPoint(pt)
	if db.client != nil {
		db.client.Write(bp)
	}
}
