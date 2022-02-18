package models

import (
	"contact_data/persistance"
)

func SavePhoneNumber(contactId int, phoneNumber PhoneNumber) int {

	var columns = []string{"contact_id", "phone_number"}
	var tableName = string("phone_numbers")

	savedId := persistance.Insert(tableName, columns, contactId, phoneNumber)

	return savedId
}
