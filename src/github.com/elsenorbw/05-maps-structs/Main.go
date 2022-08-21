package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH int
	CanFly   bool
}

type Something struct {
	Name   string `required:"true" max:"100"`
	Origin string `required:"false"`
}

func tags() {
	t := reflect.TypeOf(Something{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}

func animals() {
	fmt.Println("Animals!")

	// External access once created doesn't care about the Composition of Animal inside Bird
	a := Bird{}
	a.Name = "Emu"
	a.Origin = "Australia"
	a.SpeedKPH = 48
	a.CanFly = false
	fmt.Printf("a=%v,%T\n", a, a)

	// But initialisers does..
	b := Bird{
		Animal:   Animal{Name: "emu", Origin: "Aus"},
		SpeedKPH: 46,
		CanFly:   false,
	}
	fmt.Printf("b=%v,%T\n", b, b)

}

func main() {
	fmt.Println("Maps and structs!")

	dave := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	fmt.Printf("dave=%v, %T\n", dave, dave)

	// Ways to create
	// 1) the literal (see dave)
	// 2) the make function

	dave2 := make(map[string]int)
	dave2["blarg"] = 127
	dave2["blarg"] += 1
	fmt.Printf("dave2=%v, %T\n", dave2, dave2)

	val, ok := dave["three"]
	fmt.Printf("val=%v, ok=%v\n", val, ok)

	delete(dave, "three")
	fmt.Printf("dave=%v, %T\n", dave, dave)

	val, ok = dave["three"]
	fmt.Printf("val=%v, ok=%v\n", val, ok)

	// Maps pass byref

	// Structs
	aDoctor := Doctor{
		number:     3,
		actorName:  "Jon Pertwee",
		companions: []string{"Liz Shaw", "Jo Grant"},
	}
	fmt.Printf("aDoctor=%v\n", aDoctor)
	fmt.Printf("aDoctor.actorName=%v\n", aDoctor.actorName)

	// Can also initialise by position (but don't, maintenance etc)
	bDoctor := Doctor{4, "Dave", []string{"abc", "def"}}
	fmt.Printf("bDoctor=%v\n", bDoctor)
	fmt.Printf("bDoctor.actorName=%v\n", bDoctor.actorName)

	// Structs copy, not byref
	cDoctor := aDoctor
	cDoctor.number = 67
	cDoctor.actorName = "Plavik Venmo"

	fmt.Printf("aDoctor=%v\n", aDoctor)
	fmt.Printf("aDoctor.actorName=%v\n", aDoctor.actorName)
	fmt.Printf("cDoctor=%v\n", cDoctor)
	fmt.Printf("cDoctor.actorName=%v\n", cDoctor.actorName)

	animals()
	tags()
}
