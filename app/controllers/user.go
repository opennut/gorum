package controllers

import (
	"github.com/opennut/gorum/app/models"
	//"github.com/opennut/gorum/app/routes"
	"github.com/revel/revel"
)

// Account Controller
type Users struct {
	AdminApp
}

func (c Users) Profile() revel.Result {
	email := c.Session["user"]
	var user models.User
	c.Txn.Where("email = ?", email).First(&user)
	if c.Txn.Error != nil {
		panic(c.Txn.Error)
	}

	return c.Render(user)
}
