package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// UserData represents the structure of the JSON data expected from the client.
type UserData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Define a handler function for the root ("/") route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Handle GET requests
			fmt.Fprintf(w, "Welcome to the Go Web Server!")

		case http.MethodPost:
			// Handle POST requests with both JSON and form data
			contentType := r.Header.Get("Content-Type")
			switch contentType {
			case "application/json":
				// Handle JSON data
				var userData UserData
				err := json.NewDecoder(r.Body).Decode(&userData)
				if err != nil {
					http.Error(w, "Error parsing JSON", http.StatusBadRequest)
					return
				}

				// Print the received JSON data
				fmt.Printf("Received JSON data: %+v\n", userData)

				// Respond to the client
				fmt.Fprintf(w, "Hello, %s! Your email is %s.", userData.Name, userData.Email)

			case "application/x-www-form-urlencoded":
				// Handle form data
				err := r.ParseForm()
				if err != nil {
					http.Error(w, "Error parsing form data", http.StatusBadRequest)
					return
				}

				// Access form values
				name := r.Form.Get("name")
				email := r.Form.Get("email")

				// Respond to the client
				fmt.Fprintf(w, "Hello, %s! Your email is %s.", name, email)

			default:
				http.Error(w, "Unsupported content type", http.StatusUnsupportedMediaType)
			}

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Handle POSTForm requests with form-encoded data
	http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form data", http.StatusBadRequest)
				return
			}

			// Access form values
			name := r.Form.Get("name")
			email := r.Form.Get("email")

			// Respond to the client
			fmt.Fprintf(w, "Hello from POSTForm, %s! Your email is %s.", name, email)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Specify the port to listen on
	port := 8080

	// Start the web server
	fmt.Printf("Server is listening on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

//http://localhost:8080/postform
//http://localhost:8080/post
//http://localhost:8080/get
//http://localhost:8080
