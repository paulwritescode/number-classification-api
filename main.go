package main

import (
	"fmt"

	route "github.com/paulwritescode/numbers-api/router"
)

func main() {
	fmt.Println("Numbers classification API running...")
	fmt.Println("http://localhost:3000")
	route.Route()
}
