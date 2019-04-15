package main

import "fmt"

func main() {
	s := `{A:"B",C:{D:"E",F:{G:"H",I:"J"}}}"`
}

// {
//     A:"B",
//     C:
//     {
//         D:"E",
//         F:
//         {
//             G:"H",
//             I:"J"
//         }
//     }
// }

func prettyJson(s string) {
	space := " "
	for _, v := range s {
		if string(v) == "{" {
			fmt.Printf("%s", space)
			space += " "
		}
	}

}
