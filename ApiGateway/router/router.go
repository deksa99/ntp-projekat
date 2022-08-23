package router

import (
	"ApiGateway/handler"
	cHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	//User service
	router.Path("/users/login").Queries("role", "{role:user|admin|worker}").HandlerFunc(handler.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/users/register", handler.Register).Methods("POST", "OPTIONS")
	router.HandleFunc("/users/change-password", handler.ChangePassword).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/users/{id}", handler.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/users", handler.GetAllUsers).Methods("GET", "OPTIONS")
	router.HandleFunc("/users/{id}/block", handler.BlockUser).Methods("PATCH", "OPTIONS")

	//Vehicle service
	router.HandleFunc("/vehicles/add", handler.AddVehicle).Methods("POST", "OPTIONS")
	router.HandleFunc("/vehicles/{vehicleId}", handler.UpdateVehicle).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/vehicles/all", handler.GetVehiclesForUser).Methods("GET", "OPTIONS")

	//CarService service
	router.HandleFunc("/car-services", handler.GetAllCarServices).Methods("GET", "OPTIONS")
	router.HandleFunc("/car-services", handler.CreateCarService).Methods("POST", "OPTIONS")
	router.HandleFunc("/car-services/near-me", handler.FindNearest).Methods("POST", "OPTIONS")
	router.HandleFunc("/car-services/service", handler.CreateService).Methods("POST", "OPTIONS")
	router.HandleFunc("/car-services/service", handler.UpdateService).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/car-services/service/{id}/change-availability", handler.ChangeAvailability).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/car-services/{id}/catalog", handler.GetCatalog).Methods("GET", "OPTIONS")

	//Appointment service
	router.HandleFunc("/appointments/requests", handler.CreateAppointmentRequest).Methods("POST", "OPTIONS")
	router.HandleFunc("/appointments/requests/cancel", handler.CancelRequest).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/appointments/requests", handler.GetRequestsForUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/appointments/user", handler.GetAppointmentsForUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/appointments/requests/{requestId}/accept", handler.AcceptRequest).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/appointments/requests/{requestId}/reject", handler.RejectRequest).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/appointments/{appointmentId}/finish", handler.FinishAppointment).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/appointments/{appointmentId}/cancel", handler.CancelAppointment).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/appointments/requests/service/{carServiceId}", handler.GetRequestsForCarService).Methods("GET", "OPTIONS")
	router.HandleFunc("/appointments/car-service/{carServiceId}", handler.GetAppointmentsForCarService).Methods("GET", "OPTIONS")
	router.HandleFunc("/appointments/worker", handler.GetAppointmentsForWorker).Methods("GET", "OPTIONS")

	//Review service
	router.HandleFunc("/reviews/reported", handler.GetReports).Methods("GET", "OPTIONS")
	router.HandleFunc("/reviews/car-service/{carServiceId}", handler.GetReviews).Methods("GET", "OPTIONS")
	router.HandleFunc("/reviews/{id}/report", handler.ReportReview).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/reviews/report/{id}/process", handler.ProcessReport).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/reviews/add", handler.AddReview).Methods("POST", "OPTIONS")

	//Report service
	router.HandleFunc("/reports/manufacturer/{carServiceId}", handler.ManufacturerReport).Methods("GET", "OPTIONS")
	router.HandleFunc("/reports/service/{carServiceId}", handler.ServiceReport).Methods("GET", "OPTIONS")
	router.HandleFunc("/reports/financial/{carServiceId}", handler.FinancialReport).Methods("GET", "OPTIONS")
	router.HandleFunc("/reports/status/{carServiceId}", handler.StatusReport).Methods("GET", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8090", cHandlers.CORS(
		cHandlers.AllowedMethods([]string{"GET", "POST", "PATCH", "PUT", "OPTIONS"}),
		cHandlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin", "Authorization"}),
		cHandlers.AllowedOrigins([]string{"*"}))(router)))
}
