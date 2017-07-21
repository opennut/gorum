package controllers

import (
	"bytes"
	"fmt"
	"github.com/opennut/gorum/app/routes"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"os"
	"path/filepath"
)

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

// Account Controller
type Users struct {
	AdminApp
}

func (c Users) Profile() revel.Result {
	user := c.connected()
	return c.Render(user)
}

func (c Users) ChangeEmail(email string, password string) revel.Result {
	c.Validation.Required(email)
	c.Validation.Email(email)
	c.Validation.Required(password)
	user := c.connected()
	err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
	c.Validation.Required(err == nil).Key("password").Message("Email or Password incorrect")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		return c.Redirect(routes.Accounts.Logout())
	}
	user.Email = email
	c.Txn.Save(&user)
	c.Flash.Success("Email changed")
	return c.Redirect(routes.Users.Profile())
}

func (c Users) ChangePassword(current_password string, password string, password2 string) revel.Result {
	c.Validation.Required(current_password)
	c.Validation.Required(password)
	c.Validation.Required(password2)
	c.Validation.Required(password == password2).Key("password").Message("Claves no validas")
	if c.Validation.HasErrors() {
		return c.Redirect(routes.Users.Profile())
	}
	user := c.connected()
	err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(current_password))
	c.Validation.Required(err == nil).Key("current_password").Message("Email or Password incorrect")
	if c.Validation.HasErrors() {
		return c.Redirect(routes.Users.Profile())
	}
	user.SetNewPassword(password)
	c.Txn.Save(&user)
	c.Flash.Success("Password changed")
	return c.Redirect(routes.Users.Profile())
}

func (c Users) ChangeAvatar(avatar []byte) revel.Result {
	// Validation rules.
	c.Validation.Required(avatar)
	c.Validation.MinSize(avatar, 2*KB).Message("Minimum a file size of 2KB expected")
	c.Validation.MaxSize(avatar, 2*MB).Message("File cannot be larger than 2MB")

	// Check format of the file.
	conf, format, err := image.DecodeConfig(bytes.NewReader(avatar))
	c.Validation.Required(err == nil).Key("avatar").Message("Incorrect file format")
	c.Validation.Required(format == "jpeg" || format == "png").Key("avatar").Message("JPEG or PNG file format is expected")

	// Check resolution.
	c.Validation.Required(conf.Height >= 150 && conf.Width >= 150).Key("avatar").Message("Minimum allowed resolution is 150x150px")
	user := c.connected()
	// Handle errors.
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Users.Profile())
	}
	if len(avatar) > 0 {
		var nameFile = "location.png"
		nameFile = c.Params.Files["avatar"][0].Filename
		fmt.Println(c.Params.Files["avatar"][0].Open())
		originalFile, err := c.Params.Files["avatar"][0].Open()
		if err != nil {
			log.Fatal(err)
		}
		defer originalFile.Close()
		// Create new file
		newFile, err := os.Create(filepath.Join(revel.BasePath, "public", "uploads", nameFile))
		if err != nil {
			log.Fatal(err)
		}
		defer newFile.Close()
		// Copy the bytes to destination from source
		bytesWritten, err := io.Copy(newFile, originalFile)
		if err != nil {
			log.Fatal(err)
		}
		user.FileName = "uploads/" + nameFile
		c.Txn.Save(&user)
	}

	return c.Redirect(routes.Users.Profile())
}
