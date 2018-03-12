package controllers

import (
	"fmt"
	"github.com/opennut/gorum/app/models"
	gorm "github.com/revel/modules/orm/gorm/app"

	"github.com/revel/revel"
	"github.com/russross/blackfriday"
	"html/template"
)

func InitializeDB() {
	gorm.InitDB()
	gorm.DB.AutoMigrate(&models.User{})
	gorm.DB.AutoMigrate(&models.Tag{})
	gorm.DB.AutoMigrate(&models.Discussion{})
	gorm.DB.AutoMigrate(&models.DiscussionComment{})
	// var firstUser = models.User{Name: "Demo", Email: "demo@demo.com"}
	// firstUser.SetNewPassword("demo")
	// firstUser.Active = true
	// gorm.DB.Create(&firstUser)
}

func init() {
	revel.OnAppStart(InitializeDB)
	revel.InterceptMethod(PublicApp.AddUser, revel.BEFORE)
	revel.InterceptMethod(App.checkUser, revel.BEFORE)
	revel.InterceptMethod(AdminApp.checkUser, revel.BEFORE)

	revel.TemplateFuncs["markdown"] = func(str interface{}) template.HTML {
		s := blackfriday.MarkdownCommon([]byte(fmt.Sprintf("%s", str)))
		return template.HTML(s)
	}

	revel.TemplateFuncs["subtags"] = func(tag uint) []models.Tag {
		var tags = []models.Tag{}
		gorm.DB.Where("active = true and parent_id = ?", tag).Find(&tags)
		return tags
	}

}
