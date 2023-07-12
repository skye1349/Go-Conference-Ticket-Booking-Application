# Go Conference Ticket Booking Application

This is a simple web application written in Go language. It simulates booking tickets for a conference. Users can input their first name, last name, email address, and the number of tickets they wish to purchase.

## Feature

1. Form Submission: Users can fill out a form with their first name, last name, email address, and the number of tickets they wish to purchase.

2. Input Validation: The application checks whether the user input is valid, i.e., whether the name is of sufficient length, whether the email contains '@', and whether the number of tickets requested is valid.

3. Ticket Booking: If user input is valid, the application proceeds to book the requested number of tickets and reduces the total number of available tickets accordingly.

4. Ticket Sending Simulation: The application simulates the process of sending a confirmation email to the user after they have successfully booked their tickets.

5. Concurrency: The ticket sending simulation is run concurrently using goroutines, providing an asynchronous user experience.

## Packages and Functions Used

This web application utilizes several built-in packages in the Go standard library:

- `fmt`: Used for formatted I/O with functions like Println and Printf.
- `net/http`: Used to set up the HTTP server and handle HTTP requests.
- `html/template`: Used to parse and execute HTML templates.
- `sync`: Used for managing concurrent execution with the WaitGroup type.
- `time`: Used to create delays with Sleep.

## Code Details

### Data Types and Variables

- `String Data Type`: Used for storing users' first names, last names, and email addresses.
- `Integer Data Type`: Used for storing the number of tickets each user wants to book and the remaining number of tickets.
- `Struct`: The UserData struct is used to store information for each booking. It includes first name, last name, email, and number of tickets.

### Logic Statement

- `IF Statement`: Used for validating user input and managing the ticket booking process.
- `For loop`: Used for iterating over the bookings to get first names.

### Goroutines

- `Goroutines`: Used for simulating sending tickets to users concurrently.

## Dockerfile

The Dockerfile in this project allows you to containerize the application. This involves two stages: building the binary in a temporary container, and copying the built binary into a new, final container to run the application.

1. Build Stage: During this stage, the Go SDK is used to compile the application into a binary file. All Go dependencies are also installed during this stage.

2. Final Stage: In this stage, the built binary file is copied into a new Docker container. This final Docker container is much smaller in size since it doesn't contain the Go SDK or any other build-time dependencies, only the binary file and any other necessary runtime dependencies.
