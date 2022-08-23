package response

import (
	rr "github.com/hlts2/round-robin"
	"io"
	"net/http"
	"net/url"
)

var UserServiceRoundRobin, _ = rr.New(
	&url.URL{Host: "http://localhost:8081"},
)

var VehicleServiceRoundRobin, _ = rr.New(
	&url.URL{Host: "http://localhost:8082"},
)

var CarServiceServiceRoundRobin, _ = rr.New(
	&url.URL{Host: "http://localhost:8083"},
)

var AppointmentServiceRoundRobin, _ = rr.New(
	&url.URL{Host: "http://localhost:8084"},
)

var ReviewServiceRoundRobin, _ = rr.New(
	&url.URL{Host: "http://localhost:8085"},
)

var ReportServiceRoundRobin, _ = rr.New(
	&url.URL{Host: "http://localhost:8087"},
)

func DelegateResponse(response *http.Response, w http.ResponseWriter) {
	w.WriteHeader(response.StatusCode)
	_, err := io.Copy(w, response.Body)
	if err != nil {
		return
	}
	err = response.Body.Close()
	if err != nil {
		return
	}
}

func HandleRequest(w http.ResponseWriter, r *http.Request, url string) {
	if r.Method == "OPTIONS" {
		return
	}

	req, _ := http.NewRequest(r.Method, url, r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	DelegateResponse(res, w)
}
