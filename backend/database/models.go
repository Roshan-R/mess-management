package database

import "gorm.io/datatypes"

type Menu struct {
	ID         uint
	ActiveDate datatypes.Date
	Breakfast  string
	Lunch      string
	Dinner     string
}

type PlanType struct {
	ID           uint
	PlanName     string
	NumberOfDays uint8
	TimesPerDay  uint8
	Amount       uint16
	//FoodType string
	MaxAllowedMessoutDays uint8
}

type Plan struct {
	StartDate datatypes.Date
	EndDate   datatypes.Date
	//User
	PlanTypeID uint
	IsActive   bool
	HasPaid    bool
	//MessoutDays []datatypes.Date
	PlanType PlanType
}

type User struct {
	PhoneNumber  string
	Name         string
	ActivePlanID Plan
	PlanHistory  []Plan
}
