package main

import (
	"booking-app/helper"
	"fmt"
)

	var conferenceName = "Go Conference"
	const conferenceTickets int= 50
	var remainingTickets uint= 50
	var bookings = make([]UserData, 0)

	type UserData struct{
		firstName string
		lastName string
		email string
		numberOfTickets uint
	}

func main() {
	

	greetUsers()
	
	
for{
	
	firstName, lastName, email, userTickets:=getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
	
	if isValidEmail && isValidName && isValidTicketNumber{
		
		bookTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("thes first names of bookings are: %v\n",firstNames)

		if remainingTickets==0{
			fmt.Println("Our conference is booked out. Come back next year")
			break
		}
	}else{
		if !isValidName{
			fmt.Println("first name or last name is too short")
		}
		if !isValidEmail{
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber{
			fmt.Println("number of tickets you entered is invalid")
		}
		
	}
	

}

	
}

func greetUsers(){
	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n",conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string{
	firstNames := []string{}
		for _, booking := range bookings{
			
			firstNames=append(firstNames, booking.firstName)

		}
		return firstNames
}



func getUserInput() (string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)	

	fmt.Println("Enter no of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}


func bookTicket(userTickets uint,  firstName string, lastName string, email string){
		remainingTickets=remainingTickets-userTickets



		var userData = UserData{
			firstName: firstName,
			lastName: lastName,
			email: email,
			numberOfTickets: userTickets,
		}


	
		bookings = append(bookings, userData)
		fmt.Println("List of bookings is %v\n", bookings)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",firstName,lastName,userTickets,email)

		fmt.Printf("%v tickets remanining for %v\n",remainingTickets,conferenceName)
		
}