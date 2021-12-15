package storages

type DemoDB struct {
	ID       int    `gorm:"primaryKey;column:id" json:"id"`
	Name     string `gorm:"primaryKey;column:name" json:"name"`
	Priority int    `gorm:"primaryKey;column:priority" json:"priority"`
}

func (d DemoDB) TableName() string {
	return "Demo"
}
