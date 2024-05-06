package main_2

import "fmt"

func main() {
	// Maps
	person := map[string]string{
		"name":   "Madhunisha",
		"rollno": "7376211SE123",
		"dept":   "ISE",
		"year":   "3",
	}
	fmt.Println("Person:", person)
	fmt.Println("Name:", person["name"])

	// Slices
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)
	numbers = append(numbers, 6)
	fmt.Println("After Append:", numbers)

	// Arrays
	var weekdays [7]string = [7]string{"Mon", "Tues", "Wed", "Thur", "Fri", "Sat", "Sun"}
	fmt.Println("Weekdays:", weekdays)
	fmt.Println("Length:", len(weekdays))

	// Structs
	type Person struct {
		Name   string
		Age    int
		Rollno string
		Dept   string
		Year   string
	}
	person1 := Person{
		Name:   "Madhunisha",
		Age:    20,
		Rollno: "7376211SE123",
		Dept:   "ISE",
		Year:   "3rd Year",
	}
	fmt.Println("Name:", person1.Name)
	fmt.Println("Age:", person1.Age)
	fmt.Println("ROllno:", person1.Rollno)
	fmt.Println("Dept:", person1.Dept)
	fmt.Println("Year:", person1.Year)
}
