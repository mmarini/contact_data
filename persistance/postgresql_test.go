package persistance

import (
	"testing"
)

func TestPrepareInsertStatement(t *testing.T) {

	columns := []string{"column_a", "column_b", "column_c"}

	statement := prepareInsertStatement("temp_table", columns)

	expectedResult := `INSERT INTO "temp_table"("column_a","column_b","column_c") VALUES ($1,$2,$3) RETURNING id`

	if statement != expectedResult {
		t.Fatalf("Did not get expected statement. Got: %s", statement)
	}
}
