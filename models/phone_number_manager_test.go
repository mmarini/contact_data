package models

import (
	"testing"
)

func TestSavePhoneNumber(t *testing.T) {
	result := SavePhoneNumber(1, PhoneNumber("123456"))

	if result <= 0 {
		t.Fatalf("Expected Id to be returned")
	}
}
