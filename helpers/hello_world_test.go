package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Eko",
			request:  "Eko",
			expected: "Hello World Eko",
		},
		{
			name:     "Kurniawan",
			request:  "Kurniawan",
			expected: "Hello World Kurniawan",
		},
		{
			name:     "Khannedy",
			request:  "Khannedy",
			expected: "Hello World Khannedy",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello World Budi",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello World Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		require.Equal(t, "Hello World Eko", result, "Result must be 'Hello World Eko'")
	})
	t.Run("Kurniawan", func(t *testing.T) {
		result := HelloWorld("Kurniawan")
		require.Equal(t, "Hello World Kurniawan", result, "Result must be 'Hello World Kurniawan'")
	})
	t.Run("Khannedy", func(t *testing.T) {
		result := HelloWorld("Khannedy")
		require.Equal(t, "Hello World Khannedy", result, "Result must be 'Hello World Khannedy'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldKhannedy(t *testing.T) {
	result := HelloWorld("Khannedy")

	if result != "Hello World Khannedy" {
		// error
		t.Fatal("Result must be 'Hello World Khannedy'")
	}

	fmt.Println("TestHelloWorldKhannedy Done")
}

func TestHelloWorldKurniawan(t *testing.T) {
	result := HelloWorld("Kurniawan")

	if result != "Hello World Kurniawan" {
		// error
		t.Fatal("Result must be 'Hello World Kurniawan'")
	}

	fmt.Println("TestHelloWorldKurniawan Done")
}
