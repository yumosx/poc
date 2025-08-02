package dao

import "gorm.io/gorm"

func InitTables(db *gorm.DB) error {
	db.Migrator().DropTable(&Task{})
	return db.AutoMigrate(
		&Task{},
		&Chat{},
	)
}
