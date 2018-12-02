package entity

// Person - details of an individual
type Person struct {
	name string
}

// NewPerson - create person with name
func NewPerson(name string) Person {
	return Person{name}
}

// Name - return persons name
func (p Person) Name() string {
	return p.name
}
