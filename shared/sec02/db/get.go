package db

import "errors"

func GetNextRecord() (id uint, firstName, lastName, gender string, income Currency, err error) {
	if _db.cursor < _db.max {
		r := _db.people[_db.cursor]
		_db.cursor++
		id = r.ID
		firstName = r.FirstName
		lastName = r.LastName
		gender = r.Gender
		income = r.Income
		return
	}
	err = errors.New("no more records, call ResetCuror() for a new iteration")
	return
}
