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

//function to print all students in  a section
func (sec Section) String() string {
	//declare an empty string
	var concatanatedNames string

	//for each student in the section add it and a newline
	for _, j := range sec.mStudentNames {
		concatanatedNames = concatanatedNames + "\n" + j
	}

	//return the created string to the caller
	return concatanatedNames
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
	return student.mStudentName + " is already in " + student.mSection.mDepartment + " " + student.mSection.mCourseNumber + "-" + student.mSection.mSectionNumber
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
	//print out a header
	fmt.Println("Simulates basics of class registration")

	//create a section
	//department cs, course number 101, section number 1
	//give it an empty string to start filling students names into
	aSection := Section{"cs", "101", "1", []string{}}

	//create two strings, one for yes/no, one for student's name
	var studentName, doContinue string
	for {
		//print prompt containing section details
		fmt.Print("Enter name to enroll in ")
		fmt.Print(aSection.mDepartment)
		fmt.Print(" ")
		fmt.Print(aSection.mCourseNumber)
		fmt.Print("-")
		fmt.Print(aSection.mSectionNumber)
		fmt.Print(": ")

		//grab a student's name
		fmt.Scan(&studentName)

		//attempt to register that student in the section
		registerSuccess, registerError := aSection.register(studentName)
		//if the register function doesn't return true, print the error
		if !registerSuccess {
			fmt.Println(registerError)
		}

		//ask the user whether to continue or not
		fmt.Print("Do you want to quit? (yes/no): ")
		fmt.Scan(&doContinue)
		if doContinue == "yes" {
			//anything not explicitily "yes" continues entering students
			break
		}
	}

	//dump the final list of students
	fmt.Print("The final list of students is:")
	fmt.Println(aSection)
}
