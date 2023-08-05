package data

import (
	"context"
	"fmt"
	"github.com/goliajp/gormx"
	"github.com/goliajp/http-api-gin/core"
	"github.com/goliajp/http-api-gin/data/rdm"
	"github.com/goliajp/http-api-gin/env"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Pg *gormx.Pg

var PgTables = []interface{}{
	new(rdm.Foo),
	new(rdm.Bar),
	new(rdm.User),
}

func GetPg(ctx context.Context) *gorm.DB {
	if ctx == nil {
		ctx = context.Background()
	}
	return Pg.DB().WithContext(ctx)
}

func AutoMigrate() {
	mdb := Pg.Open("postgres")
	if err := gormx.CreateDatabase(mdb, env.RdName); err != nil {
		log.Fatalf("create database %s failed: %v", env.RdName, err)
	}

	db := GetPg(nil)
	if err := db.AutoMigrate(PgTables...); err != nil {
		log.Fatalf("auto migrate tables failed: %v", err)
	}
}

func PgPrepare() {
	if env.RdRebuild {
		PgClear()
	}
	AutoMigrate()
	if env.RdRebuild && env.RdMockData {
		PgMockData()
	}
}

func PgClear() {
	mdb := Pg.Open("postgres")
	if err := gormx.DropDatabase(mdb, env.RdName); err != nil {
		log.Fatalf("drop database failed: %v", env.RdName, err)
	}
	fmt.Printf("rd: cleared\n")
}

func PgMockData() {
	rd := GetPg(nil)

	// user
	users := []core.User{
		{
			Name:     "mock-user-1",
			Password: "mock-password-1",
			Email:    "mock-1@email.com",
			Gender:   rdm.UserGenderMale,
		},
		{
			Name:     "mock-user-2",
			Password: "mock-password-2",
			Email:    "mock-2@email.com",
			Gender:   rdm.UserGenderMale,
		},
		{
			Name:     "mock-user-3",
			Password: "mock-password-3",
			Email:    "mock-3@email.com",
			Gender:   rdm.UserGenderMale,
		},
	}
	rd.Create(&users)

	// foo
	foos := []core.Foo{
		{
			Count:       1,
			Description: "mock-foo-desc-1",
			UserId:      users[0].Id,
		},
		{
			Count:       2,
			Description: "mock-foo-desc-2",
			UserId:      users[0].Id,
		},
		{
			Count:       11,
			Description: "mock-foo-desc-3",
			UserId:      users[1].Id,
		},
		{
			Count:       22,
			Description: "mock-foo-desc-4",
			UserId:      users[1].Id,
		},
		{
			Count:       111,
			Description: "mock-foo-desc-5",
			UserId:      users[2].Id,
		},
		{
			Count:       222,
			Description: "mock-foo-desc-6",
			UserId:      users[2].Id,
		},
	}
	rd.Create(&foos)

	// bar
	bars := []core.Bar{
		{
			Name:        "mock-bar-1",
			Description: "mock-bar-desc-1",
			UserId:      users[0].Id,
		},
		{
			Name:        "mock-bar-2",
			Description: "mock-bar-desc-2",
			UserId:      users[0].Id,
		},
		{
			Name:        "mock-bar-3",
			Description: "mock-bar-desc-3",
			UserId:      users[1].Id,
		},
		{
			Name:        "mock-bar-4",
			Description: "mock-bar-desc-4",
			UserId:      users[1].Id,
		},
		{
			Name:        "mock-bar-5",
			Description: "mock-bar-desc-5",
			UserId:      users[2].Id,
		},
		{
			Name:        "mock-bar-6",
			Description: "mock-bar-desc-6",
			UserId:      users[2].Id,
		},
	}
	rd.Create(&bars)
}

func init() {
	Pg = gormx.NewPg(
		&gormx.PgConfig{
			User:     env.RdUser,
			Password: env.RdPassword,
			Host:     env.RdHost,
			Port:     env.RdPort,
			Dbname:   env.RdName,
			Tz:       env.RdTz,
		},
	)
}
