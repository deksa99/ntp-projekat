package api

import (
	"AppointmentService/response"
	"bytes"
	"encoding/json"
	"errors"
	rr "github.com/hlts2/round-robin"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var emailServiceBasePath, _ = rr.New(
	&url.URL{Host: "http://localhost:8086"},
)

func AppointmentFinished(app *response.AppointmentInfo, errCh chan error) {
	defer close(errCh)
	html, err := parseTemplate("resources/app_finished.html", &app)
	if err != nil {
		errCh <- err
	}
	err = sendEmail(app.Email, "Servis završen", html)
	if err != nil {
		log.Println("Email error")
		errCh <- err
	} else {
		log.Println("Email sent successfully")
	}
	errCh <- nil
}

func AppointmentCancelled(app *response.AppointmentInfo, errCh chan error) {
	defer close(errCh)
	html, err := parseTemplate("resources/app_cancelled.html", &app)
	if err != nil {
		errCh <- err
	}
	err = sendEmail(app.Email, "Servis otkazan", html)
	if err != nil {
		log.Println("Email error")
		errCh <- err
	} else {
		log.Println("Email sent successfully")
	}
	errCh <- nil
}

func RequestAccepted(req *response.AppointmentInfo, errCh chan error) {
	defer close(errCh)
	html, err := parseTemplate("resources/req_accepted.html", &req)
	if err != nil {
		errCh <- err
	}
	err = sendEmail(req.Email, "Servis uspešno zakazan", html)
	if err != nil {
		log.Println("Email error")
		errCh <- err
	} else {
		log.Println("Email sent successfully")
	}
	errCh <- nil
}

func RequestRejected(req *response.AppointmentRequestInfo, errCh chan error) {
	defer close(errCh)
	html, err := parseTemplate("resources/req_rejected.html", &req)
	if err != nil {
		errCh <- err
	}
	err = sendEmail(req.Email, "Zahtev odbijen", html)
	if err != nil {
		log.Println("Email error")
		errCh <- err
	} else {
		log.Println("Email sent successfully")
	}
	errCh <- nil
}

func sendEmail(to string, subject string, html string) error {
	email := Email{
		To:      to,
		Subject: subject,
		Text:    "",
		Html:    html,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(email)
	if err != nil {
		return err
	}

	eUrl := emailServiceBasePath.Next().Host + "/api/email"
	resp, err := http.Post(eUrl, "application/json", &buf)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("email error")
	}

	return nil
}

func parseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
