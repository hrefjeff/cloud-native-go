package main

import "fmt"

// --- 1. Interface (Polymorphism) ---
// The Person interface defines a behavior called `DisplayDetails`.
// Any type that implements this method is considered a `Person`.
type Person interface {
	DisplayDetails() string
}

// --- 2. Base Struct (Object) ---
// The Employee struct acts like a class, holding data for an employee.
type Employee struct {
	Firstname string
	Lastname  string
}

// --- 3. Method for the Employee Struct ---
// FullName is a method associated with the Employee struct.
// It acts on an instance of an Employee.
func (e Employee) FullName() string {
	return e.Firstname + " " + e.Lastname
}

// Employee implements the Person interface by defining this method.
func (e Employee) DisplayDetails() string {
	return "Employee: " + e.FullName()
}

// --- 4. "Inheritance" via Embedding ---
// The Manager struct embeds the Employee struct. This means a Manager
// "is an" Employee and automatically gets all its fields and methods.
type Manager struct {
	Employee   // Embedding Employee struct
	Department string
}

// Manager also implements the Person interface.
// This is an example of method overriding.
func (m Manager) DisplayDetails() string {
	return fmt.Sprintf("Manager: %s, Department: %s", m.FullName(), m.Department)
}

// --- 5. Main function to run the program ---
func main() {
	fmt.Println("--- Demonstrating Object-Oriented Concepts in Go ---")

	// Create an instance of Employee
	emp := Employee{
		Firstname: "Alice",
		Lastname:  "Smith",
	}

	// Create an instance of Manager
	mgr := Manager{
		Employee: Employee{
			Firstname: "Bob",
			Lastname:  "Johnson",
		},
		Department: "Engineering",
	}

	// Calling methods directly on the objects
	fmt.Println("\nDirectly calling methods:")
	fmt.Println(emp.DisplayDetails()) // Calls Employee's method
	fmt.Println(mgr.DisplayDetails()) // Calls Manager's method

	// Because Manager embeds Employee, it can directly access Employee's methods
	fmt.Printf("Manager's full name is: %s\n", mgr.FullName())

	// --- Polymorphism in action ---
	// We create a slice of type `Person` interface.
	// It can hold any type that implements the interface (Employee and Manager).
	people := []Person{emp, mgr}

	fmt.Println("\nDemonstrating polymorphism with an interface slice:")
	// Loop through the slice and call the method. Go automatically
	// calls the correct implementation for each type.
	for _, p := range people {
		fmt.Println(p.DisplayDetails())
	}
}
