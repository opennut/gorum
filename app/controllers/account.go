package controllers

import (
	"github.com/opennut/gorum/app/models"
	"github.com/opennut/gorum/app/routes"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

// Account Controller
type Accounts struct {
	PublicApp
}

// LoginProccess with data from Login form
func (c Accounts) LoginProccess(email string, password string) revel.Result {
	c.Validation.Required(email)
	c.Validation.Email(email)
	c.Validation.Required(password)
	var user models.User
	c.Txn.Where("email = ?", email).First(&user)
	if c.Txn.Error != nil {
		panic(c.Txn.Error)
	}
	c.Validation.Required(user.ID != 0).Key("email").Message("Email or Password incorrect")
	err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
	c.Validation.Required(err == nil).Key("password").Message("Email or Password incorrect")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.Flash.Error("Login failed")
		if user.ID == 0 {
			return c.Redirect(routes.Accounts.Register())
		}
		return c.Redirect(routes.Home.Index())
	}
	c.Session["user"] = user.Email
	c.Session["username"] = user.Username
	c.Flash.Success("Welcome, " + user.Username)
	//if user.IsAdmin {
	//	return c.Redirect(routes.Dashboard.Index())
	//}
	return c.Redirect("/")
}

// Register new user
func (c Accounts) Register() revel.Result {
	return c.Render()
}

// RegisterPost is a post of information input on Register
func (c Accounts) RegisterPost(email string, username string, name string, password string, password2 string) revel.Result {
	c.Validation.Required(email)
	c.Validation.Email(email)
	c.Validation.Required(username)
	c.Validation.Required(name)
	c.Validation.Required(password)
	c.Validation.Required(password2)
	c.Validation.Required(password == password2).Key("password").Message("Claves no validas")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		return c.Redirect(routes.Accounts.Register())
	}
	var user = models.User{Email: email, Username: username, Name: name, Active: true, IsAdmin: false}
	user.SetNewPassword(password)
	c.Txn.Create(&user)
	c.Txn.Save(&user)
	c.Session["user"] = user.Email
	return c.Redirect("/")
}

// Logout ot web app
func (c Accounts) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.Home.Index())
}
