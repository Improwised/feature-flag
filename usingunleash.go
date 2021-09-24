package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/Unleash/unleash-client-go/v3"
	_ "github.com/lib/pq"
)

type Feature struct {
	FeatureName string `json:"feature_name"`
	IsEnabled   bool   `json:"is_enabled"`
}

func init() {
	unleash.Initialize(
		unleash.WithListener(&unleash.DebugListener{}),
		unleash.WithAppName("my-application"),
		unleash.WithUrl("http://localhost:4242/api/"),
		unleash.WithCustomHeaders(http.Header{"Authorization": {"93fcc4689e6b76833d796cb616625c5f1747b97f05f9e0cf3ef661b4d36f3090"}}),
	)

}

func GETHandler(w http.ResponseWriter, r *http.Request) {

	if unleash.IsEnabled("Log-analyzer") {
		var feature Feature
		feature.FeatureName = "Log-analyzer"
		feature.IsEnabled = unleash.IsEnabled("Log-analyzer")

		t := template.Must(template.ParseFiles("./basic.html"))
		if err := t.Execute(w, feature); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		var feature Feature
		feature.FeatureName = "Log-analyzer"
		feature.IsEnabled = unleash.IsEnabled("Log-analyzer")
		t := template.Must(template.ParseFiles("./basic.html"))
		if err := t.Execute(w, feature); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func main() {
	http.HandleFunc("/", GETHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
