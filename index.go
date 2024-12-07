package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
)

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! This is your API response.")
	})

	http.HandleFunc("/submitForm", formDataAndEmail)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func formDataAndEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read and parse the JSON body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var formData map[string]interface{}
	err = json.Unmarshal(body, &formData)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	fmt.Println("Received Data from Form:", formData)

	// Construct the email body as a formatted string
	subject := "New Hiring Request Application"
	emailBody := fmt.Sprintf(`New Application Received for Hiring Application:

Application Details:
Name: %s
Company: %s
Email: %s
Phone: %s
Country: %s
City: %s
Employment Type: %s
Skill Level: %s
Position: %s
Budget: %s
Additional Information: %s
`,
		formData["name"],
		formData["company"],
		formData["email"],
		formData["contact"],
		formData["country"],
		formData["city"],
		formData["employmentType"],
		formData["skillLevel"],
		formData["position"],
		formData["budget"],
		formData["additionalInfo"],
	)

	from := "sohamxladftp@gmail.com"
	password := "twwt apln ubyx rtdz" // Replace this with your actual App Password
	smtpServer := "smtp.gmail.com"
	smtpPort := "587" // Port for sending emails using TLS
	//to := "deepali.chopra@timechainlabs.io"
	to := "soham.lad16793@sakec.ac.in"
	// Create authentication for sending the email
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// Prepare the email message
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"From: " + from + "\r\n" +
		"\r\n" + emailBody)

	// Send the email
	err = smtp.SendMail(smtpServer+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		fmt.Println("Error sending email:", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	// Log success and return a response
	fmt.Println("Email sent successfully to:", to)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Form Data Received Successfully and Email Sent Successfully to %s", to)))
}
