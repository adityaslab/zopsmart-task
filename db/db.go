package db

import (
	"github.com/adityaslab/zopsmart-task/models"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

func GetAllTrains(ctx *gofr.Context) ([]models.Train, error) {

	rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM trains")

	if err != nil {
		return nil, errors.DB{Err: err}
	}

	defer rows.Close()

	trains := make([]models.Train, 0)

	for rows.Next() {
		var t models.Train
		if err := rows.Scan(&t.Number, &t.Name, &t.Status); err != nil {
			return nil, errors.DB{Err: err}

		}

		trains = append(trains, t)
	}

	return trains, nil
}

func GetTrainByNumber(ctx *gofr.Context, trainNo int) (models.Train, error) {
	var res models.Train
	err := ctx.DB().QueryRowContext(ctx,
		"SELECT * FROM trains WHERE number = ?", trainNo).Scan(&res.Number, &res.Name, &res.Status)

	if err != nil {
		return models.Train{}, errors.DB{Err: err}
	}
	return res, err
}

func AddNewTrain(ctx *gofr.Context, t models.Train) (models.Train, error) {
	//check if number already exists
	_, err := ctx.DB().ExecContext(ctx,
		"INSERT INTO trains (number, name, status) VALUES (?, ?, ?)", t.Number, t.Name, t.Status)

	if err != nil {
		return models.Train{}, errors.DB{Err: err}

	}
	var resp models.Train
	resp, e := GetTrainByNumber(ctx, t.Number)
	if e != nil {
		return models.Train{}, errors.DB{Err: err}
	}
	return resp, nil
}

func UpdateTrainByNumber(ctx *gofr.Context, n int, t models.Train) (models.Train, error) {
	//check if this train number(n) exist in db or throw error
	_, e := GetTrainByNumber(ctx, n)
	if e != nil {
		return models.Train{}, errors.InvalidParam{Param: []string{"train number"}}
	}

	_, err := ctx.DB().ExecContext(ctx,
		"UPDATE trains SET name = ?, status = ? WHERE number = ?", t.Name, t.Status, n)

	if err != nil {
		return models.Train{}, errors.DB{Err: err}

	}

	var resp models.Train
	resp, error := GetTrainByNumber(ctx, n)
	if error != nil {
		return models.Train{}, errors.DB{Err: error}
	}
	return resp, nil
}

func DeleteTrainByNumber(ctx *gofr.Context, trainNumber int) (interface{}, error) {
	_, error := ctx.DB().ExecContext(ctx, "DELETE FROM trains WHERE number = ?", trainNumber)
	return nil, error
}

// Updates train status when arrival or departed functions are called
// This is a helper function to be used internally and not exposed to the API
func UpdateTrainStatusByNumber(ctx *gofr.Context, trainno int, status string) {
	ctx.DB().ExecContext(ctx,
		"UPDATE trains SET status = ? WHERE number = ?", status, trainno)
}
