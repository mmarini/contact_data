package models

import (
	"contact_data/persistance"
)

func SaveContact(contact Contact) int {

	saveableContact := contact.Format()

	sql := `INSERT into "contacts"("full_name", "email") VALUES($1, $2) RETURNING id`
	contactId := persistance.Insert(sql, saveableContact.FullName, saveableContact.Email)

	for _, phoneNumber := range saveableContact.PhoneNumbers {
		SavePhoneNumber(contactId, phoneNumber)
	}

	return contactId
}
