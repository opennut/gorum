package controllers

import (
	"github.com/opennut/gorum/app/models"
	"github.com/opennut/gorum/app/routes"
	"github.com/revel/revel"
	//gormc "github.com/revel/modules/orm/gorm/app/controllers"
	gormc "github.com/revel/modules/orm/gorm/app/controllers"
)

type FileInfo struct {
	ContentType string
	Filename    string
	RealFormat  string `json:",omitempty"`
	Resolution  string `json:",omitempty"`
	Size        int
	Status      string `json:",omitempty"`
}

// PublicApp without permissions
type PublicApp struct {
	*revel.Controller
	gormc.TxnController
}

// AddUser func
// This function add user variable to ViewArgs
func (c PublicApp) AddUser() revel.Result {
	if user := c.connected(); user != nil {
		c.ViewArgs["user"] = user
	}
	return nil
}

// connected return connected user
func (c PublicApp) connected() *models.User {
	if c.ViewArgs["user"] != nil {
		return c.ViewArgs["user"].(*models.User)
	}
	if username, ok := c.Session["user"]; ok {
		return c.currentUser(username)
	}
	return nil
}

// Ger current user by email
func (c PublicApp) currentUser(email string) *models.User {
	var currentUser = &models.User{}
	c.Txn.Where("email = ?", email).First(&currentUser)
	if c.Txn.Error != nil {
		return nil
	}
	c.ViewArgs["currentController"] = c.Name
	c.ViewArgs["currentAction"] = c.MethodName
	return currentUser
}

// App user loggued
type App struct {
	PublicApp
}

// Check if user is connected
func (c App) checkUser() revel.Result {
	if user := c.connected(); user == nil {
		c.Validation.Required(user != nil).Key("Email").Message("Permissions required")
		return c.Redirect(routes.Home.Index())
	}
	return nil
}

// AdminApp user loggued
type AdminApp struct {
	PublicApp
}

// Check if user is connected
func (c AdminApp) checkUser() revel.Result {
	if user := c.connected(); user == nil {
		c.Validation.Required(user != nil).Key("Email").Message("Permissions required")
		return c.Redirect(routes.Home.Index())
	} else {
		/*if !user.IsAdmin {
			c.Validation.Required(user != nil).Key("Email").Message("Permissions required")
			return c.Redirect(routes.Home.Index())
		}*/
	}
	return nil
}
