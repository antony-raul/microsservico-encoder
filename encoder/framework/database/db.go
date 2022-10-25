package database

import (
	"github.com/antony-raul/microsservico-encoder/domain"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstace := NewDb()

	dbInstace.Env = "Test"
	dbInstace.DbTypeTest = "sqlite3"
	dbInstace.DsnTest = ":memory:"
	dbInstace.AutoMigrateDb = true
	dbInstace.Debug = true

	connection, err := dbInstace.Connect()
	if err != nil {
		log.Fatalf("Test db error: %v", err)
	}

	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	if d.Env != "Test" {
		d.Db, err = gorm.Open(postgres.Open(d.Dsn))
	} else {
		d.Db, err = gorm.Open(sqlite.Open(d.Dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	}

	if err != nil {
		return nil, err
	}

	if d.Debug {
		d.Db.Debug()
	}

	if d.AutoMigrateDb {
		err = d.Db.AutoMigrate(&domain.Video{}, &domain.Job{})
	}
	if err != nil {
		return nil, err
	}

	return d.Db, nil
}
