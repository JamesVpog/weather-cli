package main

import (
	"fmt"
	"strconv"

	"github.com/JamesVpog/weather-cli/internal/weather"
)

func main() {
	weather.Hello()
	num, _ := strconv.Atoi("not a number")
	fmt.Println(num)
}
