package main

import (
	"atoss/entities"
	"fmt"
)

func main() {
	companies := entities.InitCompanies()
	/*name := "Atoss"
	city := "Timisoara"*/
	var name string
	fmt.Print("Insert company name: ")
	fmt.Scanln(&name)
	var city string
	fmt.Print("Insert employees' city: ")
	fmt.Scanln(&city)
	fmt.Println("There are:", count(companies, name, city), "employees that work at", name, ", that live in", city)
}

func count(companies []entities.Company, name string, city string) int {
	no := 0

	for _, c := range companies {
		if c.Name == name {
			for _, e := range c.Employees {
				if e.Location.GetCity() == city {
					no++
				}
			}
		}
	}

	return no
}
