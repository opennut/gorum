package controllers

import (
	"github.com/opennut/gorum/app/models"
	"github.com/revel/revel"
)

type Discussions struct {
	PublicApp
}

func (c Discussions) Index(slug string) revel.Result {
	var currentDiscussion = &models.Discussion{}
	var locale = c.Request.Locale
	c.Txn.Where("slug = ?", slug).First(&currentDiscussion)
	c.Txn.Model(currentDiscussion).Related(&currentDiscussion.User)
	if c.Txn.Error != nil {
		return nil
	}

	var comments = []models.DiscussionComment{}
	c.Txn.Where("discussion_id = ?", currentDiscussion.ID).Find(&comments)
	for i, _ := range comments {
		c.Txn.Model(comments[i]).Related(&comments[i].User)
	}
	return c.Render(currentDiscussion, comments, locale)
}
