package models

type Schedule struct {
	ID			 uint           `gorm:"primaryKey"`
	RouteCode 	 string `gorm:"not null" json:"route_code"`
	Route		 Route    `gorm:"foreignKey:RouteCode;references:RouteCode" json:"route"`
	DaysOfWeek []int `gorm:"type:integer[]" json:"days_of_week"`
	DepartureTime  string `gorm:"not null"`
}