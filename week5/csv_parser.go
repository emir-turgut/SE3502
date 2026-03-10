package main

import {
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
	"encoding/csv"
	"errors"
}

type ValidationError struct {
	Line int
	Msg string
	Err error
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("Validation error on line %d: %s (cause %w)", v.Line, v.Msg, v.Err)
}

type Student struct {
	Name string
	Grade int
}

func parse_csv(filename string) ([]Student, error){
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Can't open file: %w", err)
	}

	var students [] Student
	reader := csv.NewReader(f)

	lineNum := 0
	for{
		record, err := reader.Read()

		if err != nil {
			return nil, fmt.Errorf("CSV parsing failed: %w", err)
		}
		lineNum++

		if len(record) < 2 {
			return nil, &ValidationError{
				Line: lineNum,
				Msg: "Missing required fields (Name, Grade)",
			}
		}
		name := strings.TrimSpace(record[0])
		gradeStr := strings.TrimSpace(record[1])

		grade, err := strconv.Atoi(gradeStr)
		
		if err != nil {
			return nil, &ValidationError{
				Line : lineNum,
				Msg : "Grade must be integer",
				Err : err,
			}
		}

		if grade < 0 || grade > 100 {
			return nil, &ValidationError{
				Line: lineNum,
				Msg: "Grade must be between 0 and 100",
			}
		}

		students = append(students, Student{Name: name, Grade: grade})
	}

	return students, nil

}