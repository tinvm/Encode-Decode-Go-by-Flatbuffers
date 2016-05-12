package contact;

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
