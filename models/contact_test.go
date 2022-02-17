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
