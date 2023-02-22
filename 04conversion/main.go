package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please rate our service from 1 to 5:")
	reader := bufio.NewReader(os.Stdin)

	rating, _ := reader.ReadString('\n')
	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Your rating is:", numRating, " Thank you for your feedback!")
	}

}
