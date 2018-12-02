package usecase

import (
	"inversion/entity"
)

// printInterface - access infrastructure methods
type printInterface interface {
	Print(s string)
}

// PersonUse - give access to usecase methods
type PersonUse struct {
	printer printInterface
}

// NewPersonUse - inject interface methods and return usecase methods
func NewPersonUse(pi printInterface) PersonUse {
	return PersonUse{pi}
}

// PrintPersonName - print name of person
func (u PersonUse) PrintPersonName(p entity.Person) {
	u.printer.Print(p.Name())
}
