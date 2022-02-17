package models

type Contact struct {
	FullName     string        `json:"full_name"`
	Email        string        `json:"email"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers"`
}
