// +build example

package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	flipt "github.com/markphelps/flipt-grpc-go"

	"google.golang.org/grpc"
)

type FetureFlag struct {
	FlagKey     string
	FlagName    string
	FlagEnabled bool
	Plan        string
	ErrorMsg    string
	IsError     bool
}

var (
	fliptServer string
	flagKey     string
)

func init() {
	flag.StringVar(&fliptServer, "server", ":9000", "address of Flipt backend server")
	flag.StringVar(&flagKey, "plan-rollout", "plan-rollout", "flag key to query")
}

func main() {
	log.Println("demo ui available at http://localhost:8000")
	log.Printf("flag key: %s\n", flagKey)
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", Default).Methods("GET")
	myRouter.HandleFunc("/plan", Plan).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", myRouter))

}

func Default(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("plan.html"))
	data := FetureFlag{
		IsError:  true,
		ErrorMsg: "Enter http://localhost:8000/plan?name=YourPlanName",
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Plan(w http.ResponseWriter, r *http.Request) {
	flag.Parse()

	conn, err := grpc.Dial(fliptServer, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Printf("connected to Flipt server at: %s", fliptServer)

	client := flipt.NewFliptClient(conn)
	params := r.URL.Query().Get("name")
	t := template.Must(template.ParseFiles("plan.html"))

	var data FetureFlag
	plan := params
	flagKey := "plan-rollout"
	entityId := "18568591hhuyuyuyyyyyyftgfsesxfc"
	flag, err := client.Evaluate(context.Background(), &flipt.EvaluationRequest{
		FlagKey:  flagKey,
		EntityId: entityId,
		Context: map[string]string{
			"plan": plan,
		},
	})

	if flag.GetMatch() == true && flag.GetSegmentKey() == plan {
		data = FetureFlag{
			FlagKey:     flagKey,
			FlagName:    flagKey,
			FlagEnabled: flag.GetMatch(),
			Plan:        plan,
		}
	} else {
		data = FetureFlag{
			IsError:  true,
			ErrorMsg: "You need to register plan",
		}
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if flag == nil {
		http.Error(w, "flag not found", http.StatusNotFound)
		return
	}

	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
