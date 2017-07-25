package controllers

import (
	"github.com/opennut/gorum/app/models"
	"github.com/revel/revel"
	"regexp"
)

// Home Controller anon user
type Home struct {
	PublicApp
}

// Index page
func (c Home) Index() revel.Result {
	var discussions = []models.Discussion{}
	c.Txn.Where("active = true").Find(&discussions).Limit(25)
	for i, _ := range discussions {
		c.Txn.Model(discussions[i]).Related(&discussions[i].User)
	}
	var tags = []models.Tag{}
	c.Txn.Where("active = true and parent_id is null").Find(&tags)
	return c.Render(discussions, tags)
}

// Index page
func (c Home) Tagged(tag string) revel.Result {
	var discussions = []models.Discussion{}
	c.Txn.LogMode(true)
	c.Txn.Joins(
		"JOIN discussion_tags ON discussion_tags.discussion_id = discussions.id").Joins(
		"JOIN tags ON  tags.id = discussion_tags.tag_id").Joins(
		"left outer join tags as tags_parents  on tags_parents.parent_id = tags.id").Where(
		"tags.slug = ? or tags.parent_id in (select id from tags where slug = ?)", tag, tag).Find(&discussions)
	for i, _ := range discussions {
		c.Txn.Model(discussions[i]).Related(&discussions[i].User)
	}
	var tags = []models.Tag{}
	c.Txn.Where("active = true and parent_id is null").Find(&tags)
	return c.Render(discussions, tags, tag)
}

// Detail page
func (c Home) Detail(ID int, name string) revel.Result {
	c.Validation.Required(name).Key("name").Message("El nombre es requerido")
	c.Validation.Required(name)
	c.Validation.MaxSize(name, 15)
	c.Validation.MinSize(name, 4)
	c.Validation.Match(name, regexp.MustCompile("^\\w*$"))
	c.Validation.Keep()
	c.Flash.Success("Email changed")
	return c.Render()
}
