package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const EXPECTED_MESSAGE = "Expected: %v, got: %v"

func TestAddition(t *testing.T) {
	arith := New()

	answer, err := arith.Addition(1, 1)
	if err != nil {
		t.Fatalf(EXPECTED_MESSAGE, nil, err)
	}

	require.Equal(t, answer, int32(2))
}

func TestSubstraction(t *testing.T) {
	arith := New()

	answer, err := arith.Subtraction(1, 1)
	if err != nil {
		t.Fatalf(EXPECTED_MESSAGE, nil, err)
	}

	require.Equal(t, answer, int32(0))
}

func TestMultiplication(t *testing.T) {
	arith := New()

	answer, err := arith.Multiplication(1, 1)
	if err != nil {
		t.Fatalf(EXPECTED_MESSAGE, nil, err)
	}

	require.Equal(t, answer, int32(1))
}

func TestDivision(t *testing.T) {
	arith := New()

	answer, err := arith.Division(1, 1)
	if err != nil {
		t.Fatalf(EXPECTED_MESSAGE, nil, err)
	}

	require.Equal(t, answer, int32(1))
}
