package pipelines_test

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ppechkurov/sobesrazborgoogle/internal/pipelines"
)

func TestPipeline(t *testing.T) {
	pipelines.Print(t.Context(),
		pipelines.Square(t.Context(),
			pipelines.Generate(t.Context(), 1, 2, 3),
		))
}

func TestDatabase(t *testing.T) {
	db, err := sql.Open("sqlite3", "file:test.db?cache=shared&mode=memory")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	rows, err := pipelines.GetRows(t.Context(), db)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	read := pipelines.Read(t.Context(), rows)
	even := pipelines.Filter(t.Context(),
		read,
		func(v pipelines.Row) bool { return v.ID%2 == 0 },
	)

	for r := range even {
		if r.ID%2 != 0 {
			t.Errorf("odd ids should be filtered, got %v", r)
		}
	}
}
