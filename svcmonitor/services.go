package main

type service struct {
	name   string
	port   int
	domain string
}

func getServices() map[string]service {
	services := make(map[string]service)
	services["sysmetrics"] = service{
		name:   "sysmetrics",
		port:   3000,
		domain: "http://localhost:3000",
	}
	return services
}
