package handler

import (
	"strconv"

	"github.com/adityaslab/zopsmart-task/models"
	"gofr.dev/pkg/gofr"
)

type train struct{}

func GetAllTrains(ctx *gofr.Context) (interface{}, error) {

	rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM trains")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	trains := make([]models.Train, 0)

	for rows.Next() {
		var t models.Train
		if err := rows.Scan(&t.Number, &t.Name, &t.Status); err != nil {
			return nil, err
		}

		trains = append(trains, t)
	}

	return trains, nil
}

func GetTrainByNumber(ctx *gofr.Context) (interface{}, error) {
	num := ctx.PathParam("n")

	n, err := strconv.Atoi(num)
	//not a number
	if err != nil {
		return nil, err
	}

	rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM trains WHERE number = ?", n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var t models.Train
	for rows.Next() {
		if err := rows.Scan(&t.Number, &t.Name, &t.Status); err != nil {
			return nil, err
		}
	}
	return t, err
}

func AddNewTrain(ctx *gofr.Context) (interface{}, error) {
	var t models.Train
	if err := ctx.Bind(&t); err != nil {
		return nil, err
	}

	res, err := ctx.DB().ExecContext(ctx,
		"INSERT INTO trains (number, name, status) VALUES (?, ?, ?)", t.Number, t.Name, t.Status)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func UpdateTrainByNumber(ctx *gofr.Context) (interface{}, error) {
	num := ctx.PathParam("n")

	n, err := strconv.Atoi(num)
	//not a number
	if err != nil {
		return nil, err
	}
	//check if this n exist in db or throw err
	var t models.Train
	if err := ctx.Bind(&t); err != nil {
		return nil, err
	}

	res, err := ctx.DB().ExecContext(ctx,
		"UPDATE trains SET name = ?, status = ? WHERE number = ?", t.Name, t.Status, n)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func UpdateTrainStatusByNumber(ctx *gofr.Context, trainno int, status string) (interface{}, error) {
	res, err := ctx.DB().ExecContext(ctx,
		"UPDATE trains SET status = ? WHERE number = ?", status, trainno)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func DeleteTrainByNumber(ctx *gofr.Context) (interface{}, error) {
	num := ctx.PathParam("n")

	n, err := strconv.Atoi(num)
	//not a number
	if err != nil {
		return nil, err
	}
	_, error := ctx.DB().ExecContext(ctx, "DELETE FROM trains WHERE number = ?", n)
	return nil, error
}
