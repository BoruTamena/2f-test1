# Go Programming Challenges

This repository contains solutions to various Go programming challenges. Each challenge focuses on different aspects of Go programming, including concurrency, data structures, and interface implementation.

## Stack Implementation

### Problem:

Implement a stack in Go with the following methods: Push, Pop, Empty.

### Solution:

Explore the provided Go implementation in `stack.go`. This program showcases a basic stack with essential methods such as Push, Pop, and Empty.

## Phone Number Formatting Function

### Problem:

Create the "FormatPhone" function in Golang.

### Solution:

Explore the provided Go implementation in `formatph.go`. This program showcases a basic stack with essential methods such as Push, Pop, and Empty.

1. **FormatPhone Function:**
    a. Write a function called FormatPhone in Golang.
    b. It formats an Ethiopian phone number.
    c. The function throws an error "invalid phone" if the input is not a valid Ethiopian tel.
    d. It accepts the phone number as a string in any of the formats: [251901040506, 0901040506, +251901040506, 901040506].
    e. The output always returns in this format: 251901040506 if the phone number is valid.

2. **Usage Example:**
    ```go
    package main
    
    import (
    	"fmt"
    	"log"
    )
    
    func main() {
    	phoneNumber := "251901040506"
    
    	formattedNumber, err := FormatPhone(phoneNumber)
    	if err != nil {
    		log.Fatal(err)
    	}
    
    	fmt.Println("Formatted Phone Number:", formattedNumber)
    }
    ```

## Company Department Details

### Problem:

Implement a code that allows you to get the details of a company department.

1. **Department Interface:**
    a. Create two struct objects, "programming" and "design".
    b. Both should implement the Department interface which includes methods like
    EmployeesList() []string, Title() string, TotalEmployeesCount()
    int, HeadOfDepartment() string, NetSalary() float64.
    c. Assume all employees in their respective departments are paid the same
    amount.
    d. NetSalary for the design team is calculated as salary * 0.75, and for the
    programming department, it's calculated as salary * 0.5.

2. **Usage Example:**
    ```go
    package main
    
    import "fmt"
    
    func main() {
    	// Create department instances
    	programming := NewProgrammingDepartment()
    	design := NewDesignDepartment()
    
    	// Display department details
    	displayDepartmentDetails(programming)
    	displayDepartmentDetails(design)
    }
    
    func displayDepartmentDetails(department Department) {
    	fmt.Println("Department Title:", department.Title())
    	fmt.Println("Employees List:", department.EmployeesList())
    	fmt.Println("Total Employees Count:", department.TotalEmployeesCount())
    	fmt.Println("Head of Department:", department.HeadOfDepartment())
    	fmt.Println("Net Salary:", department.NetSalary())
    	fmt.Println()
    }
    ```

### Solution

Explore the provided Go implementation in `stack.go`. This program showcases a basic stack with essential methods such as Push, Pop, and Empty.

Feel free to explore the source code for each task and adapt them to your specific needs or learning objectives. If you have any questions or suggestions, please feel free to open an issue or contribute to the repository.
