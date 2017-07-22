package controllers

import (
	"fmt"
	"github.com/opennut/gorum/app/models"
	"github.com/revel/revel"
)

type Discussions struct {
	PublicApp
}

func (c Discussions) Index(slug string) revel.Result {
	var currentDiscussion = &models.Discussion{}
	fmt.Println("c.Request")
	fmt.Println(c.Request.Locale)
	var locale = c.Request.Locale
	c.Request.Locale = "es"
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
