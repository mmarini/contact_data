package models

import (
	"contact_data/persistance"
)

func SaveContact(contact Contact) int {
	sql := `INSERT into "contacts"("full_name", "email") VALUES($1, $2) RETURNING id`
	contactId := persistance.Insert(sql, contact.FullName, contact.Email)

	for _, phoneNumber := range contact.PhoneNumbers {
		SavePhoneNumber(contactId, phoneNumber)
	}

	return contactId
}
