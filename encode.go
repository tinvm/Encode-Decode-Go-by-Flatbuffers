package main

import (
	"io/ioutil"

	model "./model"
	flatbuffers "github.com/google/flatbuffers/go"
)

func main() {
	builder := flatbuffers.NewBuilder(0)

	phoneType := builder.CreateString("phone type 1")
	number := builder.CreateString("number 1")

	model.PhoneStart(builder)
	model.PhoneAddPhoneType(builder, phoneType)
	model.PhoneAddNumber(builder, number)
	phone := model.PhoneEnd(builder)

	model.ContactStartPhonesVector(builder, 1)
	builder.PrependUOffsetT(phone)
	phones := builder.EndVector(1)

	id := builder.CreateString("contact id")
	firstName := builder.CreateString("first name")
	lastName := builder.CreateString("last name")
	description := builder.CreateString("description")

	model.ContactStart(builder)
	model.ContactAddId(builder, id)
	model.ContactAddFirstName(builder, firstName)
	model.ContactAddLastName(builder, lastName)
	model.ContactAddDescription(builder, description)
	model.ContactAddPhones(builder, phones)
	contact := model.ContactEnd(builder)

	model.MessageStartContactsVector(builder, 1)
	builder.PrependUOffsetT(contact)
	contacts := builder.EndVector(1)

	messageId := builder.CreateString("message id")

	receiver := builder.CreateString("receiver 1")

	model.MessageStartReceiversVector(builder, 1)
	builder.PrependUOffsetT(receiver)
	receivers := builder.EndVector(1)

	model.MessageStart(builder)
	model.MessageAddId(builder, messageId)
	model.MessageAddContacts(builder, contacts)
	model.MessageAddReceivers(builder, receivers)
	message := model.MessageEnd(builder)

	builder.Finish(message)
}
