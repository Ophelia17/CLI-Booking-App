package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
	// "strconv"
)

var conferenceName string = "GoConference"

const conferenceTickets = 100

var remainingTickets int = conferenceTickets

// var bookings = [50]string{}
// var bookings []string{}
// bookings = make([]string, 50)
// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberofTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	for {
		var userFirstName string
		var userLastName string
		var userEmail string

		chk := helper.GetUserInput(&userFirstName, &userLastName, &userEmail)

		if !chk {
			continue
		}

		var userTickets int

		chk = bookTicket(&userTickets, &remainingTickets, userFirstName, userLastName, userEmail)
		if !chk {
			continue
		}

		wg.Add(1)
		go sendTicket(userTickets, userFirstName, userLastName, userEmail)

		firstName := getFirstNames()
		fmt.Printf("The first names of the bookings are : %v\n", firstName)

		// fmt.Printf("The first booking: %v\n", bookings[0])
		// fmt.Printf("The booking type is: %T\n", bookings)
		// fmt.Printf("The number of bookings: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive the booking confirmation at %v.\n", userFirstName, userLastName, userTickets, userEmail)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		if remainingTickets == 0 {
			fmt.Printf("Sorry, we are sold out! Please come back next year")
			break
		}

	}

	wg.Wait()

}

func greetUser() {
	fmt.Printf("Welcome to %v Booking Application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

func getFirstNames() []string {
	firstName := []string{}
	for _, booking := range bookings {
		// firstName = append(firstName, booking["firstName"])
		firstName = append(firstName, booking.firstName)

	}
	return firstName
}

func bookTicket(userTickets *int, remainingTickets *int, userFirstName string, userLastName string, userEmail string) bool {
	fmt.Println("How many tickets would you like to book?")
	fmt.Scan(userTickets)

	if *userTickets > *remainingTickets {
		fmt.Printf("Sorry, we only have %v tickets remaining.\n", *remainingTickets)
		return false
	}

	*remainingTickets = *remainingTickets - *userTickets

	// create user map
	// var user = make(map[string]string)
	// user["firstName"] = userFirstName
	// user["lastName"] = userLastName
	// user["email"] = userEmail
	// user["numberOfTickets"] = strconv.FormatUint(uint64(*userTickets), 10)

	var userData = UserData{
		firstName:       userFirstName,
		lastName:        userLastName,
		email:           userEmail,
		numberofTickets: uint(*userTickets),
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	// bookings = append(bookings, userFirstName+" "+userLastName)

	return true
}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nTo email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
