package models

type Contact struct {
	FullName     string        `json:"full_name"`
	Email        string        `json:"email"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers"`
}

func (contact Contact) Format() Contact {
	newContact := Contact{FullName: contact.FullName, Email: contact.Email}

	for _, phoneNumber := range contact.PhoneNumbers {
		newPhoneNumber, _ := phoneNumber.Format()
		newContact.PhoneNumbers = append(newContact.PhoneNumbers, newPhoneNumber)
	}

	return newContact
}
