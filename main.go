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

	app.GET("/getTrainByNumber/{n}", handler.GetTrainByNumber)

	app.POST("/addNewTrain", handler.AddNewTrain)

	app.PUT("/updateTrainByNumber/{n}", handler.UpdateTrainByNumber)

	app.POST("/trainArrival", handler.TrainArrival)

	app.DELETE("/trainDeparture", handler.TrainDeparture)

	app.DELETE("/deleteTrainByNumber/{n}", handler.DeleteTrainByNumber)

	app.POST("/createNPlatforms/{n}", handler.CreateNPlatforms)

	app.DELETE("/deleteAllPlatforms", handler.DeleteAllPlatforms)

	app.Start()
}
