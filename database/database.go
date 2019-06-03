package database

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/trongdth/gin_template/config"
)

// Init : config
func Init(config *config.Config) (*gorm.DB, error) {
	databaseConn, err := gorm.Open("mysql", config.Db)
	if config.Environment != "production" {
		databaseConn.LogMode(true)
	}
	if err != nil {
		return nil, errors.Wrap(err, "gorm.Open")
	}
	//databaseConn = databaseConn.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8_unicode_ci auto_increment=1")
	// skip save associations of gorm -> manual save by code
	databaseConn = databaseConn.Set("gorm:save_associations", false)
	databaseConn = databaseConn.Set("gorm:association_save_reference", true)
	databaseConn.DB().SetMaxOpenConns(20)
	databaseConn.DB().SetMaxIdleConns(10)
	return databaseConn, err
}
