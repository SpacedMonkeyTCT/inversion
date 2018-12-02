package infrastructure

import (
	"fmt"
)

// ConsoleInfra - access infrastructure methods
type ConsoleInfra struct {
}

// NewConsoleInfra - return infra methods
func NewConsoleInfra() ConsoleInfra {
	return ConsoleInfra{}
}

// Print person
func (ci ConsoleInfra) Print(s string) {
	fmt.Println(s)
}
