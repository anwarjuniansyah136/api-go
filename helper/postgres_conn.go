package helper

import (
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
		panic("failed connect to database")
	}

	db.Exec("CREATE SCHEMA IF NOT EXISTS " + schm)

	err = db.AutoMigrate()

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