package interactions_test

import (
	"testing"
	"go-specs-greet/domain/interactions"
	"go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t,
		specifications.GreetAdapter(interactions.Greet))
}
