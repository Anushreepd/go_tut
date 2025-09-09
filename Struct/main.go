package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	name  string
	age   int
	Grade float64
}

// Define a method on Student
func (s Student) Display() {
	fmt.Printf("Student Name: %s\n", s.name)
	fmt.Printf("Age: %d\n", s.age)
	fmt.Printf("Grade: %.2f\n", s.Grade)
}

func (s *Student) UserInput() {
	fmt.Println("Enter student name")
	reader := bufio.NewReader(os.Stdin)
	nameInput, _ := reader.ReadString('\n')
	s.name = strings.TrimSpace(nameInput)

	fmt.Println("Enter student age")
	fmt.Scanln(&s.age)

	fmt.Println("Enter student Grade")
	fmt.Scanln(&s.Grade)
	/*
		fmt.Printf("new student name, age, Grade is %s,%d,%.2f", s.name, s.age, s.Grade) */
}

func main() {
	/* student1 := Student{}
	student1.name = "anu"
	student1.age = 25
	student1.Grade = 95

	fmt.Println("student data", student1)

	student1.Display()
	student2 := Student{}
	student2.UserInput() */

	var n int
	fmt.Println("how many students")
	fmt.Scanln(&n)

	students := make([]Student, n)
	for i := 0; i < n; i++ {
		students[i].UserInput()
	}

	// Display loop
	fmt.Println("\n--- All Student Details ---")
	for i, student := range students {
		fmt.Printf("\nStudent %d:\n", i+1)
		student.Display()
	}

}
