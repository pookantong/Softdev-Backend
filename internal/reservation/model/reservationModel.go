package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatusType string

const (
	Accepted  StatusType = "Accepted"
	Denied    StatusType = "Denied"
	Completed StatusType = "Completed"
	Cancelled StatusType = "Cancelled"
	Pending   StatusType = "Pending"
)

type Reservation struct {
	gorm.Model
	ID           uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID       uuid.UUID  `gorm:"not null"` // foreign key to Customer
	RestaurantID uuid.UUID  `gorm:"not null"`
	TableID      uuid.UUID  `gorm:"type:uuid;not null;foreignkey:Table"`
	StartTime    time.Time  `gorm:"not null;index"`
	EndTime      time.Time  `gorm:"not null;index"`
	Status       StatusType `gorm:"not null"`
	TotalPrice   int        `gorm:"not null"`
	ReviewID     uuid.UUID  `gorm:"type:uuid"`
	DishItems    []DishItem `gorm:"foreignKey:ReservationID"`
}

type DishItem struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	DishID        uuid.UUID `gorm:"not null"`
	ReservationID uuid.UUID `gorm:"not null"`
	Quantity      int       `gorm:"not null"`
	Price         int       `gorm:"not null"`
	Option        string
	Comment       string
}
