package data

// MedicationDTO представляет структуру лекарства
type MedicationDTO struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	Dosage    string `json:"dosage"`
	Frequency string `json:"frequency"`
	StartDate string `json:"start_date"` // ISO date format: "2025-04-18"
	EndDate   string `json:"end_date"`
}
