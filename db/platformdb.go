package db

import (
	"github.com/adityaslab/zopsmart-task/models"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

func GetAllPlatformStatus(ctx *gofr.Context) ([]models.Platform, error) {

	rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM platforms")

	if err != nil {
		return nil, errors.DB{Err: err}
	}

	defer rows.Close()

	platforms := make([]models.Platform, 0)

	for rows.Next() {
		var p models.Platform
		if err := rows.Scan(&p.PlatformNumber, &p.TrainNumber); err != nil {
			return nil, err
		}

		platforms = append(platforms, p)
	}

	return platforms, nil
}

func CreateNPlatforms(ctx *gofr.Context, n int) error {

	for i := 1; i <= n; i++ {
		_, err := ctx.DB().ExecContext(ctx, "INSERT INTO platforms (number, train) VALUES (?, ?)", i, 0)
		if err != nil {
			return errors.DB{Err: err}
		}
	}
	return nil
}

func DeleteAllPlatforms(ctx *gofr.Context) {

	ctx.DB().ExecContext(ctx, "DELETE FROM platforms")

}

func TrainArrival(ctx *gofr.Context, p models.Platform) error {

	//put a check if the platform exist and is empty
	//put a check if train is valid
	_, err := ctx.DB().ExecContext(ctx, "UPDATE platforms SET train = ? WHERE number = ?", p.TrainNumber, p.PlatformNumber)
	if err != nil {
		return errors.DB{Err: err}
	}
	UpdateTrainStatusByNumber(ctx, p.TrainNumber, "ARRIVED")

	return nil
}

func TrainDeparture(ctx *gofr.Context, p models.Platform) error {

	//put a check if train is valid and status is arrived
	//make it depart by replacing train no with 0 on platform table

	_, err := ctx.DB().ExecContext(ctx, "UPDATE  platforms set train = ? WHERE number = ?", 0, p.PlatformNumber)

	if err != nil {
		return errors.DB{Err: err}
	}
	UpdateTrainStatusByNumber(ctx, p.TrainNumber, "DEPARTED")
	return nil
}
