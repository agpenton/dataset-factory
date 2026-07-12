package runtime_test

import (
	"testing"

	"github.com/agpenton/dataset-factory/internal/application/runtime"
)

func TestContextStoresValues(t *testing.T) {
	ctx := runtime.New()

	ctx.Set("answer", "Kubernetes is a container orchestrator.")

	value, ok := ctx.Get("answer")
	if !ok {
		t.Fatal("expected value")
	}

	if value.(string) != "Kubernetes is a container orchestrator." {
		t.Fatal("unexpected value")
	}
}
