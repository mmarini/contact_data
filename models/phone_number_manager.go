package models

import (
	"contact_data/persistance"
)

func SavePhoneNumber(contactId int, phoneNumber PhoneNumber) int {
	sql := `INSERT into "phone_numbers"("contact_id", "phone_number") VALUES($1, $2) RETURNING id`
	savedId := persistance.Insert(sql, contactId, phoneNumber)

	return savedId
}
