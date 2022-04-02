// Maps: All keys are of same type, & all values are of same type.
package main

import "fmt"

// func main() {
// 	// Different ways of declaring a map.
// 	// 1) map[<key-type>]<value-type>
// 	colors := map[string]string{
// 		"red":   "#FF0000",
// 		"green": "#45gftr",
// 	}
// 	fmt.Println(colors) // map[green:#45gftr red:#FF0000]

// 	// 2)
// 	var colors2 map[string]string
// 	fmt.Println(colors2) // map[]

// 	// 3) Built-in function
// 	colors3 := make(map[int]string)
// 	colors3[10] = "#QWE345"
// 	colors3[11] = "#TYH%$#"
// 	delete(colors3, 11)

// 	fmt.Println(colors3) // map[10:#QWE345]

// }

// ---------- How to iterate over a map? --------
func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#45gftr",
		"white": "#FFFFFF",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	// for key, value : range map {}
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}

// ------ Structs vs Maps ------
// Map - Reference type
// Struct - Value type
