package main
import "fmt"


type Department struct {
	DepName		string
	Location	string
	Project		string
	Employee1	Employee
}

type Employee struct{
	Name string
	Rank string
	No uint32
	Department string
	Vaction int8
	Email string
}

func main(){
	var employee Employee

	employee = Employee{
		"Park",
		"Manager",
		240428,
		"DEV",
		20,
		"jhintmain@gmail.com",
	}

	fmt.Println(employee) // {Park Manager 240428 DEV 20 jhintmain@gmail.com}


	employee = Employee{Name: "Eon", Email: "vamos_eon@email.com"}
	fmt.Println(employee) // {Eon  0  0 vamos_eon@email.com}

	employee = Employee{Name: "Kim"}
	fmt.Println(employee)	// {Kim  0  0 }

	var department Department

	department = Department{
		"Development",
		"8F",
		"Hello Jihea",
		Employee{
			"Hong",
			"Admin",
			210528,
			"DEV",
			25,
			"rodkfjdk@gmail.com",
		},
	}

	fmt.Println(department) // {Development 8F Hello Jihea {Hong Admin 210528 DEV 25 rodkfjdk@gmail.com}}
}