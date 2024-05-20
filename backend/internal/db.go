package internal

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/influxdata/influxdb-client-go/v2"
)

func ConnectDB() {
	client = influxdb2.NewClient(os.Getenv("INFLUXDB_URL"), os.Getenv("DOCKER_INFLUXDB_INIT_ADMIN_TOKEN"))
	fmt.Println("Connected to InfluxDB")
}

func CloseDB() {
	client.Close()
	fmt.Println("Closed InfluxDB connection")
}

func QueryDB(measurement string) interface{} {
	// Define types for json response
	type response struct {
		Measurement string      `json:"measurement"`
		Field       string      `json:"field"`
		Labels      []time.Time `json:"labels"`
		Values      []float64   `json:"values"`
	}
	var res response

	// Get query client
	queryAPI := client.QueryAPI("my-org")
	// get QueryTableResult
	result, err := queryAPI.Query(context.Background(), `from(bucket:"my-bucket")|> range(start: -1y) |> filter(fn: (r) => r._measurement == "`+measurement+`")`)
	if err == nil {
		// Init response
		res.Labels = make([]time.Time, 0)
		res.Values = make([]float64, 0)
		// Iterate over query response
		for result.Next() {
			// Notice when group key has changed
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// Access data
			//fmt.Printf("time: %v | value: %v\n", result.Record().Time(), result.Record().Value())
			res.Measurement = result.Record().Measurement()
			res.Field = result.Record().Field()
			res.Labels = append(res.Labels, result.Record().Time())
			res.Values = append(res.Values, result.Record().Value().(float64))
		}
		// check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
	} else {
		panic(err)
	}

	return res
}

var client influxdb2.Client
