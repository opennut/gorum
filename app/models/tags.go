package models

// Tag model
type Tag struct {
	Model
	Title    string `gorm:"size:255"`
	Slug     string `gorm:"type:varchar(75);unique_index"`
	Color    string `gorm:"size:7"`
	Parent   User   `gorm:"ForeignKey:ParentID"`
	ParentID uint   `gorm:"null;"`
	Active   bool
}
