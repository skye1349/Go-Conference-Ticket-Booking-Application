# Go Conference Ticket Booking Application

This is a simple console-based ticket booking application for a conference written in Go. 

## Packages Used

The program makes use of the following Go standard library packages:

- `fmt`: For console I/O operations, like printing to the console and scanning user input.

- `sync`: Specifically, the `WaitGroup` from the `sync` package is used for managing the ticket sending goroutines. This ensures that the main function waits for all tickets to be sent before terminating the program.

- `time`: The `time` package is used to simulate the delay in sending the ticket.

## Functionality

The main functionality of the application is to book tickets for a conference. It achieves this through a series of operations:

1. **Greeting the User**: The application starts by greeting the user and informing them about the total number of tickets and how many are remaining.

2. **Getting User Input**: The application then asks the user to enter their first name, last name, email address, and the number of tickets they wish to book.

3. **Validating User Input**: The application checks if the input entered by the user is valid. For example, it checks if the name is not too short, if the email address contains the "@" symbol, and if the number of tickets entered is valid.

4. **Booking the Ticket**: If the user's input is valid, the application books the tickets and deducts the booked tickets from the remaining total. It also thanks the user and sends a confirmation message to the console.

5. **Sending the Ticket**: A goroutine is started to simulate sending the ticket to the user's email address. This operation is delayed by 50 seconds to simulate the time it takes to send an email.

Please note that this is a simple application and doesn't actually send an email. The "email sending" is simulated and the program only displays a message on the console.

## Data Structures

The application uses an array of a custom `UserData` type to keep track of the bookings. Each booking is stored with the user's first name, last name, email, and the number of booked tickets. The remaining tickets are tracked using a simple unsigned integer.

## Concurrency

The program makes use of Go's goroutines to simulate the delay in sending tickets. This helps in making the ticket sending operation non-blocking.
