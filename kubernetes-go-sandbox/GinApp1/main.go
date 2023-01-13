package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

type DDContextLogHook struct{}

func (d *DDContextLogHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel, logrus.WarnLevel, logrus.InfoLevel, logrus.DebugLevel, logrus.TraceLevel}
}

func (d *DDContextLogHook) Fire(e *logrus.Entry) error {
	span, found := tracer.SpanFromContext(e.Context)
	if !found {
		return nil
	}
	e.Data["dd.trace_id"] = span.Context().TraceID()
	e.Data["dd.span_id"] = span.Context().SpanID()
	return nil
}

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.AddHook(&DDContextLogHook{})
	logrus.SetOutput(os.Stdout)

	// Start the tracer and defer the Stop method.
	tracer.Start(
		// tracer.WithDebugMode(true),
		tracer.WithEnv("sandbox"),
		tracer.WithService("ginapp1"),
		tracer.WithServiceVersion("1.0.0"),
		tracer.WithRuntimeMetrics(),
	)
	defer tracer.Stop()

	profiler.Start(
		profiler.WithService("ginapp1"),
		profiler.WithEnv("sandbox"),
		profiler.WithVersion("1.0.0"),
		profiler.WithProfileTypes(
			profiler.CPUProfile,
			profiler.HeapProfile,
			// The profiles below are disabled by default to keep overhead
			// low, but can be enabled as needed.

			// profiler.BlockProfile,
			// profiler.MutexProfile,
			// profiler.GoroutineProfile,
		),
	)

	defer profiler.Stop()

	// Create a gin.Engine
	r := gin.New()

	// Use the tracer middleware with your desired service name.
	r.Use(gintrace.Middleware("ginapp1"))

	// Set up some endpoints.
	r.GET("/", func(c *gin.Context) {
		ctx := c.Request.Context()
		if span, ok := tracer.SpanFromContext(ctx); ok {
			// Set tag
			span.SetTag("test_tag", "sucessful")
		}
		cLog := logrus.WithContext(ctx)

		cLog.Info("Completed some work!")
		c.String(200, "Welcome to Gopherdog! ʕ◔ϖ◔ʔ")
	})

	//Distributed Tracing
	r.GET("/test", func(c *gin.Context) {
		ctx := c.Request.Context()
		if span, ok := tracer.SpanFromContext(ctx); ok {
			fmt.Printf("The SpanID is %v", span.Context().SpanID())
			fmt.Printf(" The TraceID is %v", span.Context().TraceID())
			var url string
			kubernetesGinApp2Ip := os.Getenv("KUBERNETES_GINAPP_IP")
			if kubernetesGinApp2Ip == "" {
				url = "http://localhost:8081/service"
			} else {
				url = "http://" + kubernetesGinApp2Ip + "/service"
			}
			//Makes request to GinApp2 endpoint
			req, err := http.NewRequest("GET", url, nil)
			req = req.WithContext(ctx)
			// Inject the span Context in the Request headers
			err = tracer.Inject(span.Context(), tracer.HTTPHeadersCarrier(req.Header))
			if err != nil {
				c.String(500, "Internal Server Error")
				log.Fatalf("the error was: %s", err)
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Printf("client: error making http request: %s\n", err)
				os.Exit(1)
			}

			resBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Printf("client: could not read response body: %s\n", err)
				os.Exit(1)
			}
			c.String(200, string(resBody))
		}
	})
	//Error Tracking
	r.GET("/error", func(c *gin.Context) {

		ctx := c.Request.Context()
		if span, ok := tracer.SpanFromContext(ctx); ok {
			span.SetTag("track_error", true)
			defer span.Finish(tracer.WithError(errors.New("error upon request")))
		}
	})

	r.Run(":8080")
}
