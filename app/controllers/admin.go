package controllers

import (
	"github.com/revel/revel"
)

// Dashboard Controller
type Dashboard struct {
	AdminApp
}

// Index method
func (c Dashboard) Index() revel.Result {
	return c.Render()
}
