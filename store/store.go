package store

import (
	"errors"
	"fmt"
)

type Account struct {
	FirstName string
	LastName  string
}

type Employee struct {
	Account
	Credits float64
}

func (a *Account) ChangeName(newname string) {
	a.FirstName = newname
}

func (e Employee) String() string {
	return fmt.Sprintf("Name: %s %s\nCredits: %.2f\n", e.FirstName, e.LastName, e.Credits)
}

func CreateEmployee(firstName, lastName string, credits float64) (*Employee, error) {
	return &Employee{Account{firstName, lastName}, credits}, nil
}

func (e *Employee) AddCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		e.Credits += amount
		return e.Credits, nil
	}
	return 0.0, errors.New("INVALID CREDIT AMOUNT")
}

func (e *Employee) RemoveCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		if amount <= e.Credits {
			e.Credits -= amount
			return e.Credits, nil
		}
		return 0.0, errors.New("YOU CAN'T REMOVE MORE CREDITS THAN THE ACCOUNT HAS")
	}
	return 0.0, errors.New("YOU CAN'T REMOVE NEGATIVE NUMBERS")
}

func (e *Employee) CheckCredits() float64 {
	return e.Credits
}
