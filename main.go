package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jaumecapdevila/gocap"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Gowmany")
	fmt.Println("---------------------")

	fmt.Print("-> How many different types can you choose from? ")

	types, _ := reader.ReadString('\n')
	types = strings.Replace(types, "\n", "", -1)

	fmt.Print("-> How many types can you choose? ")

	options, _ := reader.ReadString('\n')
	options = strings.Replace(options, "\n", "", -1)

	fmt.Print("-> Does the order of the items matter? (y/n) ")

	order, _ := reader.ReadString('\n')
	order = strings.Replace(order, "\n", "", -1)

	var permutation = false

	if order == "y" || order == "yes" {
		permutation = true
	}

	fmt.Print("-> Is it possible to repeat a type? (y/n) ")

	isRepeatable, _ := reader.ReadString('\n')
	isRepeatable = strings.Replace(isRepeatable, "\n", "", -1)

	var repetition = false

	if isRepeatable == "y" || isRepeatable == "yes" {
		repetition = true
	}

	n, err := strconv.Atoi(types)

	if err != nil {
		log.Fatal(err)
	}

	r, err := strconv.Atoi(options)

	if err != nil {
		log.Fatal(err)
	}

	operation := gocap.NewOperation(n, r, repetition)

	var calculator gocap.Calculator

	if permutation {
		calculator = gocap.NewPermutationCalculator()
	} else {
		calculator = gocap.NewCombinationCalculator()
	}

	result := calculator.Calculate(operation)

	fmt.Printf("Given your input, there are %d different possibilities!", result)
}
