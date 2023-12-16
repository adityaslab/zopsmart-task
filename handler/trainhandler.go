package handler

import (
	"strconv"

	"github.com/adityaslab/zopsmart-task/db"
	"github.com/adityaslab/zopsmart-task/models"
	"gofr.dev/pkg/gofr"
)

type train struct{}

func GetAllTrains(ctx *gofr.Context) (interface{}, error) {
	trains, err := db.GetAllTrains(ctx)
	if err != nil {
		return nil, err
	}
	return trains, nil
}

func GetTrainByNumber(ctx *gofr.Context) (interface{}, error) {
	num := ctx.PathParam("n")

	trainNumber, err := strconv.Atoi(num)
	//not a number
	if err != nil {
		return nil, err
	}

	train, err := db.GetTrainByNumber(ctx, trainNumber)
	if err != nil {
		return nil, err
	}
	return train, nil
}

func AddNewTrain(ctx *gofr.Context) (interface{}, error) {
	var t models.Train
	if err := ctx.BindStrict(&t); err != nil {
		return nil, err
	}

	response, err := db.AddNewTrain(ctx, t)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func UpdateTrainByNumber(ctx *gofr.Context) (interface{}, error) {
	num := ctx.PathParam("n")

	trainNumber, err := strconv.Atoi(num)
	//not a number
	if err != nil {
		return nil, err
	}
	var t models.Train
	if err := ctx.Bind(&t); err != nil {
		return nil, err
	}

	response, error := db.UpdateTrainByNumber(ctx, trainNumber, t)
	if error != nil {
		return nil, error
	}
	return response, nil
}

func DeleteTrainByNumber(ctx *gofr.Context) (interface{}, error) {
	tn := ctx.PathParam("n")

	trainNumber, err := strconv.Atoi(tn)
	//not a number
	if err != nil {
		return nil, err
	}
	db.DeleteTrainByNumber(ctx, trainNumber)
	return nil, err
}
