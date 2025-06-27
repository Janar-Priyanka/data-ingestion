package db

import "time"

type Data struct {
	Timestamp   time.Time `json:"timestamp"`
	CPULoad     float64   `json:"cpu_load"`
	Concurrency int       `json:"concurrency"`
}
type GetDataResponseStruct struct {
	Data []Data `json:"data"`
}
