package helper

import (
	"api/modules/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func OpenDB(conn, schm, ver string) *gorm.DB {
	dsn := conn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: schm + ".",
			SingularTable: true,
		},
	})

	if err  != nil {
		panic("failed connect to database " + err.Error())
	}

	db.Exec("CREATE SCHEMA IF NOT EXISTS " + schm)

	err = db.AutoMigrate(
		&model.AttendanceRecord{},
		&model.AttendanceSession{},
		&model.DeviceLog{},
		&model.Role{},
		&model.Room{},
		&model.ScheduleClass{},
		&model.School{},
		&model.Student{},
		&model.Subject{},
		&model.Teacher{},
		&model.User{},
	)

	if err  != nil {
		panic("failed to migrate database")
	}

	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to closed database")
	}
	sqlDB.Close()
}