package main

func (h *handler) initSubRoutes() {
	h._defaultHandle = Home
	h._subRoutes = []route{
		route{pattern: "login", fn: Login},
		route{pattern: "login2", fn: Login2},
		route{pattern: "logout", fn: Logout},
		route{pattern: "regis", fn: Regis},
		route{pattern: "regis2", fn: Regis2},
		route{pattern: "members/*", fn: Member},
		route{pattern: "thread", fn: ViewThread},
		route{pattern: "newthread", fn: NewThread},
		route{pattern: "newthread2", fn: NewThread2},
		route{pattern: "feed", fn: Feed},
		route{pattern: "newcat", fn: NewCat},
		route{pattern: "newcat2", fn: NewCat2},
	}
}
