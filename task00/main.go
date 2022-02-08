package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Animal interface declare contract
// all inheritors must implement the Sound method
type Animal interface {
	Sound() string
}

/* =========================================================== */

// Cat
// Some data type
type Cat struct{}

// Sound
// Duck-typing inheritance,
// if Cat implements Sound method, that means Cat is an Animal
func (Cat) Sound() string {
	return "Meow"
}

/* =========================================================== */

// Dog
// Another data type
type Dog struct{}

// Sound
// There Duck-typing works too
func (Dog) Sound() string {
	return "Bark"
}

/* =========================================================== */

// Cow
// Another data type
type Cow struct{}

// Sound
// Another one animal
func (Cow) Sound() string {
	return "Moo"
}

/* =========================================================== */

// CreateAnimal
// Implements fabric method pattern
// ToDo: use constants or enum class instead magic numbers
func CreateAnimal(code int) Animal {
	switch code {
	case 0:
		return Cat{}
	case 1:
		return Cow{}
	case 2:
		return Dog{}
	default:
		panic("Unknown animal code")
	}
}

func GetRandomAnimalCode() int {
	rand.Seed(time.Now().UnixNano())
	min, max := 0, 3
	return rand.Intn(max - min)
}

func CreateRandomAnimals(count int) []Animal {
	var animals []Animal
	for i := 0; i < count; i++ {
		animals = append(animals, CreateAnimal(GetRandomAnimalCode()))
	}
	return animals
}

func CallAnimals(zoo *[]Animal) {
	for _, animal := range *zoo {
		fmt.Println(animal.Sound())
	}
}

func main() {
	zoo := CreateRandomAnimals(20)
	CallAnimals(&zoo)
}
