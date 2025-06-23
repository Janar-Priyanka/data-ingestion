package utils

import (
	"context"
	"data-ingestion/cmd/server/db"
	"data-ingestion/cmd/server/models"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func HandleAggregateQuery(pg *pgx.Conn, c *gin.Context, ctx context.Context, requestBody models.GetSpecificDataSetRequest, startTime, endTime time.Time) (*models.GetSpecificDataSetAggResponse, error) {

	var columnName string
	switch requestBody.OppParam {
	case "cpu_load":
		columnName = "cpu_load"
	case "concurrency":
		columnName = "concurrency"
	default:
		//c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'opp_param'. Must be 'cpu_load' or 'concurrency'."})
		return nil, fmt.Errorf(`error:Invalid opp_param. Must be cpu_load or concurrency`)
	}

	var query string
	var operation string

	switch requestBody.OppCode {
	case "max":
		operation = "MAX"
	case "avg":
		operation = "AVG"
	default:
		//c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'oppcode'. Currently supporting only 'max' or 'avg'."})
		return nil, fmt.Errorf(`error:Invalid oppcode. Currently supporting only MAX or AVG `)
	}

	query = fmt.Sprintf(`SELECT %s(%s) FROM comcast WHERE timestamp >= $1 AND timestamp < $2`, operation, columnName)

	var result float64
	err := pg.QueryRow(ctx, query, startTime, endTime).Scan(&result)
	if err != nil {
		if err == pgx.ErrNoRows {
			//c.JSON(http.StatusOK, gin.H{"message": "No data found for the given time frame", "value": nil})
			return nil, fmt.Errorf(`error: No data found for the given time frame`)
		}
		log.Printf("Query failed for GetSpecificDataSet (aggregate): %v\n", err)
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed"})
		return nil, fmt.Errorf(`error: Database query failed`)
	}
	response := &models.GetSpecificDataSetAggResponse{
		Operation: requestBody.OppCode,
		Parameter: requestBody.OppParam,
		Value:     result,
	}
	return response, nil
}

func HandleRawDataQuery(pg *pgx.Conn, c *gin.Context, ctx context.Context, startTime, endTime time.Time) (*models.GetDataResponseStruct, error) {

	query := `
        SELECT timestamp, cpu_load, concurrency
        FROM comcast
        WHERE timestamp >= $1 AND timestamp < $2
        ORDER BY timestamp DESC
    `
	rows, err := pg.Query(ctx, query, startTime, endTime)
	if err != nil {
		log.Printf("Query failed for GetSpecificDataSet (raw): %v\n", err)
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed"})
		return nil, fmt.Errorf(`error: Database query failed`)
	}
	defer rows.Close()

	var result []db.Data
	result, err = pgx.CollectRows(rows, pgx.RowToStructByPos[db.Data])
	if err != nil {
		log.Printf("GetSpecificDataSet: Error converting rows to struct: %v\n", err)
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process database results"})
		return nil, fmt.Errorf(`error: Failed to process database results`)
	}

	response := &models.GetDataResponseStruct{}
	var responseData []models.Data
	if len(result) == 0 {
		response.Data = responseData // Return an empty slice
		return response, nil
	}

	for _, val := range result {
		data := models.Data{
			Timestamp:   val.Timestamp.Unix(),
			CPULoad:     val.CPULoad,
			Concurrency: val.Concurrency,
		}
		responseData = append(responseData, data)
	}
	response.Data = responseData
	return response, nil
	// response.Data = responseData
	// c.JSON(http.StatusOK, response)
}
