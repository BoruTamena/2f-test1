package main

import "fmt"

type programming struct {
	title     string
	salary    float64
	dephead   string
	employees []string
}
type design struct {
	title     string
	salary    float64
	dephead   string
	employees []string
}

type Department interface {
	EmployeesList() []string
	Title() string
	TotalEmployeesCount() int
	HeadOfDepartment() string
	NetSalary() float64
}

func (p programming) EmployeesList() []string {

	if len(p.employees) == 0 {

		fmt.Println(" no employe yet")
		return p.employees
	}

	return p.employees
}

func (d design) EmployeesList() []string {
	if len(d.employees) == 0 {

		fmt.Println(" no employe yet")
		return d.employees
	}

	return d.employees
}

func (p programming) Title() string {

	if p.title == "" {
		fmt.Println(" no title yet ")
		return ""
	}

	return p.title
}

func (d design) Title() string {

	if d.title == "" {
		fmt.Println(" no title yet ")
		return ""
	}

	return d.title
}

func (p programming) TotalEmployeesCount() int {

	return len(p.employees)

}

func (d design) TotalEmployeesCount() int {

	return len(d.employees)

}

func (p programming) HeadOfDepartment() string {
	if p.dephead == "" {
		return ""
	}
	return p.dephead

}

func (d design) HeadOfDepartment() string {

	if d.dephead == "" {
		return ""
	}
	return d.dephead

}

func (p programming) NetSalary() float64 {

	return p.salary * 0.5

}

func (d design) NetSalary() float64 {
	return d.salary * 0.5

}

func ShowDepartmentDetails(dep Department) {
	fmt.Println("Dispalying Department detail")
	fmt.Printf("Dep:%v\n", dep.Title())
	fmt.Printf("Dep-head:%v\n", dep.HeadOfDepartment())
	fmt.Printf("Emp_Count:%v\n", dep.TotalEmployeesCount())
	fmt.Printf("Emp_List:%v\n", dep.EmployeesList())
	fmt.Printf("Emp_Salary:%v\n", dep.NetSalary())
}

func main() {

	// init the programming struct
	p1 := programming{
		title:     "python",
		salary:    500000,
		dephead:   "Boru",
		employees: []string{"Shime", "Adil"},
	}

	d1 := design{
		title:     "UI/UX",
		salary:    300000,
		dephead:   "Malese",
		employees: []string{"Heran", "Almaz"},
	}

	// Displaying the department detail
	ShowDepartmentDetails(p1)
	ShowDepartmentDetails((d1))

}
