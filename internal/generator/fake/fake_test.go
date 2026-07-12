package fake_test

import (
	"context"
	"testing"

	"github.com/agpenton/dataset-factory/internal/generator/fake"
)

func TestGenerate(t *testing.T) {
	g := fake.New("What is Kubernetes?")

	answer, err := g.Generate(
		context.Background(),
		"ignored prompt",
	)
	if err != nil {
		t.Fatal(err)
	}

	if answer != "What is Kubernetes?" {
		t.Fatalf("unexpected answer %q", answer)
	}
}
