package models

import "time"

type Train struct {
	ID                 uint   `gorm:"primaryKey"`
	UID                string `gorm:"unique;not null"`
	Headcode           string `gorm:"not null"`
	ATOCCode           string `gorm:"not null"`
	ATOCDesc           string
	OriginTIPLOC       string    `gorm:"column:origin_tiploc;not null"`
	DestinationTIPLOC  string    `gorm:"column:destination_tiploc;not null"`
	ScheduledDeparture time.Time `gorm:"not null"`
	ExpectedArrival    *time.Time
	LastUpdated        time.Time `gorm:"autoUpdateTime"`
}
