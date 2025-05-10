package basics

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	// PascaCase: Structs, interfaces, enums.
	// UserInfo, HTTPRequest

	// Scake case: Conmopnly used for naming variables, constants and file names. Multiple words or phrases
	//  iser_id, first_name, http_requests

	// UPEPRCASE
	// For constants

	// mixedCase
	// javaScript, htmlDocument, userInfo, isValid
	// Identifiers of external variables, pr normal variables

	// Packages names always in lowercase

	const MASRETRIES = 5
	var employeeId = 1001
	fmt.Println("EmployeeID: ", employeeId)

}
