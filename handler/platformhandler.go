package handler

import (
	"fmt"
	"strconv"

	"github.com/adityaslab/zopsmart-task/db"
	"github.com/adityaslab/zopsmart-task/models"
	"gofr.dev/pkg/gofr"
)

type stringResp struct {
	Message string
}

type platformResp struct {
	PlatformNumber int
	IsFree         bool
	TrainName      string
}

func TrainArrival(ctx *gofr.Context) (interface{}, error) {

	var p models.Platform
	if err := ctx.Bind(&p); err != nil {
		return nil, err
	}

	err := db.TrainArrival(ctx, p)

	if err != nil {
		return stringResp{Message: fmt.
			Sprintf("Train couldn't arrive at platform number %d!", p.PlatformNumber)}, err
	}

	return stringResp{Message: fmt.Sprintf("Train arrived at platform number %d", p.PlatformNumber)}, nil
}

func TrainDeparture(ctx *gofr.Context) (interface{}, error) {

	var p models.Platform
	if err := ctx.Bind(&p); err != nil {
		return nil, err
	}

	err := db.TrainDeparture(ctx, p)

	if err != nil {
		return stringResp{Message: "Task failed due to error"}, err
	}

	return stringResp{Message: "Train departed from the platform"}, nil
}

func GetAllPlatformDetails(ctx *gofr.Context) (interface{}, error) {

	platforms, err := db.GetAllPlatformDetails(ctx)
	resp := make([]platformResp, 0)
	for _, p := range platforms {
		var train models.Train
		checkIfPlatformFree := false
		if p.TrainNumber == 0 {
			checkIfPlatformFree = true
		}

		train, err = db.GetTrainByNumber(ctx, p.TrainNumber)
		var trainname string

		//we will hit this error only if train number is 0 meaning the platform is free
		if err != nil {
			trainname = ""
		} else {
			trainname = train.Name
		}

		r := platformResp{
			PlatformNumber: p.PlatformNumber,
			IsFree:         checkIfPlatformFree,
			TrainName:      trainname,
		}
		resp = append(resp, r)
	}

	return resp, nil
}

func GetPlatformDetailsByPlatformNo(ctx *gofr.Context, plaformNo int) (interface{}, error) {

	p, err := db.GetAllPlatformDetailsByPlatformNo(ctx, plaformNo)
	var train models.Train
	checkIfPlatformFree := false
	if p.TrainNumber == 0 {
		checkIfPlatformFree = true
	}

	train, err = db.GetTrainByNumber(ctx, p.TrainNumber)
	var trainname string

	//we will hit this error only if train number is 0 meaning the platform is free
	if err != nil {
		trainname = ""
	} else {
		trainname = train.Name
	}

	r := platformResp{
		PlatformNumber: p.PlatformNumber,
		IsFree:         checkIfPlatformFree,
		TrainName:      trainname,
	}

	return r, nil
}

func FindTrainOnStation(ctx *gofr.Context) (interface{}, error) {
	num := ctx.PathParam("n")
	trainNumber, err := strconv.Atoi(num)
	//not a number
	if err != nil {
		return nil, err
	}
	platform, err := db.FindTrainOnStation(ctx, trainNumber)
	if err != nil {
		return nil, err
	}
	var msg string
	if platform != 0 {
		msg = fmt.Sprintf("The train is on platform %d", platform)
	} else {
		msg = fmt.Sprintf("The train is not on the station")
	}
	return stringResp{Message: msg}, nil
}

func CreateNPlatforms(ctx *gofr.Context) (interface{}, error) {
	num := ctx.PathParam("n")
	n, err := strconv.Atoi(num)
	//not a number
	if err != nil {
		return nil, err
	}
	db.DeleteAllPlatforms(ctx)
	db.CreateNPlatforms(ctx, n)

	return stringResp{Message: "Platforms created successfully!"}, nil
}

func DeleteAllPlatforms(ctx *gofr.Context) (interface{}, error) {

	db.DeleteAllPlatforms(ctx)
	return nil, nil
}
