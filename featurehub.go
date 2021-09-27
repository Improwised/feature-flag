package main

import (
	"log"
	"net/http"
	"text/template"

	client "github.com/featurehub-io/featurehub-go-sdk"

	"github.com/sirupsen/logrus"
)

type Feature struct {
	FeatureName string `json:"feature_name"`
	IsEnabled   bool   `json:"is_enabled"`
}

const (
	listenAddress = ":8082"
	logLevel      = logrus.TraceLevel
	sdkKey        = "default/98f34bcc-8367-4672-8388-ef4bed619150/Jerc5JkSQTA7rJl28q155kUzYBP6OblR4Zd1K1Vv"
	serverAddress = "http://localhost:8085"
)

func main() {
	http.HandleFunc("/", GETHandler)
	log.Fatal(http.ListenAndServe(listenAddress, nil))


}

func GETHandler(w http.ResponseWriter, r *http.Request) {

	// Set up a TRACE level logger:
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	// Prepare a config:
	fhConfig, err := client.New(serverAddress, sdkKey).WithLogLevel(logLevel).WithWaitForData(true).Connect()
	if err != nil {
		logrus.Fatalf("Error creating config")
	}

	fhClient := fhConfig.NewContext()

	// Configure a logging analytics collector:
	//fhClient.AddAnalyticsCollector(analytics.NewLoggingAnalyticsCollector(logger))
	var featureName =  "FET_1"
	someJSON, err := fhClient.GetFeature(featureName)
	if err != nil {
		log.Fatalf("Error retrieving a JSON feature: %s", err)
	}
	var feature Feature
	feature.FeatureName = featureName
	feature.IsEnabled = true;

	t := template.Must(template.ParseFiles("./basic.html"))
	if err := t.Execute(w, feature); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
}	