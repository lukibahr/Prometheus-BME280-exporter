package handlers

import (
	"encoding/json"
	"net/http"
)

func IndexHandler (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<html>
		  <head>
		  <title>BME280 exporter</title>
		  </head>
		  <body>
		  <h1>BME280 exporter</h1>
          <p>Go to <a href="/metrics">/metrics</a></p>
		  </body>
		  </html>`))
}

type response struct {
	Status string
}
type message struct {
	Message string
}

func HealthHandler (w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{Status: "up"})
}
func JsonMetricsHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message{Message: "json"})
}