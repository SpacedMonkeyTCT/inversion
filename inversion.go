package main

import (
	"inversion/entity"
	"inversion/infrastructure"
	"inversion/usecase"
)

func main() {
	ci := infrastructure.NewConsoleInfra()
	pu := usecase.NewPersonUse(ci)
	p := entity.NewPerson("Ben")

	pu.PrintPersonName(p)
}
