package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

type Data struct {
	Timestamp   time.Time `json:"timestamp"`
	CPULoad     float64   `json:"cpu_load"`
	Concurrency int       `json:"concurrency"`
}

func IngestData(pgInstance *pgx.Conn, ctx context.Context, data *Data) error {
	insertSQL := `INSERT INTO comcast (timestamp, cpu_load, concurrency) VALUES ($1, $2, $3)`

	_, err := pgInstance.Exec(ctx, insertSQL, data.Timestamp, data.CPULoad, data.Concurrency)
	if err != nil {
		// If one insert fails, we log it and stop the process.
		return fmt.Errorf("failed to insert data for timestamp %v: %w", data.Timestamp, err)
	}
	return nil

}
func GenerateData(pgInstance *pgx.Conn, ctx context.Context) []*Data {
	currentTime := time.Now()
	res := []*Data{}

	for i := 0; i < 300; i++ {
		data := &Data{
			Timestamp:   currentTime.Add(time.Second * -1),
			CPULoad:     math.Round((rand.Float64()*100.0)*100) / 100,
			Concurrency: rand.IntN(500001), //0-50000
		}
		err := IngestData(pgInstance, ctx, data)
		if err != nil {
			log.Fatalf("Fatal error: could not ingest data : %v", err)
		}
		currentTime = currentTime.Add(time.Second * -1)
		res = append(res, data)
	}

	return res
}

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:password@localhost:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	data := GenerateData(conn, ctx)
	for _, val := range data {
		fmt.Println(" val : ", *val)
	}

}
