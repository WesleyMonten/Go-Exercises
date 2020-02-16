package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/WesleyMonten/Go-Exercises/shared/input"
	log "github.com/sirupsen/logrus"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("Incorrect usage")
		return
	}

	records, err := loadData(os.Args[1])
	if err != nil {
		log.Errorf("Error reading records: %v", err)
		return
	}

	g := check(5)

	fmt.Println("Countries with 5 or more of the largest cities in the world:\n")

	for _, cities := range records {
		if g(cities) {
			fmt.Printf("\t%v\n", cities[0].Country)
		}
	}
}

func check(number int) func([]City) bool {
	return func(cities []City) bool {
		return len(cities) >= number
	}
}
func loadData(name string) (records Records, err error) {

	var reader *input.FileReader

	if reader, err = input.NewFileReader(name); err == nil {
		reader.ReadLine()

		records = make(Records)

		var record string
		for err == nil {
			record, err = reader.ReadLine()
			if io.EOF == err {

				return records, nil
			}
			if nil != err {
				break
			}
			city, _ := ExtractCity(strings.Trim(record, "\n"))

			records[city.Country] = append(records[city.Country], city)
		}
	}
	return
}
