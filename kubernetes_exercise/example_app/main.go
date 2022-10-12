package main

import (
	"fmt"
	"net/http"

	"gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	tracer.Start(
		tracer.WithEnv("prod"),
		tracer.WithService("example_app"),
		tracer.WithServiceVersion("0.0.1"),
	)
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":80", r)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It's a Homepage Party!!!\n")
	fmt.Fprintf(w, "|-----------------------------------------------------------------------|\n")
	fmt.Fprintf(w, "|    o   | o /  _ o         __|    | /     |__        o _  | o /   o    |\n")
	fmt.Fprintf(w, "|   /||    |     /|   ___|o   |o    |    o/    o/__   /|     |    /||   |\n")
	fmt.Fprintf(w, "|   / |   / |   | |  /)  |    ( |  /o|  / )    |  (|  / |   / |   / |   |\n")
	fmt.Fprintf(w, "|-----------------------------------------------------------------------|\n")

}
