package main

import (
	"github.com/adityaslab/zopsmart-task/handler"
	"gofr.dev/pkg/gofr"
)

type Car struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	app := gofr.New()

	app.GET("/getAllTrains", handler.GetAllTrains)

	app.POST("/addNewTrain", handler.AddNewTrain)

	app.PUT("/updateTrainByNumber/{n}", handler.UpdateTrainByNumber)

	app.POST("/createNPlatforms/{n}", handler.CreateNPlatforms)

	app.Start()
}
