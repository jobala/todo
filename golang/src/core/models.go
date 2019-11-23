package core

import "time"

// ToDo data structure
type ToDo struct {
	ID          int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
