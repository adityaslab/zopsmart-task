package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/gofr"

	"github.com/adityaslab/zopsmart-task/models"
)

func TestCoreLayer(*testing.T) {
	app := gofr.New()

	seeder := datastore.NewSeeder(&app.DataStore, "../db")
	seeder.ResetCounter = true

	createTable(app)
}

func createTable(app *gofr.Gofr) {
	_, err := app.DB().Exec("DROP TABLE IF EXISTS trains;")

	if err != nil {
		return
	}
	_, err = app.DB().Exec("CREATE TABLE IF NOT EXISTS trains (" +
		"train_number INT NOT NULL, name varchar(100), status varchar(100), CONSTRAINT trains_PK PRIMARY KEY (train_number));")
	if err != nil {
		return
	}
}

func TestAddNewTrain(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		ctx.Logger.Error("mock connection failed")
	}

	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()

	tr := models.Train{
		TrainNumber: 122001,
		Name:        "CDG LKO EXP",
		Status:      "Arriving",
	}

	mock.ExpectExec("INSERT INTO trains (train_number, name, status) VALUES (?, ?, ?)").
		WithArgs(tr.TrainNumber, tr.Name, tr.Status).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Set up the expectations for the SELECT query
	rows := sqlmock.NewRows([]string{"train_number", "name", "status"}).
		AddRow(tr.TrainNumber, tr.Name, tr.Status)
	mock.ExpectQuery("SELECT train_number, name, status FROM trains WHERE train_number = ?").
		WithArgs(tr.TrainNumber).
		WillReturnRows(rows)

	resp, err := AddNewTrain(ctx, tr)
	assert.Error(t, err)
	fmt.Printf("%v and %v", resp.Name, resp.TrainNumber)
}

func TestFindTrainByNumberError(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		ctx.Logger.Error("an error '%s' was not expected when opening a stub database connection", err)
	}

	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()

	query := "SELECT train_number, name, status FROM trains WHERE train_number = \\?"
	tr := models.Train{
		TrainNumber: 122001,
		Name:        "CDG LKO EXP",
		Status:      "Arriving",
	}
	rows := sqlmock.NewRows([]string{"train_number", "name", "status"})

	mock.ExpectQuery(query).WithArgs(tr.TrainNumber).WillReturnRows(rows)

	tres, err := GetTrainByNumber(ctx, tr.TrainNumber)
	assert.Empty(t, tres)
	assert.Error(t, err)
}

func TestFindTrainByNumber(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		ctx.Logger.Error("an error '%s' was not expected when opening a stub database connection", err)
	}

	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()

	query := "SELECT train_number, name, status FROM trains WHERE train_number = \\?"
	tr := models.Train{
		TrainNumber: 122001,
		Name:        "CDG LKO EXP",
		Status:      "Arriving",
	}
	rows := sqlmock.NewRows([]string{"train_number", "name", "status"}).
		AddRow(tr.TrainNumber, tr.Name, tr.Status)

	mock.ExpectQuery(query).WithArgs(tr.TrainNumber).WillReturnRows(rows).WillReturnError(nil)

	tres, err := GetTrainByNumber(ctx, tr.TrainNumber)
	assert.NotNil(t, tres)
}
