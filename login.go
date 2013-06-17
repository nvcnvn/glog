package main

func Login(c *controller) {
	c.View("login_form.tmpl", c.ViewData("test"))
}

func Login2(c *controller) {
	if user, err := c.auth.ValidateUser(c.Post("email", false),
		c.Post("password", false)); err == nil {
		err = c.auth.LogginUser(user.GetId().Encode(), 60*60)
		if err != nil {
			c.Print("Loggin falied!")
		} else {
			c.Print("Loggin success!")
		}
	}
}
