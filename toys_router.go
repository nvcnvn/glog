package main

func (h *handler) initSubRoutes() {
	h._defaultHandle = Home
	h._subRoutes = []route{
		route{"login", Login},
		route{"login2", Login2},
		route{"thread", ViewThread},
		route{"newthread", NewThread},
		route{"newthread2", NewThread2},
		route{"feed", Feed},
		route{"newcat", NewCat},
		route{"newcat2", NewCat2},
	}
}
