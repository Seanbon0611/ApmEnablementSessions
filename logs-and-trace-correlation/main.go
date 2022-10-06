package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	tracer.Start(
		tracer.WithEnv("test-service"),
		tracer.WithService("test-go"),
		tracer.WithServiceVersion("0.0.1"),
	)
	log.SetFormatter(&log.JSONFormatter{})

	defer tracer.Stop()

	//Creates log file
	f, err := os.OpenFile("logpath.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.SetOutput(f)

	//handles routing
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":8080", r)
}

//function that renders home path
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if span, ok := tracer.SpanFromContext(r.Context()); ok {
		traceID := span.Context().TraceID()
		spanID := span.Context().TraceID()
		log.Printf("dd.service:test-go, dd.env:test-service, dd.version=0.0.1, dd.trace_id: %v, dd.span_id: %v", traceID, spanID)
	}
	fmt.Fprintf(w, "Welcome to Gopherdog! ʕ◔ϖ◔ʔ")
}
