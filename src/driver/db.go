package driver

import (
	"fmt"
	"log"
	"record_song_key_service/src/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseDriver struct {
	DB *gorm.DB
}

var dbInstance *DatabaseDriver

func GetInstance() *DatabaseDriver {
	if dbInstance == nil {
		dbInstance = &DatabaseDriver{}
	}

	return dbInstance
}

func (this DatabaseDriver) GetConnection(host, user, password, dbname, port string) (db *gorm.DB) {
	if this.DB == nil {

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)

		var err error
		this.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panicln(err)
			panic("failed to connect database")
		}

		err = this.DB.AutoMigrate(
			&model.Key{},
			&model.Note{},
			&model.Singer{},
			&model.Song{},
			&model.CoverSongKey{},
			&model.User{},
			&model.UserSong{},
		)

		if err != nil {
			log.Panicln(err)
			panic("failed to migrate database")
		}
	}

	return this.DB
}
