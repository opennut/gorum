package models

// Discussion model
type Discussion struct {
	Model
	Title  string `gorm:"size:255"`
	Body   string `gorm:"TEXT"`
	Tags   []Tag  `gorm:"many2many:discussion_tags;"`
	Active bool
}

// DiscussionPlus model
type DiscussionPlus struct {
	Model
	Discussion   Discussion `gorm:"ForeignKey:DiscussionID"`
	DiscussionID uint       `gorm:"not null;"`
	User         User       `gorm:"ForeignKey:UserID"`
	UserID       uint       `gorm:"not null;"`
}
