package namegen

import (
	"testing"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewName()
	}
}

func BenchmarkFill(b *testing.B) {
	var name Name
	var empty Name
	for i := 0; i < b.N; i++ {
		name = empty
		FillName(&name)
	}
}

// Shoes how to create a two names with the same last name but with different
// genders.
func ExampleFillName_family() {
	// Ask for a male entity
	husbend := Name{Gender: Male}
	// Fill with random data
	FillName(&husbend)
	// Declare a new female entity with the same last name
	wife := Name{Gender: Female, Last: husbend.Last}
	FillName(&wife)
}
