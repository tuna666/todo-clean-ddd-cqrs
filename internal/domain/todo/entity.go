package todo

import "time"

type ID uint

type Todo struct {
	ID        ID        `gorm:"primaryKey"`
	Title     string    `gorm:"size:255;not null"`
	Done      bool      `gorm:"default:false"`
	CreatedAt time.Time
}
