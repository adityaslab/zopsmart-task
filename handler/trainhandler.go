package handler

import (
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
