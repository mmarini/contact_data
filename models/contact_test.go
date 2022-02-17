package models

import (
	"encoding/json"
	"testing"
)

func TestContactMarshalling(t *testing.T) {

	jsonText := `{"full_name": "Alex Bell", "email": "alex@bell-labs.com", "phone_numbers":["03 8578 6688", "1800728069"]}`

	var newContact Contact

	err := json.Unmarshal([]byte(jsonText), &newContact)

	if err != nil {
		t.Fatalf("Issue unmarshalling, %v", err)
	}
}

func TestFormat(t *testing.T) {

	phoneNumbers := []PhoneNumber{"03 8578 6688"}
	contact := Contact{FullName: "Michael Marini", Email: "mmarini@a.com", PhoneNumbers: phoneNumbers}

	formattedContact := contact.Format()
	formattedPhoneNumber := formattedContact.PhoneNumbers[0]

	if formattedPhoneNumber != "+61385786688" {
		t.Fatalf("Contact PhoneNumber not formatted, expected +61385786688, got %s", formattedPhoneNumber)
	}
}
