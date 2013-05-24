package main

func (c *controller) ViewData(title string) map[string]interface{} {
	m := make(map[string]interface{})
	m["Title"] = title
	m["BasePath"] = func(path string) string { return c.BasePath(path) }
	return m
}
