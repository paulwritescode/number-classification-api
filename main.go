package main

import (
	"log"

	route "github.com/paulwritescode/numbers-api/router"
)

func main() {
	log.Println("Numbers classification API running on port :3000")
	route.Route()
}
