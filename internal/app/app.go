package app

import (
	"fmt"
	"log"
	"medtracker/medtracker/internal/app/config"
	"medtracker/medtracker/internal/deliveries"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Run(cfg *config.Config,
	patientHandler *deliveries.PatientHandler,
	medHandler *deliveries.MedicationHandler,
	userHandler *deliveries.UserHandler,
	reminderHandler *deliveries.ReminderHandler,
	intakeHandler *deliveries.IntakeLogHandler) {

	// Инициализируем маршруты
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}).Methods("GET")

	// Пациенты
	r.HandleFunc("/patient", patientHandler.CreatePatientHandler).Methods("POST")
	r.HandleFunc("/patient/{id}", patientHandler.GetPatientHandler).Methods("GET")
	r.HandleFunc("/patient/{id}", patientHandler.UpdatePatientHandler).Methods("PUT")
	r.HandleFunc("/patient/{id}", patientHandler.DeletePatientHandler).Methods("DELETE")

	// Лекарства
	r.HandleFunc("/medication", medHandler.CreateMedication).Methods("POST")
	r.HandleFunc("/medication/{id}", medHandler.GetMedication).Methods("GET")
	r.HandleFunc("/medication/{id}", medHandler.UpdateMedication).Methods("PUT")
	r.HandleFunc("/medication/{id}", medHandler.DeleteMedication).Methods("DELETE")
	r.HandleFunc("/medications", medHandler.ListMedicationsByUser).Methods("GET")

	// Пользователи
	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/login", userHandler.LoginUser).Methods("POST")

	// Напоминания
	r.HandleFunc("/reminder", reminderHandler.CreateReminder).Methods("POST")
	r.HandleFunc("/reminder/{id}", reminderHandler.GetReminder).Methods("GET")
	r.HandleFunc("/reminder/{id}", reminderHandler.UpdateReminder).Methods("PUT")
	r.HandleFunc("/reminder/{id}", reminderHandler.DeleteReminder).Methods("DELETE")
	r.HandleFunc("/reminder/{id}/confirm", reminderHandler.ConfirmReminder).Methods("POST") // <--- добавили

	// 🔄 Разделение маршрутов
	r.HandleFunc("/reminders/user", reminderHandler.ListRemindersByUser).Methods("GET")
	r.HandleFunc("/reminders/medication", reminderHandler.ListRemindersByMedication).Methods("GET")

	// Журнал приёма
	r.HandleFunc("/intake", intakeHandler.CreateLog).Methods("POST")
	r.HandleFunc("/intake/{id}", intakeHandler.GetLog).Methods("GET")
	r.HandleFunc("/intake/{id}", intakeHandler.DeleteLog).Methods("DELETE")
	r.HandleFunc("/intakes", intakeHandler.ListLogsByMedication).Methods("GET")

	// ✅ Глобальная обработка OPTIONS для всех путей (Preflight запросы от браузера)
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		http.NotFound(w, r)
	})

	// Оборачиваем маршруты в CORS middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3002"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}).Handler(r)

	addr := fmt.Sprintf(":%s", cfg.HTTPServer.Port)
	log.Printf("🚀 Starting server on http://localhost%s", addr)

	if err := http.ListenAndServe(addr, corsHandler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[HTTP] %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
