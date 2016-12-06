package main

import (
  "os"
  "fmt"
  "testing"
)

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


func TestEnvironmentVariable(t *testing.T) {
  result := os.Getenv("VERIFICATION_PASS")

  fmt.Println(result)

  if result != "Yes" {
    t.Error(
      "For", result,
      "expected", "Yes",
      "got", result,
    )
  }
}
