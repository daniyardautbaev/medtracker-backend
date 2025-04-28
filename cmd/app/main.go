package main

import (
	"flag"
	"log"

	"medtracker/medtracker/internal/app"
	"medtracker/medtracker/internal/app/config"
	"medtracker/medtracker/internal/app/connections"
	"medtracker/medtracker/internal/deliveries"
	"medtracker/medtracker/internal/repositories"
	"medtracker/medtracker/internal/services"
)

func main() {
	// Чтение конфига
	configFile := flag.String("config", "./configs/.env", "Path to config file")
	flag.Parse()

	cfg, err := config.NewConfig(*configFile)
	if err != nil {
		log.Fatalf("❌ Error loading config: %v", err)
	}

	// Подключение к БД
	conn, err := connections.NewConnections(cfg)
	if err != nil {
		log.Fatalf("❌ Error connecting to DB: %v", err)
	}
	defer conn.Close()

	// Репозитории
	patientRepo := repositories.NewMySQLPatientRepository(conn.DB)
	medicationRepo := repositories.NewMySQLMedicationRepository(conn.DB)
	userRepo := repositories.NewMySQLUserRepository(conn.DB)
	reminderRepo := repositories.NewMySQLReminderRepository(conn.DB)
	intakeRepo := repositories.NewMySQLIntakeLogRepository(conn.DB)

	// Сервисы
	patientService := services.NewPatientService(patientRepo)
	medicationService := services.NewMedicationService(medicationRepo)
	userService := services.NewUserService(userRepo)
	intakeService := services.NewIntakeLogService(intakeRepo)
	reminderService := services.NewReminderService(reminderRepo, intakeRepo) // !! два репозитория

	// Handlers
	patientHandler := deliveries.NewPatientHandler(patientService)
	medicationHandler := deliveries.NewMedicationHandler(medicationService)
	userHandler := deliveries.NewUserHandler(userService)
	reminderHandler := deliveries.NewReminderHandler(reminderService)
	intakeHandler := deliveries.NewIntakeLogHandler(intakeService)

	// Запуск приложения
	app.Run(cfg, patientHandler, medicationHandler, userHandler, reminderHandler, intakeHandler)
}
