package day5

import "fmt"

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	if initialAge < 0 {
		p.age = 0
		fmt.Println("Age is not valid, setting age to 0.")
	}
	p.age = initialAge
	return p
}

func (p person) amIOld() {
	if p.age < 13 {
		fmt.Println("You are young.")
	} else if p.age >= 13 && p.age < 18 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are old.")
	}

}

func (p person) yearPasses() person {
	p.age += 1
	return p
}
