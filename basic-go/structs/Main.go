package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	fmt.Println("Structs")

	aDoctor := Doctor{
		number:    3,
		actorName: "jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	fmt.Printf("Doctor struct: %v\n", aDoctor)

	// Getting values from a struct we use the dot syntax
	fmt.Printf("Actor name: %v\n", aDoctor.actorName)
	fmt.Println()

	// another way to instantiate a struct is using the positional syntax
	// aDoctor := Doctor{
	// 		3,
	// 		"jon Pertwee",
	// 		[]string{
	// 			"Liz Shaw",
	// 			"Jo Grant",
	// 			"Sarah Jane Smith",
	// 		},
	// }

	// anonymous structs
	doctor := struct{ name string }{name: "Jon Pertwee"}
	fmt.Printf("Anonymous struct:  %v\n", doctor)
	fmt.Println()

	// structs are passed by value, whole struct is copied over
	anotherDoctor := doctor
	anotherDoctor.name = "Tom Walker"
	fmt.Printf("doctor: %v\n", doctor)
	fmt.Printf("anotherDoctor: %v\n", anotherDoctor)
	fmt.Println()

	// golang does not support inheritance, it support compositon instead
	// inheritance defines "is-a" relationship
	// composition defines "has-a" relationship
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	// could also be instantiated in the following way
	// using the embedding syntax
	// b := Bird {
	// 	Animal: Animal{Name: "Emu", Origin: "Australia"},
	// 	SpeedKPH: 48,
	// 	CanFly: false,
	// }

	fmt.Printf("Bird: %v\n", b)
}

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}
