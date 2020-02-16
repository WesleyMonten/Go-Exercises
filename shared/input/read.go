package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadBool() bool {
	input := getInput()
	b, err := strconv.ParseBool(input)

	for nil != err {
		fmt.Println(`Sorry, expected a boolean value.
			For example: 1, t, true, True, 0, f, false, False
			Please try again.`)
		input := getInput()
		b, err = strconv.ParseBool(input)
	}

	return b
}

func ReadFloat64() float64 {
	input := getInput()
	f, err := strconv.ParseFloat(input, 64)

	for nil != err {
		fmt.Println(`Sorry, expected an integer value. 
			For example: 3.14 or 1e6. 
			Please try again.`)
		input := getInput()
		f, err = strconv.ParseFloat(input, 64)
	}

	return float64(f)
}

func ReadInt() int {
	input := getInput()
	i, err := strconv.ParseInt(input, 10, 32)

	for nil != err {
		fmt.Println("Sorry, expected an integer value. For example, 5. Please try again.")
		input := getInput()
		i, err = strconv.ParseInt(input, 10, 32)
	}

	return int(i)
}

func ReadString() string {
	input := getInput()

	return strings.TrimSpace(input)
}
func getInput() (input string) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	for nil != err {
		fmt.Println("Sorry, unable to read entry. Please try again.")
		input, err = reader.ReadString('\n')
	}

	return
}
