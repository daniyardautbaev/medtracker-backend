package data

// ReminderDTO описывает напоминание о приёме лекарства
type ReminderDTO struct {
	ID             string `json:"id"`
	MedicationID   string `json:"medication_id"`
	Time           string `json:"time"` // формат HH:MM
	Frequency      string `json:"frequency"`
	MedicationName string `json:"medication_name"`
	UserID         string `json:"user_id"`
}
