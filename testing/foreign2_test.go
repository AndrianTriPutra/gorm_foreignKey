package testing_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"foreignKey/pkg/database/model"
	"foreignKey/pkg/database/postgres"
	"foreignKey/pkg/logger"
	"testing"
	"time"

	"gorm.io/gorm"
)

func Test_F2_Createx1a(t *testing.T) {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable", "localhost", "5432",
		"postgres_test", "xxxxxxx", "yyyyyy")

	db, err := postgres.NewPostgres(dsn, true)
	if err != nil {
		logger.Level("panic", "Test", "failed connect to database:"+err.Error())
	}

	ctx := context.Background()
	errTx := db.WithTransaction(ctx, func(ctxWithTx context.Context, dbt *gorm.DB) error {
		dev := model.Devices{}
		log, err := dev.Find(dbt, "A0001")
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		loc, _ := time.LoadLocation("Asia/Jakarta")
		ts := time.Now().In(loc)
		data := model.Foreign2{
			Timestamp: ts,
			//Device_PID: &log.ID,
			Device_UID: log,
			Data:       1,
		}

		err = data.Create(dbt)
		if err != nil {
			return err
		}
		return nil
	})
	if errTx != nil {
		logger.Level("error", "test", "DoTransaction->"+errTx.Error())
	}
}

func Test_F2_Findx1(t *testing.T) {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable", "localhost", "5432",
		"postgres_test", "xxxxxxx", "yyyyyy")

	db, err := postgres.NewPostgres(dsn, true)
	if err != nil {
		logger.Level("panic", "Test", "failed connect to database:"+err.Error())
	}

	var data model.Foreign2
	ctx := context.Background()
	errTx := db.WithTransaction(ctx, func(ctxWithTx context.Context, dbt *gorm.DB) error {
		data, err = data.Find(dbt, 1)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return nil
	})
	if errTx != nil {
		logger.Level("error", "test", "DoTransaction->"+errTx.Error())
	}

	js, _ := json.MarshalIndent(data, " ", " ")
	logger.Trace("data:", string(js))
}

func Test_F2_Updatex1(t *testing.T) {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable", "localhost", "5432",
		"postgres_test", "xxxxxxx", "yyyyyy")

	db, err := postgres.NewPostgres(dsn, true)
	if err != nil {
		logger.Level("panic", "Test", "failed connect to database:"+err.Error())
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	ts := time.Now().In(loc)

	ctx := context.Background()
	errTx := db.WithTransaction(ctx, func(ctxWithTx context.Context, dbt *gorm.DB) error {
		dev := model.Devices{}
		dev_master, err := dev.Find(dbt, "A0002")
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		data := model.Foreign2{
			Timestamp: ts,
			//Device_PID: 1,
			Device_UID: dev_master,
			Data:       10,
		}
		log, err := data.Find(dbt, dev_master.ID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		data.ID = log.ID
		data.CreatedAt = log.CreatedAt
		data.UpdatedAt = ts
		data.DeletedAt = log.DeletedAt
		js, _ := json.MarshalIndent(data, " ", " ")
		logger.Trace("data:", string(js))

		err = data.Update(dbt)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return nil
	})
	if errTx != nil {
		logger.Level("error", "test", "DoTransaction->"+errTx.Error())
	}
}

func Test_F2_Deletedx2(t *testing.T) {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable", "localhost", "5432",
		"postgres_test", "xxxxxxx", "yyyyyy")

	db, err := postgres.NewPostgres(dsn, true)
	if err != nil {
		logger.Level("panic", "Test", "failed connect to database:"+err.Error())
	}

	dev_id := uint(2)
	ctx := context.Background()
	errTx := db.WithTransaction(ctx, func(ctxWithTx context.Context, dbt *gorm.DB) error {
		data := model.Foreign2{}
		log, err := data.Find(dbt, dev_id)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		//err = log.Delete(dbt)
		err = log.Remove(dbt)
		if err != nil {
			return err
		}
		return nil
	})
	if errTx != nil {
		logger.Level("error", "test", "DoTransaction->"+errTx.Error())
	}
}
