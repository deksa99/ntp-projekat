package handler

import (
	"EmailService/helper"
	"EmailService/model"
	"EmailService/request"
	"EmailService/response"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"
)

func SendMail(w http.ResponseWriter, r *http.Request) {
	var email request.Email

	err := helper.DecodeJSONBody(w, r, &email)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		var e *response.Error
		if errors.As(err, &e) {
			http.Error(w, e.Message, e.Status)
		} else {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	sender := os.Getenv("EMAIL")

	user := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASS")

	server := "smtp-mail.outlook.com"
	port := "587"
	addr := server + ":" + port

	to := []string{email.To}

	req := model.Mail{
		Sender:  sender,
		To:      []string{email.To},
		Subject: email.Subject,
		Body:    email.Text + "\n" + email.Html,
	}

	msg := buildMessage(req)
	auth := model.LoginAuth(user, password)
	err = smtp.SendMail(addr, auth, sender, to, []byte(msg))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(response.Error{Message: err.Error(), Status: http.StatusBadRequest})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(response.Success{Message: "Email sent successfully"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func buildMessage(mail model.Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}
