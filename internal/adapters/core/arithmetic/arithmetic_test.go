package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Addtion(1, 1)
	if err != nil {
		t.Fatalf("Expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(2))
}

func TestSubstraction(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Subtraction(1, 1)
	if err != nil {
		t.Fatalf("Expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(0))
}

func TestMultiplication(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Multiplication(1, 1)
	if err != nil {
		t.Fatalf("Expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(1))
}

func TestDivision(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Divison(1, 1)
	if err != nil {
		t.Fatalf("Expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(1))
}
