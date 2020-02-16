package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/WesleyMonten/Go-Exercises/shared/input"
	log "github.com/sirupsen/logrus"
)

const (
	jsonFile   = "people.json"
	minRecords = 500
	maxRecords = 1000
)

type database struct {
	people []person
	cursor int
	max    int
}

var _db database

func Load() (err error) {

	dbFile := fmt.Sprintf(".%c%s", os.PathSeparator, jsonFile)
	f, err := os.Open(dbFile)
	if nil != err {
		return errors.New("Can't open people.json. Did you copy people.json from shared/sec02/db to current directory?")
	}
	defer f.Close()

	dec := json.NewDecoder(f)

	d := []person{}
	err = dec.Decode(&d)
	if nil != err {
		return err
	}
	_db.people = d[:input.GetRandInt(minRecords, maxRecords)]
	_db.max = len(_db.people)
	_db.cursor = 0
	log.Infof("%v Records loaded", _db.max)

	return
}
