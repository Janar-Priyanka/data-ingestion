package db

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type ControllersStruct struct {
}

func GetDayData_Db(c *gin.Context, ctx context.Context, day int, pg *pgx.Conn) []Data {

	query := `
        SELECT timestamp, cpu_load, concurrency 
        FROM comcast 
        WHERE timestamp >= NOW() - ($1 * INTERVAL '1 day')
        ORDER BY timestamp DESC
		`
	rows, err := pg.Query(ctx, query, day)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query failed for Fetcing past Day Data"})
		return nil
	}

	var result []Data
	result, err = pgx.CollectRows(rows, pgx.RowToStructByPos[Data])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error in converting rows to struct"})
		return nil
	}

	return result
}

// GetHoursData
func GetHoursData_Db(c *gin.Context, ctx context.Context, hours int, pg *pgx.Conn) []Data {
	query := `
        SELECT timestamp, cpu_load, concurrency 
        FROM comcast 
        WHERE timestamp >= NOW() - ($1 * INTERVAL '1 hour')
        ORDER BY timestamp DESC
		`
	rows, err := pg.Query(ctx, query, hours)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query failed for Fetcing past Hour Data"})
		return nil
	}
	var result []Data
	result, err = pgx.CollectRows(rows, pgx.RowToStructByPos[Data])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error in converting rows to struct"})
		return nil
	}
	return result
}
