package main

import "testing"

type testpair struct {
  values []float64
  result float64
}

var testAverages = []testpair{
  { []float64{1,2}, 1.5 },
  { []float64{1,1,1,1}, 1 },
  { []float64{-2,2}, 0 },
}

var testAddem = []testpair{
  { []float64{4,4,4}, 12},
  { []float64{5,9,20}, 34},
  { []float64{9,204}, 213},
}

func TestAverage(t *testing.T) {
  for _, pair := range testAverages {
    v := Average(pair.values)
    if v != pair.result {
      t.Error(
        "For", pair.values,
        "expected", pair.result,
        "got", v,
      )
    }
  }
}

func Average(xs []float64) float64 {
  total := float64(0)
  for _, x := range xs {
    total += x
  }
  return total / float64(len(xs))
}

func TestAddEm(t *testing.T) {
  for _, pair := range testAddem {
    v := AddEm(pair.values)
    if v != pair.result {
      t.Error(
        "For", pair.values,
        "expected", pair.result,
        "got", v,
      )
    }
  }
}

func AddEm(xs []float64) float64 {
  total := float64(0)
  for _, x := range xs {
    total += x
  }
  return total
}
