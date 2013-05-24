package main

func (h *handler) initSubRoutes() {
	h._defaultHandle = Home
	h._subRoutes = []route{
		route{"login", Login},
		route{"login2", Login2},
		route{"newthread", NewThread},
		route{"newthread2", NewThread2},
		route{"feed", Feed},
	}
}
