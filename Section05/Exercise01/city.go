package main

import (
	"errors"
	"strconv"
	"strings"
)

type (
	City struct {
		Name       string
		Country    string
		Population uint
	}

	Records map[string][]City
)

func ExtractCity(s string) (c City, err error) {
	const numberOfFields = 3
	fields := strings.Split(s, ",")

	if numberOfFields > len(fields) {
		return c, errors.New("Fields are incorrect")
	}

	c.Name = fields[0]
	c.Country = fields[1]
	population, _ := strconv.ParseUint(fields[2], 10, 64)
	c.Population = uint(population)

	return
}
