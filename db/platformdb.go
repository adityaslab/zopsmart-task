package db

import (
	"fmt"
	"strings"

	"github.com/adityaslab/zopsmart-task/models"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

func GetAllPlatformDetails(ctx *gofr.Context) ([]models.Platform, error) {

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

func GetPlatformDetailsByPlatformNo(ctx *gofr.Context, plaformNo int) (models.Platform, error) {
	if !validatePlatformNumber(ctx, plaformNo) {
		return models.Platform{}, &errors.Response{Reason: "Invalid platform number"}
	}
	var res models.Platform
	err := ctx.DB().QueryRowContext(ctx,
		"SELECT * FROM platforms WHERE platform_number = ?", plaformNo).Scan(&res.PlatformNumber, &res.TrainNumber)

	if err != nil {
		return models.Platform{}, errors.DB{Err: err}
	}
	return res, err
}

func CreateNPlatforms(ctx *gofr.Context, n int) error {

	for i := 1; i <= n; i++ {
		_, err := ctx.DB().ExecContext(ctx, "INSERT INTO platforms (platform_number, train_number) VALUES (?, ?)", i, -1)
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

	//check if the platform exist
	if !validatePlatformNumber(ctx, p.PlatformNumber) {
		return &errors.Response{Reason: "Invalid platform number"}
	}

	//put a check if train is valid
	if !validateTrainNumberInTrainDb(ctx, p.TrainNumber) {
		return &errors.Response{Reason: "Invalid train number"}
	}

	//check if platform is empty
	platform, err := GetPlatformDetailsByPlatformNo(ctx, p.PlatformNumber)
	if platform.TrainNumber != -1 {
		return &errors.Response{Reason: "There is already another train on the platform"}
	}

	//check if the train's status is valid
	train, err := GetTrainByNumber(ctx, p.TrainNumber)
	if strings.ToLower(train.Status) != "arriving" {
		message := fmt.Sprintf("Trains current status: %v", train.Status)
		return &errors.Response{Reason: message}
	}

	_, err = ctx.DB().ExecContext(ctx, "UPDATE platforms SET train_number = ? WHERE platform_number = ?", p.TrainNumber, p.PlatformNumber)
	if err != nil {
		return errors.DB{Err: err}
	}
	UpdateTrainStatusByNumber(ctx, p.TrainNumber, "ARRIVED")

	return nil
}

func TrainDeparture(ctx *gofr.Context, p models.Platform) error {

	//check if the platform exist
	if !validatePlatformNumber(ctx, p.PlatformNumber) {
		return &errors.Response{Reason: "Invalid platform number"}
	}

	//put a check if train is valid
	if !validateTrainNumberInTrainDb(ctx, p.TrainNumber) {
		return &errors.Response{Reason: "Invalid train number"}
	}

	//check if this train is on the platform given in the request body
	platform, err := GetPlatformDetailsByPlatformNo(ctx, p.PlatformNumber)
	if platform.TrainNumber != p.TrainNumber {
		msg := fmt.Sprintf("Train no %d is not on platform no %d", p.TrainNumber, p.PlatformNumber)
		return &errors.Response{Reason: msg}
	}

	//check if the train's status is valid
	// train, err := GetTrainByNumber(ctx, p.TrainNumber)
	// if strings.ToLower(train.Status) != "arrived" {
	// 	message := fmt.Sprintf("Trains current status: %v", train.Status)
	// 	return &errors.Response{Reason: message}
	// }
	_, err = ctx.DB().ExecContext(ctx, "UPDATE  platforms set train_number = ? WHERE platform_number = ?", -1, p.PlatformNumber)

	if err != nil {
		return errors.DB{Err: err}
	}
	UpdateTrainStatusByNumber(ctx, p.TrainNumber, "DEPARTED")
	return nil
}

func FindTrainOnStation(ctx *gofr.Context, trainNumber int) (int, error) {
	validationFlag := validateTrainNumberInTrainDb(ctx, trainNumber)
	if !validationFlag {
		return -1, &errors.Response{Reason: "Invalid train number"}
	}
	var res models.Platform
	err := ctx.DB().QueryRowContext(ctx,
		"SELECT * FROM platforms WHERE train_number = ?", trainNumber).Scan(&res.PlatformNumber, &res.TrainNumber)

	if err != nil {
		if validationFlag {
			return -1, nil
		} else {
			return -1, errors.DB{Err: err}
		}
	}
	return res.PlatformNumber, nil
}

// returns true if the platform number already exists
func validatePlatformNumber(ctx *gofr.Context, platformNo int) bool {

	var res models.Platform
	ctx.DB().QueryRowContext(ctx,
		"SELECT * FROM platforms WHERE platform_number = ?", platformNo).Scan(&res.PlatformNumber, &res.TrainNumber)

	if res.PlatformNumber == platformNo {
		return true
	} else {
		return false
	}
}
