package pipelines_test

import (
	"testing"

	"github.com/ppechkurov/sobesrazborgoogle/internal/pipelines"
)

func TestPipeline(t *testing.T) {
	pipelines.Print(t.Context(),
		pipelines.Square(t.Context(),
			pipelines.Generate(t.Context(), 1, 2, 3),
		))
}
