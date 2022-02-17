package square

import (
	"math"
	"testing"
)

func TestCalcSquareForCircle(t *testing.T) {
	t.Log("Check square for circle")

	expect := math.Pi * math.Pow(10, 2)
	result := CalcSquare(10, 0)

	if result != expect {
		t.Error("Circle's square was wrong calculated")
	}
}

func TestCalcSquareForTriangle(t *testing.T) {
	t.Log("Check square for triangle")

	expect := (math.Pow(10, 2) * math.Sqrt(3)) / 4
	result := CalcSquare(10, 3)

	if result != expect {
		t.Error("Triangle's square was wrong calculated")
	}
}

func TestCalcSquareForSquare(t *testing.T) {
	t.Log("Check square for square")

	expect := math.Pow(10, 2)
	result := CalcSquare(10, 4)

	if result != expect {
		t.Error("Square's square was wrong calculated")
	}
}

func TestCalcSquareForDefault(t *testing.T) {
	t.Log("Check square for square")

	expect := 0.0
	result := CalcSquare(10, 10)

	if result != expect {
		t.Error("The result is unexpected")
	}
}
