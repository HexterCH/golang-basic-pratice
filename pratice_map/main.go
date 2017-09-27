package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var map_with_literals = map[string]Vertex{
	"Google": Vertex{
		37.42202, -122.08408,
	},
	"Facebook": Vertex{
		36.42202, -122.08408,
	},
	"Amazon": { 30.42202, -122.08408 },
}

func main() {
	// The make function returns a map of the given type, initialized and ready for use
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(map_with_literals)
}