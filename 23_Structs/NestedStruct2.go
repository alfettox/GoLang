/*
Author: Giovanni De Franceschi
*/



package main

import "fmt"

type Student struct {
    FirstName string
    LastName  string
    Age       int
}

type Course struct {
    Title       string
    Instructor  string
    EnrolledStudents []Student
}

func main() {
    student1 := Student{
        FirstName: "John",
        LastName:  "Doe",
        Age:       20,
    }

    student2 := Student{
        FirstName: "Jane",
        LastName:  "Smith",
        Age:       22,
    }

    course1 := Course{
        Title:       "Introduction to Programming",
        Instructor:  "Professor Smith",
        EnrolledStudents: []Student{student1, student2}, //slice of structs
    }

    course2 := Course{
        Title:       "Data Structures and Algorithms",
        Instructor:  "Professor Johnson",
        EnrolledStudents: []Student{student1},
    }

    fmt.Println("Course:", course1.Title)
    fmt.Println("Instructor:", course1.Instructor)
    fmt.Println("Enrolled Students:")
    for _, student := range course1.EnrolledStudents {
        fmt.Printf("- %s %s\n", student.FirstName, student.LastName)
    }

    fmt.Println("\nCourse:", course2.Title)
    fmt.Println("Instructor:", course2.Instructor)
    fmt.Println("Enrolled Students:")
    for _, student := range course2.EnrolledStudents {
        fmt.Printf("- %s %s\n", student.FirstName, student.LastName)
    }
}
