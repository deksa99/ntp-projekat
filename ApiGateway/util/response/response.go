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
	w.Header().Set("Content-Type", response.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", response.Header.Get("Content-Length"))
	w.Header().Set("Access-Control-Allow-Origin", "*")
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

func CorsResponseHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func HandleRequest(w http.ResponseWriter, r *http.Request, url string) {
	CorsResponseHeaders(&w)
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
