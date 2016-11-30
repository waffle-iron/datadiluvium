package main

import "testing"

func TestThingExists(t *testing.T) {
  schemaRdbms := getPages()[0]
  expectedResult := "rdbms"

  if schemaRdbms.Schema != expectedResult {
    t.Error(
      "For", schemaRdbms.Schema,
      "expected", expectedResult,
      "got", schemaRdbms.Schema,
    )
  }
}
