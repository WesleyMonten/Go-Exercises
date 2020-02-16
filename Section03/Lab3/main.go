package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/WesleyMonten/Go-Exercises/shared/input"
	"github.com/sirupsen/logrus"
)

type (
	Student []int
	Class   []Student
)

const (
	courseID           = iota
	amountOfExamsID    = iota
	amountOfStudentsID = iota
	maxID              = iota
	studentID          = iota
)

var amountExams, amountStudents, max int
var class string

func main() {
	if len(os.Args) != 5 {

		logrus.Error("Not the right amount")

		return
	}

	var amount int64
	var err error

	if amount, err = strconv.ParseInt(os.Args[amountOfExamsID], 10, 8); amount <= 0 || amount > 20 || nil != err {

		logrus.Fatal("Invalid amount: %v", amount)

		return
	}
	amountExams = int(amount)

	if amount, err = strconv.ParseInt(os.Args[amountOfStudentsID], 10, 8); amount <= 0 || nil != err {

		logrus.Fatal("Invalid amount: %v", amount)

		return
	}

	amountStudents = int(amount)

	if amount, err = strconv.ParseInt(os.Args[maxID], 10, 8); amount < 100 || amount > 120 || nil != err {

		logrus.Fatal("Invalid amount: %v", amount)

		return
	}
	max = int(amount)

	className := os.Args[courseID]

	if 0 == len(className) {

		logrus.Fatal("No classname")

		return
	}

	classGrades := initClassGrades(amountStudents, amountExams, max)
	printClassGrades(className, classGrades)
}
func printClassGrades(name string, class Class) {

	// print report header
	fmt.Printf("Classname: %v\n", strings.Title(name))
	fmt.Printf("Students: \n\n")

	fmt.Print("StudentID: ")

	ex := "Exam %v"

	for i := 0; i < len(class[0]); i++ {

		s := fmt.Sprintf(ex, i+1)
		fmt.Printf("%8v", s)

	}
	fmt.Printf("%14v\n", "Letter Grade")

	for i, s := range class {

		fmt.Printf("%10v", i+1)

		for _, g := range s {
			fmt.Printf("%8v", g)
		}

		fmt.Printf("%-v\n", getGrade(getAvgScore(s)))

	}

	var scores []int
	fmt.Printf("%-10v", "Average: ")

	for i := 0; i < amountExams; i++ {

		for j := 0; j < amountStudents; j++ {

			scores = append(scores, class[j][i])
		}

		fmt.Printf("%8v", getAvgScore(scores))
	}

}

func initClassGrades(sA, eA, max int) (class Class) {
	class = make(Class, eA)
	for i := range class {

		class[i] = make(Student, eA)

		for j := range class[i] {

			class[i][j] = input.GetRandInt(20, max)
		}

	}

	return class
}

func getAvgScore(s Student) int {
	var total int

	for _, v := range s {
		total += v
	}
	return total / len(s)
}

func getGrade(avg int) string {

	grades := []string{"A+", "A", "B+", "B", "C+", "C", "D+", "D"}
	scores := []int{98, 95, 90, 80, 70, 60, 55, 45}

	for i, s := range scores {
		if avg >= s {
			return grades[i]
		}
	}
	return "F"
}
