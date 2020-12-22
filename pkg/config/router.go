package config

// 获取当前路由信息
func GetCurrentRoute() map[string]string {
	var route map[string]string
	if Request == nil {
		return route
	}

	return map[string]string{
		"url":    Request.RequestURI,
		"method": Request.Method,
	}
}

// 获取系统全部路由
func GetAllRoutes() []map[string]string {
	appRouters := Router.Routes()
	routers := []map[string]string{}

	for _, route := range appRouters {
		// fmt.Printf("Method: %s, Path: %s \n", route.Method, route.Path)
		routers = append(routers, map[string]string{
			"url":    route.Path,
			"method": route.Method,
		})
	}
	return routers
}
