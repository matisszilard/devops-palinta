package influxdb

import (
	log "github.com/sirupsen/logrus"

	"github.com/matisszilard/devops-palinta/pkg/store"

	"github.com/influxdata/influxdb/client/v2"
)

// New - store instance is created
func New(host string, port string) store.Store {
	// TODO handle reconnection
	connection, _ := createConnection(host, port)
	log.Debug("Create database")
	createDatabase(connection, "demeter")

	return store.New("influxdb", &influxdbTemperatureStore{connection})
}

// createDatabase - create a given database
func createDatabase(connection client.Client, database string) {
	log.Info("Create database - [", database, "]")
	q := client.NewQuery("CREATE DATABASE "+database, "", "")
	if response, err := connection.Query(q); err == nil && response.Error() == nil {
		log.Info(response.Results)
	}
}

// createConnection - create a db connection to the given host and port
func createConnection(host string, port string) (client.Client, error) {
	log.Info("Create InlfuxDB connection")
	url := "http://" + host + ":" + port
	log.Info("Connecting to [" + url + "]")
	connection, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: url,
	})
	if err != nil {
		log.Error("Error creating InfluxDB Client: ", err.Error())
	} else {
		log.Info("InfluxDB connection created")
	}
	return connection, err
}

// closeConnection - close a db connection
func closeConnection(connection client.Client) {
	log.Info("Closing InfluxDB connection")
	connection.Close()
	log.Info("InfluxDB connection closed")
}
