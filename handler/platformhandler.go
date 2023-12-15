package handler

import (
	"strconv"

	"gofr.dev/pkg/gofr"
)

func CreateNPlatforms(ctx *gofr.Context) (interface{}, error) {
	num := ctx.PathParam("n")
	n, err := strconv.Atoi(num)
	//not a number
	if err != nil {
		return nil, err
	}

	for i := 1; i <= n; i++ {
		_, err := ctx.DB().ExecContext(ctx, "INSERT INTO platforms (number, train) VALUES (?, ?)", i, 0)
		if err != nil {
			return nil, err
		}
	}
	return nil, err
}
