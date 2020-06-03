package handlers

import "fmt"

// Event : Might need to export
type Event struct {
	Id    uint16 `json :"id"`
	Title string `json :"title"`
	Prize uint16 `json :"prize"`
	Head  string `json :"head"`
	Phone string `json :"phone"`
}

// Events : Might need to export
type Events struct {
	Events []Event `json:"events"`
}

// CheckError : Common function for the package
func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
