package postgres

import (
	"context"
	"foreignKey/pkg/database/model"
	"foreignKey/pkg/logger"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	ormlog "gorm.io/gorm/logger"
)

type postgresDb struct {
	db *gorm.DB
}

type DatabaseI interface {
	Db(ctx context.Context) interface{}
	WithTransaction(ctx context.Context, fn func(ctxWithTx context.Context, dbt *gorm.DB) error) error
}

func NewPostgres(dsn string, migrate bool) (DatabaseI, error) {
	newLogger := ormlog.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		ormlog.Config{
			SlowThreshold:             5 * time.Second, // Slow SQL threshold
			Colorful:                  true,            // Disable color
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			LogLevel:                  ormlog.Error,    // Log level
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return postgresDb{db: db}, err
	}

	if migrate {
		db.AutoMigrate(
			model.Devices{},

			model.Foreign1{},
			model.Foreign2{},
			model.Foreign3{},
		)
	}

	logger.Level("info", "NewPostgres", "successed connected to database")

	return postgresDb{db: db}, nil
}

func (d postgresDb) Db(ctx context.Context) interface{} {
	tx := ctx.Value("txContext")
	if tx == nil {
		return d.db
	}
	return tx.(*gorm.DB)
}

func (d postgresDb) WithTransaction(ctx context.Context, fn func(ctxWithTx context.Context, dbt *gorm.DB) error) error {
	tx := d.db.Begin()
	ctxWithTx := context.WithValue(ctx, tx, "txContext")
	errFn := fn(ctxWithTx, tx)
	if errFn != nil {
		//logger.Level("debug", "WithTransaction", "Rollback")
		errRlbck := tx.Rollback().Error
		if errRlbck != nil {
			logger.Level("error", "WithTransaction", "failed on rollback transaction:"+errRlbck.Error())
		}
		return errFn
	}

	//logger.Level("debug", "WithTransaction", "Commit")
	errCmmt := tx.Commit().Error
	if errCmmt != nil {
		logger.Level("error", "WithTransaction", "failed on commit transaction:"+errCmmt.Error())
	}
	return errFn
}
