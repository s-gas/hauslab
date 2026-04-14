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
	services["prometheus"] = service{
		name:   "prometheus",
		port:   9090,
		domain: "http://localhost:9090",
	}
	services["grafana"] = service{
		name:   "grafana",
		port:   3000,
		domain: "http://localhost:3000",
	}
	return services
}
