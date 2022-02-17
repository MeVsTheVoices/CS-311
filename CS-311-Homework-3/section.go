// section.go
package main

import (
	"fmt"
)

//Section type specifies details relevant to each section of a class
type Section struct {
	//strings for department, course number, section number
	mDepartment    string
	mCourseNumber  string
	mSectionNumber string
	//mStudentNames specifies a slice of students enrolled in that particular section
	mStudentNames []string
}

//Struct assosciating a student by name to a particular section
type AlreadyEnrolled struct {
	//name of the student
	mStudentName string
	//section student is currently enrolled in
	mSection *Section
}

//provide an interface for AlreadyEnrolled so that it can be reduced to a string
//and returned as an error for register
func (student AlreadyEnrolled) Error() string {
	return student.mStudentName + " is already in " + student.mSection.mCourseNumber + "-" + student.mSection.mSectionNumber
}

//try to add a student 'name' to a section 'sec'

//preconditions: student is either enrolled in section, or isn't
//postconditions: student is enrolled in section

//if student is enrolled:
//returns false and AlreadyEnrolled containing section to which student is enrolled
//if student isn't enrolled:
//returns true and nil

func (this *Section) register(name string) (bool, error) {
	//look through those names already in the section
	for _, j := range this.mStudentNames {
		//if we find that name already in the section, return as specified
		if j == name {
			return false, AlreadyEnrolled{name, this}
		}
	}
	//if the student's name wasn't in the section, add it, return as specified
	this.mStudentNames = append(this.mStudentNames, name)
	return true, nil
}

func main() {
	fmt.Println("Hello World!")
}
