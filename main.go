package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	ConferenceTickets = 50
	ConferenceName    = "Go Conference"
)

var (
	RemainingTickets = uint(ConferenceTickets)
	Bookings         = make([]UserData, 0)
	wg               = sync.WaitGroup{}
)

type UserData struct {
	FirstName       string
	LastName        string
	Email           string
	NumberOfTickets uint
}

func main() {
	http.HandleFunc("/", renderForm)
	http.HandleFunc("/submit", submitForm)

	fmt.Println("Server is running on localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func renderForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	form := `
		<html>
		<body>
			<form method="POST" action="/submit">
				First name: <input name="firstName" type="text"><br>
				Last name: <input name="lastName" type="text"><br>
				Email: <input name="email" type="text"><br>
				Number of tickets: <input name="tickets" type="number"><br>
				<input type="submit">
			</form>
		</body>
		</html>
	`
	fmt.Fprint(w, form)
}

func submitForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	email := r.FormValue("email")
	ticketsStr := r.FormValue("tickets")

	tickets, err := strconv.ParseUint(ticketsStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid number of tickets", http.StatusBadRequest)
		return
	}

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := tickets > 0 && uint(tickets) <= RemainingTickets

	if isValidName && isValidEmail && isValidTicketNumber {
		RemainingTickets -= uint(tickets)

		userData := UserData{
			FirstName:       firstName,
			LastName:        lastName,
			Email:           email,
			NumberOfTickets: uint(tickets),
		}

		Bookings = append(Bookings, userData)

		wg.Add(1)
		go sendTicket(userData)

		fmt.Fprintf(w, "Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, tickets, email)
		fmt.Fprintf(w, "%v tickets remaining for %v\n", RemainingTickets, ConferenceName)
	} else {
		if !isValidName {
			fmt.Fprintln(w, "First name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Fprintln(w, "Email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Fprintln(w, "Number of tickets you entered is invalid")
		}
	}
}

func sendTicket(user UserData) {
	defer wg.Done()

	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", user.NumberOfTickets, user.FirstName, user.LastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, user.Email)
	fmt.Println("#################")
}
