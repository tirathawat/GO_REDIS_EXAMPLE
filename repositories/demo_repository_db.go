package repositories

import (
	"example/database"
	"example/models/storages"
)

type demoRepository struct {
	Database database.IDatabase
}

func NewDemoRepository(db database.IDatabase) demoRepository {
	db.GetDB().AutoMigrate(storages.DemoDB{})
	db.FeedData()
	return demoRepository{Database: db}
}

func (r demoRepository) GetDemo() (demo []storages.DemoDB, err error) {
	err = r.Database.GetDB().
		Order("priority desc").
		Limit(30).
		Find(&demo).
		Error
	return demo, err
}
