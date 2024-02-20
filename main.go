package main

import "fmt"

func main() {
	// ====================Example usage of single responsblity principle===============

	// Create a FileReader for a specific file
	fileReader := NewFileReader(`C:\Users\Anurag singh\Downloads\nested_json.txt`)

	// Read content from the file
	content, err := fileReader.ReadContent()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Create a ConsolePrinter
	consolePrinter := ConsolePrinter{}

	// Print the content to the console
	consolePrinter.PrintContent(content)

	// =============================Example usage of OPEN/CLOSE principle=================================

	// Create instances of Rectangle and Circle
	rectangle := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 2}

	// Calculate the area of individual shapes
	fmt.Printf("Rectangle Area: %.2f\n", rectangle.Area())
	fmt.Printf("Circle Area: %.2f\n", circle.Area())

	// Create a slice of shapes
	shapes := []Shape{rectangle, circle}

	// Calculate the total area without modifying existing shapes or the CalculateTotalArea function
	totalArea := CalculateTotalArea(shapes)
	fmt.Printf("Total Area: %.2f\n", totalArea)

	//==============================Liscov Substitution principle==============================================
	// Create instances of Sparrow and Ostrich
	sparrow := Sparrow{Bird{Name: "Sparrow"}}
	ostrich := Ostrich{Bird{Name: "Ostrich"}}

	// Fly function can be called on Sparrow, adhering to the Flyer interface
	fmt.Println(sparrow.Fly())

	// Attempting to call Fly on Ostrich, which cannot fly
	// This adheres to Liskov Substitution Principle as calling Fly on Ostrich doesn't break the program
	fmt.Println(ostrich.Fly()) // Output: "Ostrich cannot fly"

	//=======================================Interface segregation principle================================

	// Create instances of Robot and Human
	robot := Robot{}
	human := Human{}

	// Both Robot and Human can work
	performWork(robot)
	performWork(human)

	// Only Human can eat
	performEat(human)

	//====================================Dependency Inversion principle=================
	// Create instances of concrete notification services
	emailService := EmailService{}
	smsService := SMSService{}

	// Create instances of NotificationManager with different services injected
	emailNotifier := NotificationManager{Service: emailService}
	smsNotifier := NotificationManager{Service: smsService}

	// Use the notification managers to send notifications
	emailNotifier.Notify("Hello via Email")
	smsNotifier.Notify("Hello via SMS")
}
