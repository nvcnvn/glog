package main

import (
	"github.com/kidstuff/toys/secure/membership"
	"time"

//"github.com/kidstuff/toys/util/forms"
)

func Regis(c *controller) {
	user, err := c.auth.GetUser()
	if err == nil && user.GetId().Valid() {
		c.Redirect(c.BasePath("/"), 307)
		return
	}
	c.View("regis_form.tmpl", c.ViewData("Regis"))
}

func Regis2(c *controller) {
	user, err := c.auth.GetUser()
	if err == nil && user.GetId().Valid() {
		c.Redirect(c.BasePath("/"), 307)
		return
	}
	//vdata := c.ViewData("Regis")

	email := c.Post("Email", true)
	pwd := c.Post("password", false)

	if pwd != c.Post("repassword", false) {
		c.Print("Password doesnot match")
		return
	}

	info := membership.UserInfo{}
	info.FirstName = c.Post("Info.FirstName", true)
	info.MiddleName = c.Post("Info.MiddleName", true)
	info.LastName = c.Post("Info.LastName", true)
	info.NickName = c.Post("Info.NickName", true)
	info.BirthDay, err = time.Parse("2006-01-02", c.Post("Info.BirthDay", false))
	if err != nil {
		c.Print("Worng BirthDay format")
		return
	}

	pri := map[string]bool{}

	if _, err := c.auth.AddUserDetail(email, pwd, &info, pri,
		false, true); err != nil {
		c.Print("some error when regis...", err.Error())
		return
	}
}
