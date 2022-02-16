package models

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type PhoneNumber string

func (phoneNumber PhoneNumber) Format() (PhoneNumber, error) {

	// check if we have invalid characters
	_, err := validateValidCharacters(string(phoneNumber))
	if err != nil {
		validError := fmt.Errorf("PhoneNumber.Valid: %e", err)
		return phoneNumber, validError
	}

	basePhoneNumber := stripInvalidCharacters(string(phoneNumber))

	// check country code validation
	if !validateCountryCode(basePhoneNumber) {
		return phoneNumber, errors.New("PhoneNumber.Valid: Country Code must be Australia if supplied")
	}

	basePhoneNumber = formatCountryCode(basePhoneNumber)

	// check format
	if !validateFormat(basePhoneNumber) {
		return phoneNumber, errors.New("PhoneNumber.Valid: Phone Number is in an invalid format")
	}

	return PhoneNumber(basePhoneNumber), nil
}

func validateValidCharacters(phoneNumber string) (bool, error) {
	// remove leading and trailing whitespce
	formattedNumber := strings.TrimSpace(phoneNumber)

	// only want numbers, open and close brackets and `-`
	r, _ := regexp.Compile(`\d|\(|\)|-|\s`)

	for index, char := range formattedNumber {
		testString := string(char)

		// if the first character is a +, accept it
		if index == 0 && !(testString == "+" || r.MatchString(testString)) {
			err := fmt.Errorf("PhoneNumber.validateValidCharacters, invalid character detected, %s at position %d", testString, index)
			return false, err
		}

		// any other characters, only want a number
		if index > 0 && !r.MatchString(testString) {
			err := fmt.Errorf("PhoneNumber.validateValidCharacters, invalid character detected, %s at position %d", testString, index)
			return false, err
		}
	}

	return true, nil
}

func stripInvalidCharacters(phoneNumber string) string {

	var outputNumber []string

	// remove leading and trailing whitespce
	formattedNumber := strings.TrimSpace(phoneNumber)

	// only want numbers
	r, _ := regexp.Compile(`\d`)

	for index, char := range formattedNumber {
		testString := string(char)

		// if the first character is a +, accept it
		if index == 0 && (testString == "+" || r.MatchString(testString)) {
			outputNumber = append(outputNumber, testString)
		}

		// any other characters, only want a number
		if index > 0 && r.MatchString(testString) {
			outputNumber = append(outputNumber, testString)
		}
	}

	return strings.Join(outputNumber, "")
}

func validateFormat(phoneNumber string) bool {
	match, _ := regexp.MatchString(`^\+[1-9]\d{1,14}$`, phoneNumber)
	return match
}

func validateCountryCode(phoneNumber string) bool {
	if strings.HasPrefix(phoneNumber, `+`) {
		return strings.HasPrefix(phoneNumber, `+61`)
	}
	return true
}

func formatCountryCode(phoneNumber string) string {
	// If the string starts with a 0, remove it and replace with +61
	if strings.HasPrefix(phoneNumber, `0`) {
		phoneNumber = strings.Replace(phoneNumber, `0`, `+61`, 1)
	}

	if !strings.HasPrefix(phoneNumber, `+61`) {
		phoneNumber = "+61" + phoneNumber
	}

	return phoneNumber
}
