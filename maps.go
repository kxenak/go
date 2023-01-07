package main

import "fmt"

func main() {
	// Basic map.
	map_map := make(map[string]string)
	fmt.Println(map_map)
	map_map["l"] = "L bro."
	map_map["g"] = "Top G."
	fmt.Println(map_map)

	// Length method.
	fmt.Println("Length of map_map: ", len(map_map))

	// Delete method.
	delete(map_map, "l")
	fmt.Println(map_map)

	// The second return value indicates if a key exists or not.
	_, tell := map_map["l"]
	fmt.Println(tell)

	// Declaring and initializing together.
	maap := map[string]string{"foo": "poo", "boo": "loo"}
	fmt.Println(maap)
}
