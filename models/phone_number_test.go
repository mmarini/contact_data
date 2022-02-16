package models

import (
	"testing"
)

func TestValidateValidCharacters(t *testing.T) {
	result, err := validateValidCharacters("  +61 (123) 456-789 ")

	if !result && err != nil {
		t.Fatalf("Number expected to be valid, got %e", err)
	}
}

func TestValidateValidCharactersInvalid(t *testing.T) {
	result, err := validateValidCharacters("  +61 (123) A56-789 ")

	if result && err == nil {
		t.Fatalf("Number expected to be NOT valid")
	}
}

func TestStripInvalidCharactersWithSpaces(t *testing.T) {
	result := stripInvalidCharacters(" 123 45 67 89 ")

	if result != "123456789" {
		t.Fatalf("Number expected is 123456789, got %s", result)
	}
}

func TestStripInvalidCharactersWithLeadingPlus(t *testing.T) {
	result := stripInvalidCharacters(" +123 45 67 89 ")

	if result != "+123456789" {
		t.Fatalf("Number expected is +123456789, got %s", result)
	}
}

func TestStripInvalidCharactersWithOtherCharacters(t *testing.T) {
	result := stripInvalidCharacters(" +61 (438) 020 675 ")

	if result != "+61438020675" {
		t.Fatalf("Number expected is +61438020675, got %s", result)
	}
}

func TestValidFormat(t *testing.T) {
	phoneNumber := "+61438020675"

	if !validateFormat(phoneNumber) {
		t.Fatalf("Number expected to be valid, got invalid")
	}
}

func TestValidFormat1800(t *testing.T) {
	phoneNumber := "+611800728069"

	if !validateFormat(phoneNumber) {
		t.Fatalf("Number expected to be valid, got invalid")
	}
}

func TestInValidFormat(t *testing.T) {
	phoneNumber := "ABC123"

	if validateFormat(phoneNumber) {
		t.Fatalf("Number expected to NOT be valid, got valid")
	}
}

func TestFormatCountryCode(t *testing.T) {
	result := formatCountryCode("0438020675")
	if result != "+61438020675" {
		t.Fatalf("Number expected to be +61438020675, got %s", result)
	}
}

func TestFormatCountryCode1800(t *testing.T) {
	result := formatCountryCode("1800728069")
	if result != "+611800728069" {
		t.Fatalf("Number expected to be +611800728069, got %s", result)
	}
}

func TestFormatCountryCodeWithCountryCode(t *testing.T) {
	result := formatCountryCode("+61438020675")
	if result != "+61438020675" {
		t.Fatalf("Number expected to be +61438020675, got %s", result)
	}
}

func TestValidateCountryCodeWithCountryCode(t *testing.T) {
	result := validateCountryCode("+61438020675")
	if !result {
		t.Fatal("Expected number to be valid country code, got false")
	}
}

func TestValidateCountryCodeNoCountryCode(t *testing.T) {
	result := validateCountryCode("0438020675")
	if !result {
		t.Fatalf("Expected number to be valid country code, got false")
	}
}

func TestValidateCountryCodeInvalid(t *testing.T) {
	result := validateCountryCode("+1438020675")
	if result {
		t.Fatalf("Expected number to be INVALID country code, got true")
	}
}

func TestFormatWithAreaCode(t *testing.T) {
	phoneNumber, err := PhoneNumber("03 8578 6688").Format()

	if err != nil {
		t.Fatalf("Expected phone number to be valid, got error %s", err)
	}

	if string(phoneNumber) != "+61385786688" {
		t.Fatalf("Expected phone number to be +61385786688, got %s", phoneNumber)
	}
}

func TestFormatWithAreaCodeBrackets(t *testing.T) {
	phoneNumber, err := PhoneNumber("(03) 9333 7119").Format()

	if err != nil {
		t.Fatalf("Expected phone number to be valid, got error %s", err)
	}

	if string(phoneNumber) != "+61393337119" {
		t.Fatalf("Expected phone number to be +61393337119, got %s", phoneNumber)
	}
}

func TestFormatWith1800(t *testing.T) {
	phoneNumber, err := PhoneNumber("1800728069").Format()

	if err != nil {
		t.Fatalf("Expected phone number to be valid, got error %s", err)
	}

	if string(phoneNumber) != "+611800728069" {
		t.Fatalf("Expected phone number to be +611800728069, got %s", phoneNumber)
	}
}

func TestFormatWithCountryCode(t *testing.T) {
	phoneNumber, err := PhoneNumber("+6139888998").Format()

	if err != nil {
		t.Fatalf("Expected phone number to be valid, got error %s", err)
	}

	if string(phoneNumber) != "+6139888998" {
		t.Fatalf("Expected phone number to be +6139888998, got %s", phoneNumber)
	}
}

func TestFormatWithInvalidNumber(t *testing.T) {
	_, err := PhoneNumber("ABC123").Format()

	if err == nil {
		t.Fatalf("Expected phone number to be INVALID, got valid")
	}
}
