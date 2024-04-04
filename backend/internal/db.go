package internal

import (
	"context"
	"fmt"
	"os"

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

func QueryDB(measurement string) {
	// Get query client
	queryAPI := client.QueryAPI("my-org")
	// get QueryTableResult
	result, err := queryAPI.Query(context.Background(), `from(bucket:"my-bucket")|> range(start: -1y) |> filter(fn: (r) => r._measurement == "`+measurement+`")`)
	if err == nil {
		// Iterate over query response
		for result.Next() {
			// Notice when group key has changed
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// Access data
			fmt.Printf("time: %v | value: %v\n", result.Record().Time(), result.Record().Value())
		}
		// check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
	} else {
		panic(err)
	}
}

var client influxdb2.Client
