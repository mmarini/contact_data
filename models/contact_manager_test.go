package models

import (
	"testing"
)

func TestSaveContact(t *testing.T) {

	phoneNumbers := []PhoneNumber{"12345678", "98765432"}
	contact := Contact{FullName: "Michael Marini", Email: "mmarini@a.com", PhoneNumbers: phoneNumbers}

	result := SaveContact(contact)

	if result <= 0 {
		t.Fatalf("Expected Id to be returned")
	}
}
