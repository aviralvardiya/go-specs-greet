package main_test

import (
	"fmt"
	"testing"
	"go-specs-greet/adapters"
	"go-specs-greet/adapters/grpcserver"
	"go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "50051"
		dockerFilePath = "./Dockerfile"
		driver         = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	adapters.StartDockerServer(t, port, dockerFilePath, "grpcserver")
	specifications.GreetSpecification(t, &driver)
	specifications.CurseSepcification(t, &driver)
}
