package main

import (
	"assignment-1/models"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// students data
	students := []models.Student{
		{StudentNumber: 1, Name: "Ariq", Address: "Padang", Job: "Software Engineer", ReasonToJoinClass: "I want to learn more about Golang"},
		{StudentNumber: 2, Name: "Athallah", Address: "Padang", Job: "Freelancer", ReasonToJoinClass: "I want to be a Software Engineer"},
	}

	// get the argument
	arg := os.Args[1]
	argInt, _ := strconv.Atoi(arg)

	// check if the argument is not a number
	if argInt == 0 {
		fmt.Println("Please input a number")
		return
	}

	// check if the argument is out of range
	if argInt > len(students) {
		fmt.Println("Student not found")
		return
	}

	// find the student
	student := FindStudent(argInt, students)
	fmt.Println(student)
}

func FindStudent(studentNumber int, students []models.Student) models.Student {
	for _, student := range students {
		if student.StudentNumber == studentNumber {
			return student
		}
	}
	return models.Student{}
}
