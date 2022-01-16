package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("you entered: ", input)

	fmt.Print("Enter Number: ")
	number, _ := reader.ReadString('\n')
	floatNum, err := strconv.ParseFloat(strings.TrimSpace(number), 64)

	if err == nil {
		fmt.Println("You entered: ", floatNum)
	} else {
		fmt.Println("Error: ", err)
	}

}
