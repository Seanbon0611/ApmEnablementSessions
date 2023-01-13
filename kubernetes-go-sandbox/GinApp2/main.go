package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

func main() {
	// Start the tracer and defer the Stop method.
	tracer.Start(
		//Unified Service Tagging
		tracer.WithEnv("sandbox"),
		tracer.WithServiceVersion("1.0.0"),
		tracer.WithService("ginapp2"),
		tracer.WithRuntimeMetrics(),
	)
	defer tracer.Stop()

	//Enable Go Profiler
	profiler.Start(
		profiler.WithService("ginapp2"),
		profiler.WithEnv("sandbox"),
		profiler.WithVersion("1.0.0"),
		profiler.WithProfileTypes(
			profiler.CPUProfile,
			profiler.HeapProfile,
		),
	)

	defer profiler.Stop()

	// Create a gin.Engine
	r := gin.New()

	// Use the tracer middleware with your desired service name.
	r.Use(gintrace.Middleware("ginapp2"))

	// Set up some endpoints.
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello world!")
	})

	r.GET("/service", func(c *gin.Context) {
		ctx := c.Request
		fmt.Println(ctx)
		sctx, err := tracer.Extract(tracer.HTTPHeadersCarrier(ctx.Header))
		if err != nil {
			c.String(500, "Internal Server Error")
			log.Fatalf("There was an error %s", err)
		}
		span := tracer.StartSpan("test.span", tracer.ChildOf(sctx))
		defer span.Finish()

		c.String(200, "Headers received!")
	})

	// And start gathering request traces.
	r.Run(":8081")
}
