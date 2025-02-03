package analysis

import (
	"log"
	"strconv"
)

func ReturnNumber(slug string) int32 {
	//convert string to int
	num, err := strconv.ParseInt(slug, 10, 32)
	if err != nil {
		log.Println(err)
	}

	// convery to int32 safely
	var ApiSlug = int32(num)
	log.Printf("Number: %v\n", ApiSlug)

	//check parity
	status := parity(&ApiSlug)
	log.Println(status)

	return ApiSlug
}
func parity(ApiSlug *int32) bool {
	var is_prime bool

	if *ApiSlug%2 == 0 {
		log.Println("even number")
		is_prime = false
	} else {
		is_prime = true
		log.Println("odd number")
	}

	return is_prime
}
