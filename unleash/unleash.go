package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/Unleash/unleash-client-go/v3"
	"github.com/Unleash/unleash-client-go/v3/context"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Feature struct {
	FeatureName string `json:"feature_name"`
	IsEnabled   bool   `json:"is_enabled"`
	Variant     string `json:"variant"`
}

var featureName = "Log-analyzer"

func init() {
	unleash.Initialize(
		unleash.WithListener(&unleash.DebugListener{}),
		unleash.WithAppName("my-application"),
		unleash.WithUrl("http://localhost:4242/api/"),
		unleash.WithCustomHeaders(http.Header{"Authorization": {"d840cc631ad18fadc043f79277ebec107ef11e305fd2e90948abc8abf33f8d09"}}),
	)

}

func FeatureWithContext(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query().Get("id")
	var ctx context.Context
	if params != "" {
		ctx = context.Context{
			UserId: params,
		}
	}

	var feature Feature

	enabled := unleash.IsEnabled(featureName, unleash.WithContext(ctx))
	variant := unleash.GetVariant(featureName)

	feature.FeatureName = featureName
	feature.IsEnabled = enabled
	feature.Variant = variant.Name

	if enabled {
		displayFile(feature, w)

	} else {
		displayFile(feature, w)
	}

	r.Body.Close()
	r.Header.Set("Connection", "close")
}

func WithSessionID(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query().Get("id")
	var ctx context.Context

	ctx = context.Context{
		UserId:    params,
		SessionId: "1",
	}

	var feature Feature

	enabled := unleash.IsEnabled(featureName, unleash.WithContext(ctx))
	variant := unleash.GetVariant(featureName)

	feature.FeatureName = featureName
	feature.IsEnabled = enabled
	feature.Variant = variant.Name

	if enabled {
		displayFile(feature, w)

	} else {
		displayFile(feature, w)
	}

	r.Body.Close()
	r.Header.Set("Connection", "close")
}

func Default(w http.ResponseWriter, r *http.Request) {

	var feature Feature

	enabled := unleash.IsEnabled(featureName)

	feature.FeatureName = featureName
	feature.IsEnabled = enabled

	if enabled {
		displayFile(feature, w)

	} else {
		displayFile(feature, w)
	}

	r.Body.Close()
	r.Header.Set("Connection", "close")
}
func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", Default).Methods("GET")
	myRouter.HandleFunc("/feature", FeatureWithContext).Methods("GET")
	myRouter.HandleFunc("/session", WithSessionID).Methods("GET")
	log.Fatal(http.ListenAndServe(":8085", myRouter))
}

func displayFile(feature Feature, w http.ResponseWriter) {
	t := template.Must(template.ParseFiles("./basic.html"))
	if err := t.Execute(w, feature); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
