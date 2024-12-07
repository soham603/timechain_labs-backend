package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
)

func Handler(w http.ResponseWriter, r *http.Request) {
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
	password := "twwt apln ubyx rtdz" // Use environment variable for security
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"
	to := "soham.lad16793@sakec.ac.in"
	auth := smtp.PlainAuth("", from, password, smtpServer)

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"From: " + from + "\r\n" +
		"\r\n" + emailBody)

	err = smtp.SendMail(smtpServer+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		fmt.Println("Error sending email:", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	fmt.Println("Email sent successfully to:", to)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Form Data Received Successfully and Email Sent Successfully to %s", to)))
}
