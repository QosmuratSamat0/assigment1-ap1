package main

import (
	"assigment1/Bank"
	"assigment1/Company"
	"assigment1/Library"
	"assigment1/Shapes"
	"fmt"
)

func main() {
	fmt.Println("--------- Company ----------")

	company := Company.New()

	f := &Company.FullTimeEmployee{ID: 1, Name: "Max Yellow", Salary: 50000}
	p := &Company.PartTimeEmployee{ID: 2, Name: "Canello Saimon", HourlyRate: 20, HoursWorked: 25}

	company.AddEmployee(f)
	company.AddEmployee(p)

	fmt.Println("Employee List:")
	company.ListEmployees()
	fmt.Println()

	fmt.Println("--------- Shapes ----------")
	var shapes []Shapes.Shape

	shapes = append(shapes,
		&Shapes.Rectangle{Width: 10, Height: 10},
		&Shapes.Circle{Radius: 10},
		&Shapes.Square{A: 10},
		&Shapes.Triangle{A: 10, B: 10, C: 10},
	)

	for _, shape := range shapes {
		fmt.Printf("Area: %.2f and Perimeter: %.2f\n", shape.Area(), shape.Perimeter())
	}
	fmt.Println("------------------------")
	library := Library.New()
	bankAccount := &Bank.BankAccount{ID: 1, Owner: 12345, Balance: 1000}
	fmt.Println()
	// while operator for menu console
	for {
		fmt.Println("--------- Main ----------")
		fmt.Println("1. Library Operations")
		fmt.Println("2. Bank Account Operations")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")
		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || choice < 1 || choice > 3 {
			fmt.Println("Invalid option, please try again.")
			continue
		}

		switch choice {
		case 1:
			for {
				// Library menu
				fmt.Println("\n--------- Library ----------")
				fmt.Println("1. Add Book")
				fmt.Println("2. Borrow Book")
				fmt.Println("3. Return Book")
				fmt.Println("4. List Available Books")
				fmt.Println("5. Back to Main Menu")
				fmt.Print("Choose an option: ")

				var libraryChoice int
				_, err := fmt.Scan(&libraryChoice)
				if err != nil || libraryChoice < 1 || libraryChoice > 5 {
					fmt.Println("Invalid option, please try again.")
					continue
				}

				switch libraryChoice {
				case 1:
					// Add a book
					var title, author string

					fmt.Print("Enter book title: ")
					fmt.Scan(&title)
					fmt.Print("Enter book author: ")
					fmt.Scan(&author)
					library.AddBook(title, author)

				case 2:
					// Borrow a book
					var id string
					fmt.Print("Enter book ID to borrow: ")
					fmt.Scan(&id)
					library.BorrowBook(id)

				case 3:
					// Return a book
					var id string
					fmt.Print("Enter book ID to return: ")
					fmt.Scan(&id)
					err := library.ReturnBook(id)
					if err != nil {
						fmt.Println("Error:", err)
					}
				case 4:
					// List available books
					library.ListAvailableBooks()
				case 5:
					// Go to main menu
					break
				}

				if libraryChoice == 5 {
					break
				}
			}

		case 2:
			// Bank menu
			for {
				fmt.Println("\n--- Bank Account Menu ---")
				fmt.Println("1. Deposit Money")
				fmt.Println("2. Withdraw Money")
				fmt.Println("3. Check Balance")
				fmt.Println("4. Back to Main Menu")
				fmt.Print("Choose an option: ")

				var bankChoice int
				_, err := fmt.Scan(&bankChoice)
				if err != nil || bankChoice < 1 || bankChoice > 4 {
					fmt.Println("Invalid option, please try again.")
					continue
				}

				switch bankChoice {
				case 1:
					// Deposit money
					var amount float64
					var note string
					fmt.Print("Enter deposit amount: ")
					fmt.Scan(&amount)
					fmt.Print("Enter a note: ")
					fmt.Scan(&note)
					err := bankAccount.Deposit(amount, note)
					if err != nil {
						fmt.Println("Error:", err)
					} else {
						fmt.Println("Deposit successful!")
					}
				case 2:
					// Withdraw money
					var amount float64
					var note string
					fmt.Print("Enter withdrawal amount: ")
					fmt.Scan(&amount)
					fmt.Print("Enter a note: ")
					fmt.Scan(&note)
					err := bankAccount.WithDraw(amount, note)
					if err != nil {
						fmt.Println("Error:", err)
					} else {
						fmt.Println("Withdrawal successful!")
					}
				case 3:
					// Check balance
					fmt.Printf("Current balance: %.2f\n", bankAccount.GetBalance())
				case 4:
					// Go to main menu
					break
				}

				if bankChoice == 4 {
					break
				}
			}

		case 3:
			// Exit
			fmt.Println("Exiting...")
			return
		}
	}

}

func menu() {

}
