package models

import (
	"time"

	"github.com/jackc/pgx/v5"
)

type Data struct {
	Timestamp   int64   `json:"timestamp"`
	CPULoad     float64 `json:"cpu_load"`
	Concurrency int     `json:"concurrency"`
}
type GetDataResponseStruct struct {
	Data []Data `json:"data"`
}

type GetSpecificDataSetRequest struct {
	StartTime time.Time `json:"starttime" binding:"required"`
	EndTime   time.Time `json:"endtime" binding:"required"`
	OpCode    string    `json:"opcode"`
	Params    string    `json:"params"`
}

type GetSpecificDataSetAggResponse struct {
	Operation string  `json:"operation"`
	Parameter string  `json:"parameter"`
	Value     float64 `json:"value"`
}

type ServiceStruct struct {
	Db *pgx.Conn
}
