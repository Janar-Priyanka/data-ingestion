package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"data-ingestion/cmd/server/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {

	//Dependency Injection
	type Env struct {
		service service.ServiceStruct
	}

	//DB Connection
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	err = conn.Ping(ctx)
	if err != nil {
		log.Fatalf("Fatal error: could not ping database: %v", err)
	}

	env := &Env{
		service: service.ServiceStruct{Db: conn},
	}
	router := gin.Default()

	router.POST("/getData", env.service.GetSpecificDataSet)
	router.GET("/days/:day", env.service.GetDayData)
	router.GET("/hours/:hour", env.service.GetHoursData)
	router.GET("/minutes/:minute", env.service.GetMinutesData)
	router.GET("/seconds/:second", env.service.GetSecondsData)
	router.GET("/date/:date", env.service.GetDataByDate)

	// Server Startup
	err = router.Run()
	if err != nil {
		fmt.Println("Error Starting Application")
		return
	}
}
