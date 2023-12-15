package handler

import (
	"strconv"

	"github.com/adityaslab/zopsmart-task/models"
	"gofr.dev/pkg/gofr"
)

func TrainArrival(ctx *gofr.Context) (interface{}, error) {

	var p models.Platform
	if err := ctx.Bind(&p); err != nil {
		return nil, err
	}
	//put a check if the platform exist and is empty
	//put a check if train is valid
	_, err := ctx.DB().ExecContext(ctx, "UPDATE platforms SET train = ? WHERE number = ?", p.Train, p.Number)
	if err != nil {
		return nil, err
	}
	UpdateTrainStatusByNumber(ctx, p.Train, "ARRIVED")

	return nil, err
}

func GetAllPlatformStatus(ctx *gofr.Context) (interface{}, error) {

	rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM platforms")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	platforms := make([]models.Platform, 0)

	for rows.Next() {
		var p models.Platform
		if err := rows.Scan(&p.Number, &p.Train); err != nil {
			return nil, err
		}

		platforms = append(platforms, p)
	}

	return platforms, nil
}

// func GetPlatformStatus(ctx *gofr.Context, platformNo int) (interface{}, error) {
// }

func TrainDeparture(ctx *gofr.Context) (interface{}, error) {

	var p models.Platform
	if err := ctx.Bind(&p); err != nil {
		return nil, err
	}
	//put a check if train is valid and status is arrived
	//make it depart by replacing train no with 0 on platform table

	_, err := ctx.DB().ExecContext(ctx, "UPDATE  platforms set train = ? WHERE number = ?", 0, p.Number)

	if err != nil {
		return nil, err
	}
	UpdateTrainStatusByNumber(ctx, p.Train, "DEPARTED")

	return nil, err
}

func CreateNPlatforms(ctx *gofr.Context) (interface{}, error) {
	num := ctx.PathParam("n")
	n, err := strconv.Atoi(num)
	//not a number
	if err != nil {
		return nil, err
	}
	DeleteAllPlatforms(ctx)
	for i := 1; i <= n; i++ {
		_, err := ctx.DB().ExecContext(ctx, "INSERT INTO platforms (number, train) VALUES (?, ?)", i, 0)
		if err != nil {
			return nil, err
		}
	}
	return nil, err
}

func DeleteAllPlatforms(ctx *gofr.Context) (interface{}, error) {

	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM platforms")
	return nil, err

}
