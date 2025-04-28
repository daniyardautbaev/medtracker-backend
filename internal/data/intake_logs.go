package data

// IntakeLogDTO представляет лог приёма лекарства
type IntakeLogDTO struct {
	ID           string `json:"id"`
	MedicationID string `json:"medication_id"`
	IntakeTime   string `json:"intake_time"`
	Taken        bool   `json:"taken"`
}
