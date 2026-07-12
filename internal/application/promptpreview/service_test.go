package promptpreview_test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/agpenton/dataset-factory/internal/application/promptpreview"
)

func TestRenderInstructionPrompt(t *testing.T) {
	service := promptpreview.New()

	repoRoot, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}

	recipe := filepath.Join(
		repoRoot,
		"recipes",
		"instruction-from-answer.yaml",
	)

	text, err := service.Render(
		recipe,
		"Kubernetes is a container orchestration platform.",
	)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(
		text,
		"Kubernetes is a container orchestration platform.",
	) {
		t.Fatal("answer not rendered into template")
	}

	if !strings.Contains(
		text,
		"You are creating supervised fine-tuning datasets.",
	) {
		t.Fatal("template not rendered")
	}
}

func TestMissingRecipe(t *testing.T) {
	service := promptpreview.New()

	_, err := service.Render(
		"does-not-exist.yaml",
		"answer",
	)

	if err == nil {
		t.Fatal("expected error")
	}
}
