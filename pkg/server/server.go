package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sahlinet/go-tumbo/pkg/worker"

	"github.com/gorilla/mux"
)

func getHomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("homepage")
}

func startWorker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	worker := worker.Worker{}
	worker.Start()
	w.Header().Set("Content-Type", "application/json")
	resp := fmt.Sprintf("Worker %s started", vars["worker"])
	json.NewEncoder(w).Encode(resp)
}

// func getWorker(w http.ResponseWriter, r *http.Request) {
// 	worker := worker.Worker{}
// 	worker.Start()
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(resp)
// }

type ResponsePayload struct {
	Message string
}

func Worker() Handler {
	return Handler{
		Route: func(r *mux.Route) {
			r.Path("/{worker}/start").Methods("POST")
		},
		Func: func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			workerName := mux.Vars(r)["worker"]
			worker := worker.Worker{}
			worker.Start()
			resp := ResponsePayload{fmt.Sprintf("Worker %s started", workerName)}
			json.NewEncoder(w).Encode(resp)
		},
	}
}

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/", getHomePage)
	//r.HandleFunc("/{worker}", getWorker)
	Worker().AddRoute(r)

	r.HandleFunc("/{worker}/start", startWorker)
	http.Handle("/", r)

	// // This will serve files under http://localhost:8000/static/<filename>
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	srv := &http.Server{
		Handler: r,
		//Addr:    "127.0.0.1:8000",
		Addr: "0.0.0.0:8001",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
