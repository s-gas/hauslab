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
		port:   1024,
		domain: "http://localhost:1024/metrics",
	}
	return services
}
