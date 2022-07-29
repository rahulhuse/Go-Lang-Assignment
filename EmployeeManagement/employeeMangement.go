package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Employee struct {
	name    string
	empid   int64
	email   string
	address string
}

func (e *Employee) setEmployee(name string, empid int64, email string, address string) {
	e.empid = empid
	e.name = name
	e.email = email
	e.address = address

}

func main() {

	// var emp Employee
	emp := Employee{}
	var arremp []Employee

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter 1 to add employee: ")
		fmt.Println("Enter 2 to get employee: ")
		fmt.Println("Enter 3 to update employee: ")
		fmt.Println("Enter 4 to delete employee: ")
		fmt.Println("Enter 0 to exit employee: ")
		scanner.Scan()
		input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

		switch input {
		case 1:
			fmt.Println("Enter employee name: ")
			scanner.Scan()
			name := scanner.Text()
			fmt.Println("Enter employee id: ")
			scanner.Scan()
			empid, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			fmt.Println("Enter employee email: ")
			scanner.Scan()
			email := scanner.Text()
			fmt.Println("Enter employee address: ")
			scanner.Scan()
			address := scanner.Text()
			emp.setEmployee(name, empid, email, address)
			arremp = append(arremp, emp)
			fmt.Println("Employee added successfully")
		case 2:
			fmt.Println("Enter empid of employee you want to display")
			scanner.Scan()
			empid, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			for i := 0; i < len(arremp); i++ {
				if arremp[i].empid == empid {
					fmt.Println("Employee name: ", arremp[i].name)
					fmt.Println("Employee email: ", arremp[i].email)
					fmt.Println("Employee address: ", arremp[i].address)
				}
			}
		case 3:
			fmt.Println("Enter empid of employee you want to update")
			scanner.Scan()
			empid, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			for i := 0; i < len(arremp); i++ {
				if arremp[i].empid == empid {

					fmt.Println("Enter employee name: ", arremp[i].name)
					scanner.Scan()
					arremp[i].name = scanner.Text()
					fmt.Println("Enter employee email: ", arremp[i].email)
					scanner.Scan()
					arremp[i].email = scanner.Text()
					fmt.Println("Enter employee address: ", arremp[i].address)
					scanner.Scan()
					arremp[i].address = scanner.Text()
				}
			}

		case 4:
			fmt.Println("Enter emp id")
			scanner.Scan()
			empid, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			for i := len(arremp) - 1; i >= 0; i-- {
				if arremp[i].empid == empid {

					arremp = append(arremp[:i], arremp[i+1:]...)
					// arremp = arremp[empid:]
				}
			}
		default:
			os.Exit(0)

		}
	}

}
