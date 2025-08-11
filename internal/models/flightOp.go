package models

type FlightOp struct {
	ID			   uint		`gorm:"primaryKey" json:"id"`
	DateTime 	   string	`gorm:"not null" json:"date_time"`
	RouteCode 	   string 	`json:"route_code" gorm:"not null"`
	ACRegistration string	`gorm:"not null" json:"aircraft_registration"`
	Status    	   string   `json:"status" gorm:"type:enum('scheduled', 'boarding', 'in_flight', 'arrived', 'cancelled', 'delayed', 'no_info')"`
	Aircraft       Aircraft `gorm:"foreignKey:ACRegistration;references:Registration" json:"aircraft"`
	Route		   Route    `gorm:"foreignKey:RouteCode;references:RouteCode" json:"route"`
}