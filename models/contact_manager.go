package models

import (
	"contact_data/persistance"
)

func SaveContact(contact Contact) int {

	var columns = []string{"full_name", "email"}
	var tableName = string("contacts")

	saveableContact := contact.Format()

	contactId := persistance.Insert(tableName, columns, saveableContact.FullName, saveableContact.Email)

	for _, phoneNumber := range saveableContact.PhoneNumbers {
		SavePhoneNumber(contactId, phoneNumber)
	}

	return contactId
}
