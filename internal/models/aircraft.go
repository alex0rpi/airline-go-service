package models

type Aircraft struct {
	ID				  uint 	 `gorm:"primaryKey" json:"id"`
	Registration	  string `gorm:"not null;unique" json:"registration"`
	Model             string `gorm:"not null" json:"model"`
	ManufactureDate   string `gorm:"not null" json:"manufacture_date"`
	TotalCycles       string `gorm:"not null" json:"total_cycles"`
	TotalFlightHours  uint   `gorm:"not null" json:"total_flight_hours"`
	Status            string `gorm:"not null" json:"status"`
	
}
