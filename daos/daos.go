package daos

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/trongdth/mroom_backend/config"
	"github.com/trongdth/mroom_backend/database"
	"github.com/trongdth/mroom_backend/models"
)

var (
	tables = []interface{}{(*models.User)(nil), (*models.UserSubscribe)(nil)}
	db     *gorm.DB
)

func AutoMigrate() error {
	if db == nil {
		return errors.Wrap(errors.New("Init db instance first!"), "db.AutoMigrate")
	}

	allTables := make([]interface{}, 0)
	allTables = append(allTables, tables...)

	if err := db.AutoMigrate(allTables...).Error; err != nil {
		return errors.Wrap(err, "db.AutoMigrate")
	}

	if err := AddForeignKeys(); err != nil {
		return errors.Wrap(err, "db.AutoMigrate")
	}

	return nil
}

func AddForeignKeys() error {
	if db == nil {
		return errors.Wrap(errors.New("Init db instance first!"), "db.AutoMigrate")
	}

	return nil
}

func Init(conf *config.Config) error {
	var err error
	db, err = database.Init(conf)
	if err != nil {
		return errors.Wrap(err, "database.Init")
	}
	return nil
}

// GetDB : getter
func GetDB() *gorm.DB {
	return db
}

func WithTransaction(callback func(*gorm.DB) error) error {
	tx := db.Begin()

	if err := callback(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return errors.Wrap(err, "tx.Commit()")
	}

	return nil
}

func WithDB(callback func(*gorm.DB) error) error {
	if err := callback(db); err != nil {
		//return errors.Wrap(err, "callback(tx)")
		return err
	}

	return nil
}
