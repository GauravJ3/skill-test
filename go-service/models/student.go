package models

type Student struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Class    string `json:"class_name"`
	RollNo   string `json:"roll_no"`
}
