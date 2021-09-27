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
		unleash.WithCustomHeaders(http.Header{"Authorization": {"66e1e4cc7912b243a2710a1851ffd91ed4d68728261ab266aec88d4f574d23e5"}}),
	)

}

func GETHandler(w http.ResponseWriter, r *http.Request) {
	var featureName =  "country-based-rollout"
	if unleash.IsEnabled(featureName) {
		var feature Feature
		feature.FeatureName = featureName
		feature.IsEnabled = unleash.IsEnabled(featureName)

		t := template.Must(template.ParseFiles("./basic.html"))
		if err := t.Execute(w, feature); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		var feature Feature
		feature.FeatureName = featureName
		feature.IsEnabled = unleash.IsEnabled(featureName)
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
