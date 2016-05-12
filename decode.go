package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	model "./model"
)

type Phone struct {
	PhoneType string
	Number    string
}

type Contact struct {
	Id          string
	FirstName   string
	LastName    string
	Description string
	Phones      []Phone
}

type Message struct {
	Id        string
	Contacts  []Contact
	Receivers []string
}

func main() {
	buf, err := ioutil.ReadFile("types.go")
	if err != nil {
		panic(err)
	}

	message := model.GetRootAsMessage(buf, 0)

	m := Message{
		Id:        string(message.Id()),
		Contacts:  []Contact{},
		Receivers: []string{},
	}

	contact := new(model.Contact)

	for i := 0; i < message.ContactsLength(); i++ {
		if message.Contacts(contact, i) {
			c := Contact{
				Id:          string(contact.Id()),
				FirstName:   string(contact.FirstName()),
				LastName:    string(contact.LastName()),
				Description: string(contact.Description()),
				Phones:      []Phone{},
			}

			phone := new(model.Phone)

			for j := 0; j < contact.PhonesLength(); j++ {
				if contact.Phones(phone, j) {
					p := Phone{
						PhoneType: string(phone.PhoneType()),
						Number:    string(phone.Number()),
					}

					c.Phones = append(c.Phones, p)
				}
			}

			m.Contacts = append(m.Contacts, c)
		}
	}

	for i := 0; i < message.ReceiversLength(); i++ {
		m.Receivers = append(m.Receivers, string(message.Receivers(i)))
	}

	output, _ := json.Marshal(m)
	fmt.Println(string(output))
}
