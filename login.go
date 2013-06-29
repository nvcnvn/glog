package main

func Login(c *controller) {
	user, err := c.auth.GetUser()
	if err == nil && user.GetId().Valid() {
		c.Redirect(c.BasePath("/"), 307)
		return
	}
	c.View("login_form.tmpl", c.ViewData("test"))
}

func Login2(c *controller) {
	if user, err := c.auth.ValidateUser(c.Post("email", false),
		c.Post("password", false)); err == nil {
		err = c.auth.Login(user.GetId(), 60*60)
		if err != nil {
			c.Print("Loggin falied!")
		} else {
			c.Print("Loggin success!")
		}
	}
}

func Logout(c *controller) {
	c.auth.Logout()
	c.Redirect("/", 307)
}
