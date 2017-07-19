package controllers

import (
	//"github.com/opennut/gorum/app/models"
	"regexp"

	"github.com/revel/revel"
)

// Home Controller anon user
type Home struct {
	PublicApp
}

// Index page
func (c Home) Index() revel.Result {
	return c.Render()
}

// Detail page
func (c Home) Detail(ID int, name string) revel.Result {
	c.Validation.Required(name).Key("name").Message("El nombre es requerido")
	c.Validation.Required(name)
	c.Validation.MaxSize(name, 15)
	c.Validation.MinSize(name, 4)
	c.Validation.Match(name, regexp.MustCompile("^\\w*$"))
	c.Validation.Keep()

	return c.Render()
}
