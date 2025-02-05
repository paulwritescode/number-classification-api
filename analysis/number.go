package analysis

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

type NumberReturnJson struct {
	Number     int      `json:"number"`
	Prime      bool     `json:"is_prime"`
	Perfect    bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FuncFact   string   `json:"fun_fact"`
}

func ReturnNumber(slug int) NumberReturnJson {
	var ApiSlug = slug

	//check parity
	status := parity(&ApiSlug)

	// check perfect number
	perfectN := isPerfect(ApiSlug)

	//properties
	PROPS := Properties(ApiSlug)

	// sum of all digits
	SUMDIGITS := digitSum(ApiSlug)

	//fun fact
	FUNFACT := getFunFact(ApiSlug)

	JsonReturn := NumberReturnJson{
		Number:     ApiSlug,
		Prime:      status,
		Perfect:    perfectN,
		Properties: PROPS,
		DigitSum:   SUMDIGITS,
		FuncFact:   FUNFACT,
	}

	log.Println(JsonReturn)
	return JsonReturn
}

func getFunFact(ApiSlug int) string {
	url := fmt.Sprintf("http://numbersapi.com/%d/math", ApiSlug)
	responce, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer responce.Body.Close()

	body, err := io.ReadAll(responce.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

func digitSum(ApiSlug int) int {
	string_number := strconv.Itoa(ApiSlug)
	sum := 0
	for _, character := range string_number {
		character_int, _ := strconv.Atoi(string(character))
		sum += character_int
	}
	return sum
}

func parity(ApiSlug *int) bool {
	var is_prime bool

	if *ApiSlug%2 == 0 {
		is_prime = false
	} else {
		is_prime = true
	}

	return is_prime
}

func isPerfect(ApiSlug int) bool {
	var perfect_number bool
	sum := 0

	for i := 1; i < ApiSlug; i++ {
		if ApiSlug%i == 0 {
			sum += i
		}
	}

	if sum == ApiSlug {
		perfect_number = true
	} else {
		perfect_number = false
	}

	return perfect_number
}

func Properties(ApiSlug int) []string {
	//armstrong
	original := ApiSlug
	digits := []int{}
	sum := 0

	//extract the individual digits
	for temp := ApiSlug; temp > 0; temp /= 10 {
		digits = append(digits, temp%10)
	}

	power := len(digits)

	//compute sum of digits raised to the power of the count of digits
	for _, digit := range digits {
		sum += int(math.Pow(float64(digit), float64(power)))
	}

	var armstrong_status string
	if sum == original {
		armstrong_status = "armstrong"
	}

	// check even / odd
	odd_even := "odd"
	if ApiSlug%2 == 0 {
		odd_even = "even"
	}

	var propertiesArray []string
	if sum == original {
		propertiesArray = []string{armstrong_status, odd_even}
	} else {
		propertiesArray = []string{odd_even}
	}

	return propertiesArray
}
