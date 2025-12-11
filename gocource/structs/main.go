package main

import "fmt"

type Employee struct {
	firstname string
	lastname  string
	age       int
	salary    int
	address   Address
	Company  // Embedded struct
}

type Address struct {
	street  string
	city    string
	country string
}

 type Company struct {
	name     string
	location string
	employees []Employee
}

func main() {
	p := Employee{
		firstname: "John",
		lastname:  "Doe",
		age:       30,
		salary:    50000,
		address: Address{
			street:  "123 Main St",
			city:    "New York",
			country: "USA",
		},
		Company: Company{
			name:     "TechCorp",
			location: "Silicon Valley",
			employees: []Employee{},
		},
	}
	fmt.Println("in adress struct street address", p.address.street)
	fmt.Println("In company struct company name is", p.Company.name)
	
	fmt.Println("Employee Details:",p)
	fmt.Println("First Name:", p.firstname)
	fmt.Println("Last Name:", p.lastname)
	fmt.Println("Age:", p.age)
	fmt.Println("Salary:", p.salary)
	fullName := p.getFullName()
	fmt.Println("Full Name:", fullName)

	p.giveRaise(5000)
	fmt.Println("New Salary after raise:", p.salary)
}
func (e Employee) getFullName() string {
	return e.firstname + " " + e.lastname
}

func (e *Employee) giveRaise(amount int) {
	e.salary += amount
}	