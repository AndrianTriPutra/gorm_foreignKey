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

func Test_Dev_Createx2(t *testing.T) {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable", "localhost", "5432",
		"postgres_test", "xxxxxxx", "yyyyyy")

	db, err := postgres.NewPostgres(dsn, true)
	if err != nil {
		logger.Level("panic", "Test", "failed connect to database:"+err.Error())
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	ts := time.Now().In(loc)
	data := model.Devices{
		// ID        uint      `gorm:"primaryKey;autoIncrement"`
		Timestamp: ts,
		Device_ID: "A0002",
	}

	ctx := context.Background()
	errTx := db.WithTransaction(ctx, func(ctxWithTx context.Context, dbt *gorm.DB) error {

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

func Test_Dev_Findx1B(t *testing.T) {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable", "localhost", "5432",
		"postgres_test", "xxxxxxx", "yyyyyy")

	db, err := postgres.NewPostgres(dsn, true)
	if err != nil {
		logger.Level("panic", "Test", "failed connect to database:"+err.Error())
	}

	var data model.Devices
	ctx := context.Background()
	errTx := db.WithTransaction(ctx, func(ctxWithTx context.Context, dbt *gorm.DB) error {
		data, err = data.Find(dbt, "A0002")
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

func Test_Dev_Updatex1B(t *testing.T) {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable", "localhost", "5432",
		"postgres_test", "xxxxxxx", "yyyyyy")

	db, err := postgres.NewPostgres(dsn, true)
	if err != nil {
		logger.Level("panic", "Test", "failed connect to database:"+err.Error())
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	ts := time.Now().In(loc)
	dev_exiting := "A0002"
	dev_new := "A0003"
	data := model.Devices{
		// ID        uint      `gorm:"primaryKey;autoIncrement"`
		Timestamp: ts,
		Device_ID: dev_new,
	}

	ctx := context.Background()
	errTx := db.WithTransaction(ctx, func(ctxWithTx context.Context, dbt *gorm.DB) error {
		log, err := data.Find(dbt, dev_exiting)
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

func Test_Dev_Deletedx1A(t *testing.T) {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable", "localhost", "5432",
		"postgres_test", "xxxxxxx", "yyyyyy")

	db, err := postgres.NewPostgres(dsn, true)
	if err != nil {
		logger.Level("panic", "Test", "failed connect to database:"+err.Error())
	}

	dev_id := "A0001"
	ctx := context.Background()
	errTx := db.WithTransaction(ctx, func(ctxWithTx context.Context, dbt *gorm.DB) error {
		data := model.Devices{}
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
