package main

import (
	"github.com/kidstuff/toys/model"
	"github.com/kidstuff/toys/secure/membership"
	"strconv"
	"strings"
)

func Member(c *controller) {
	path := strings.Split(c.Request.URL.Path, "/")
	switch path[len(path)-1] {
	case "all":
		AllMember(c)
	case "online":
		OnlineMember(c)
	case "find":
		FindMemberResult(c)
	case "info":
		MemberInfo(c)
	default:
		AllMember(c)
	}
}

func AllMember(c *controller) {
	limit, err := strconv.Atoi(c.Get("limit", false))
	if err != nil {
		limit = 20
	}

	var userLst membership.UserLister

	offsetStr := c.Get("offsetId", false)
	offsetId, err := model.MustLoad(MODELDRIVER).DecodeId(offsetStr)
	if err == nil {
		userLst, err = c.auth.FindAllUser(offsetId, limit)
	} else {
		userLst, err = c.auth.FindAllUser(nil, limit)
	}

	if err == nil {
		data := c.ViewData("Member List")
		data["UserLst"] = userLst
		c.View("allmember_list.tmpl", data)
	}
}

func OnlineMember(c *controller) {
}

func FindMemberResult(c *controller) {
}

func MemberInfo(c *controller) {
}
