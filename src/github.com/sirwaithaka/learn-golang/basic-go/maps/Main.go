package main

import (
	"fmt"
)

func main() {
	// map values are passed by reference

	// creating maps
	// keys of a map should be types that can
	// be used for equivalency checking
	countyPopulations := map[string]int{
		"Nairobi":  3193800,
		"Kiambu":   2384044,
		"Machakos": 1039449,
	}

	fmt.Printf("County Populations: \n%v", countyPopulations)
	fmt.Println()

	// using make function to create a map
	classPopulations := make(map[string]int)
	classPopulations = map[string]int{
		"8South":   29,
		"8North":   34,
		"8West":    23,
		"8Central": 30,
		"8East":    26,
	}

	fmt.Printf("Class Populations: \n%v", classPopulations)
	fmt.Println()

	// manipulating values in the map
	// Getting a value from a map
	fmt.Printf("Population in Kiambu: %v\n", countyPopulations["Kiambu"])

	// Adding a value to a map
	countyPopulations["Kitui"] = 389209
	fmt.Printf("Populations with Kitui added: \n%v", countyPopulations)
	fmt.Println()

	// deleting from a map
	delete(countyPopulations, "Machakos")
	fmt.Printf("Populations without Machakos: \n%v", countyPopulations)

	// Accessing a key from a map that isnt in the map will result
	// in a default value
	// in this case accessing a key such as "Garisssa" which is in fact
	// not a key in the :countyPopulations map will return 0
	fmt.Println()
	fmt.Printf("Populations: \n%v", countyPopulations)
	fmt.Println()
	fmt.Printf("Accessing a key which is not in the map <Garissa>: %v\n", countyPopulations["Garissa"])
	fmt.Println()

	// we can use the ok value to check that indeed the key was not in the map
	_, ok := countyPopulations["Garissa"]
	fmt.Printf("Garrissa key: %v in %v\n", ok, countyPopulations)
	fmt.Println()

	// getting the length of a map
	fmt.Printf("Length of countyPopulations: %v\n", len(countyPopulations))
}
