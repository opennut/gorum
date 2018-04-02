package models

// Discussion model
type Discussion struct {
	Model
	Slug   string `gorm:"type:varchar(75);unique_index"`
	Title  string `gorm:"size:255"`
	Body   string `gorm:"TEXT"`
	Tags   []Tag  `gorm:"many2many:discussion_tags;"`
	User   User   `gorm:"ForeignKey:UserID"`
	UserID uint   `gorm:"not null;"`
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

// DiscussionComment model
type DiscussionComment struct {
	Model
	Discussion   Discussion `gorm:"ForeignKey:DiscussionID"`
	DiscussionID uint       `gorm:"not null;"`
	User         User       `gorm:"ForeignKey:UserID"`
	UserID       uint       `gorm:"not null;"`
	Body         string     `gorm:"TEXT"`
}
