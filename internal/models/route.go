package models

type Route struct {
	ID                int     `gorm:"primaryKey" json:"id"`
	Code		  	  string  `gorm:"not null;unique" json:"route_code"`
    Origin            string  `json:"origin"`
    Destination       string  `json:"destination"`
    EstimatedDuration string  `json:"estimated_duration"`
    Distance          float64 `json:"distance"`
}