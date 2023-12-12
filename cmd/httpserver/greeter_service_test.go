package main_test

import (
	// "fmt"
	"net/http"
	"testing"
	"time"
	"fmt"
	"go-specs-greet/adapters"
	"go-specs-greet/adapters/httpserver"
	"go-specs-greet/specifications"

)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "8080"
		dockerFilePath = "./Dockerfile"
		baseURL        = fmt.Sprintf("http://localhost:%s", port)

		driver = httpserver.Driver{BaseURL: baseURL, Client: &http.Client{
			Timeout: 1 * time.Second,
		}}
	)

	adapters.StartDockerServer(t, port, dockerFilePath, "httpserver")
	specifications.GreetSpecification(t, driver)
	specifications.CurseSepcification(t, driver)
}
