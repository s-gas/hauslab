package main

import (
	"fmt"
	"net/http"
)

func writeMetrics(w http.ResponseWriter, cpu float64, ram float64) {
	fmt.Fprintln(w, "# HELP cpu_usage_percent Current CPU usage as a percentage")
	fmt.Fprintln(w, "# TYPE cpu_usage_percent gauge")
	fmt.Fprintf(w, "cpu_usage_percent %.2f\n", cpu)

	fmt.Fprintln(w, "# HELP ram_usage_percent Current RAM usage as a percentage")
	fmt.Fprintln(w, "# TYPE ram_usage_percent gauge")
	fmt.Fprintf(w, "ram_usage_percent %.2f\n", ram)
}
