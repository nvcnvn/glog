package main

func Home(c *controller) {
	user, err := c.auth.GetUser()
	if err != nil {
		c.Redirect(c.BasePath("login"), 303)
		return
	}

	c.Print(user.Email)
}
