package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	listen = flag.String("listen", ":9055", "Host and port to listen on")
)

func readInformation(address string) (map[string]string, error) {
	information := map[string]string{}

	url := fmt.Sprintf("http://%s/etc/mnt_info.csv", address)
	response, err := http.Get(url)
	if err != nil {
		return information, err
	}

	records, err := csv.NewReader(response.Body).ReadAll()
	if err != nil {
		return information, err
	}

	for index, name := range records[0] {
		name = strings.ToLower(strings.Replace(name, "%", "percent", 1))
		name = regexp.MustCompile(`\W+`).ReplaceAllString(name, "_")
		information[name] = records[1][index]
	}

	return information, nil
}

func collectMetrics(address string) *prometheus.Registry {
	success := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name:        "success",
		ConstLabels: map[string]string{"host": address},
	}, []string{})
	registry := prometheus.NewRegistry()
	registry.Register(success)

	information, err := readInformation(address)
	if err != nil {
		success.WithLabelValues().Set(0)
		log.Printf("%#v", err)
		return registry
	}

	success.WithLabelValues().Set(1)

	for name, value := range information {
		floatValue, err := strconv.ParseFloat(value, 64)
		if err == nil {
			registry.Register(prometheus.NewGaugeFunc(prometheus.GaugeOpts{
				Name:        fmt.Sprintf("brother_%s", name),
				ConstLabels: map[string]string{"host": address},
			}, func() float64 {
				return floatValue
			}))
		}
	}

	return registry
}

func main() {
	flag.Parse()

	http.HandleFunc("/metrics", func(response http.ResponseWriter, request *http.Request) {
		host := request.URL.Query().Get("host")
		if host == "" {
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte("Query parameters `host` is required"))
			return
		}

		registry := collectMetrics(host)

		promhttp.HandlerFor(
			registry,
			promhttp.HandlerOpts{},
		).ServeHTTP(response, request)
	})

	log.Printf("Starting to listen on %s", *listen)
	log.Fatal(http.ListenAndServe(*listen, nil))
}
