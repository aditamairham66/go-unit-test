package helpers

import (
	"fmt"
	"testing"
)

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
