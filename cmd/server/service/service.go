package service

import (
	"data-ingestion/cmd/server/db"
	"data-ingestion/cmd/server/models"
	"data-ingestion/cmd/server/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type ServiceStruct struct {
	Db *pgx.Conn
}

func (pg ServiceStruct) GetDayData(c *gin.Context) {
	day_stringVal := c.Param("day")

	day, err := strconv.Atoi(day_stringVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid day parameter, must be an integer"})
		return
	}
	if day <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid day parameter, must be an integer greater than zero"})
		return
	}

	ctx := c.Request.Context()

	query := `
        SELECT timestamp, cpu_load, concurrency 
        FROM comcast 
        WHERE timestamp >= NOW() - ($1 * INTERVAL '1 day')
        ORDER BY timestamp DESC
		`
	rows, err := pg.Db.Query(ctx, query, day)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed for Fetcing past Day Data : %v\n", err)
		os.Exit(1)
	}

	var result []db.Data
	result, err = pgx.CollectRows(rows, pgx.RowToStructByPos[db.Data])
	if err != nil {
		fmt.Errorf("GetDay : Error in converting rows to struct")
	}

	var responseData []models.Data
	for _, val := range result {
		data := models.Data{
			Timestamp:   val.Timestamp.Unix(),
			CPULoad:     val.CPULoad,
			Concurrency: val.Concurrency,
		}
		responseData = append(responseData, data)
	}

	response := models.GetDataResponseStruct{
		Data: responseData,
	}
	c.JSON(http.StatusOK, response)
}

func (pg ServiceStruct) GetHoursData(c *gin.Context) {
	hours_stringVal := c.Param("hour")

	hours, err := strconv.Atoi(hours_stringVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hours parameter, must be an integer"})
		return
	}
	if hours <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hours parameter, must be an integer greater than zero"})
		return
	}

	ctx := c.Request.Context()

	query := `
        SELECT timestamp, cpu_load, concurrency 
        FROM comcast 
        WHERE timestamp >= NOW() - ($1 * INTERVAL '1 hour')
        ORDER BY timestamp DESC
		`
	rows, err := pg.Db.Query(ctx, query, hours)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed for Fetcing past Hour Data : %v\n", err)
		os.Exit(1)
	}

	var result []db.Data
	result, err = pgx.CollectRows(rows, pgx.RowToStructByPos[db.Data])
	if err != nil {
		fmt.Errorf("GetHourData : Error in converting rows to struct")
	}
	var response models.GetDataResponseStruct
	var responseData []models.Data
	if len(result) == 0 {
		response.Data = responseData
		c.JSON(http.StatusOK, response)
		return
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
	c.JSON(http.StatusOK, response)
}

func (pg ServiceStruct) GetMinutesData(c *gin.Context) {
	minute_stringVal := c.Param("minute")

	minute, err := strconv.Atoi(minute_stringVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid minute parameter, must be an integer"})
		return
	}
	if minute <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid minute parameter, must be an integer greater than zero"})
		return
	}

	ctx := c.Request.Context()

	query := `
        SELECT timestamp, cpu_load, concurrency 
        FROM comcast 
        WHERE timestamp >= NOW() - ($1 * INTERVAL '1 minute')
        ORDER BY timestamp DESC
		`
	rows, err := pg.Db.Query(ctx, query, minute)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed for Fetcing past Minute Data : %v\n", err)
		os.Exit(1)
	}

	var result []db.Data
	result, err = pgx.CollectRows(rows, pgx.RowToStructByPos[db.Data])
	if err != nil {
		fmt.Errorf("GetMinuteData : Error in converting rows to struct")
	}
	var response models.GetDataResponseStruct
	var responseData []models.Data
	if len(result) == 0 {
		response.Data = responseData
		c.JSON(http.StatusOK, response)
		return
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
	c.JSON(http.StatusOK, response)
}

func (pg ServiceStruct) GetSecondsData(c *gin.Context) {
	second_stringVal := c.Param("second")

	second, err := strconv.Atoi(second_stringVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid second parameter, must be an integer"})
		return
	}
	if second <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid second parameter, must be an igreater than zero"})
		return
	}

	ctx := c.Request.Context()

	query := `
        SELECT timestamp, cpu_load, concurrency 
        FROM comcast 
        WHERE timestamp >= NOW() - ($1 * INTERVAL '1 second')
        ORDER BY timestamp DESC
		`
	rows, err := pg.Db.Query(ctx, query, second)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed for Fetcing past seconds Data : %v\n", err)
		os.Exit(1)
	}

	var result []db.Data
	result, err = pgx.CollectRows(rows, pgx.RowToStructByPos[db.Data])
	if err != nil {
		fmt.Errorf("GetSecondsData : Error in converting rows to struct")
	}
	var response models.GetDataResponseStruct
	var responseData []models.Data
	if len(result) == 0 {
		response.Data = responseData
		c.JSON(http.StatusOK, response)
		return
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
	c.JSON(http.StatusOK, response)
}

func (pg ServiceStruct) GetDataByDate(c *gin.Context) {
	date_stringVal := c.Param("date")
	const layout = "02-01-2006"
	parsedDate, err := time.Parse(layout, date_stringVal)
	fmt.Println(parsedDate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing date string, Please use DD-MM-YYYY format."})
		return
	}

	ctx := c.Request.Context()

	startOfDay := parsedDate.UTC()
	endOfDay := parsedDate.Add(24 * time.Hour).UTC()

	log.Printf("Querying for metrics between %v and %v", startOfDay, endOfDay)

	query := `
        SELECT timestamp, cpu_load, concurrency 
        FROM comcast 
        WHERE timestamp >= $1 AND timestamp < $2
        ORDER BY timestamp DESC
		`
	rows, err := pg.Db.Query(ctx, query, startOfDay, endOfDay)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed for Fetcing the given Date Data : %v\n", err)
		os.Exit(1)
	}

	var result []db.Data
	result, err = pgx.CollectRows(rows, pgx.RowToStructByPos[db.Data])
	if err != nil {
		fmt.Errorf("GetDataByDate : Error in converting rows to struct")
	}
	var response models.GetDataResponseStruct
	var responseData []models.Data
	if len(result) == 0 {
		response.Data = responseData
		c.JSON(http.StatusOK, response)
		return
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
	c.JSON(http.StatusOK, response)
}

func (pg ServiceStruct) GetSpecificDataSet(c *gin.Context) {
	var requestBody models.GetSpecificDataSetRequest

	err := c.BindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body format",
			"details": err.Error(),
		})
		return
	}
	if requestBody.StartTime.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required field: starttime"})
		return
	}
	if requestBody.EndTime.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required field: endtime"})
		return
	}
	if !requestBody.StartTime.Before(requestBody.EndTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time range: starttime must be before endtime"})
		return
	}

	startTime := requestBody.StartTime.UTC()
	endTime := requestBody.EndTime.UTC()

	ctx := c.Request.Context()

	if requestBody.OppCode != "" {
		res, err := utils.HandleAggregateQuery(pg.Db, c, ctx, requestBody, startTime, endTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Aggregate Query for GetSpecificDataSet FAILED",
				"details": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"operation": res.Operation,
			"parameter": res.Parameter,
			"value":     res.Value,
		})
	} else {
		res, err := utils.HandleRawDataQuery(pg.Db, c, ctx, startTime, endTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Aggregate Query for GetSpecificDataSet FAILED",
				"details": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}
