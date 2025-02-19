package model

import "time"

// https://go.dev/doc/tutorial/web-service-gin

const IdentityKey = "id"

type Widget struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string `json:"name"`
	Description string `json:"description"`
	Count       int    `json:"count"`
	Creator     string `json:"creator"`
	Updater     string `json:"updater"`
}

// https://www.baeldung.com/rest-api-error-handling-best-practices
type ErrorDetail struct {
	Type     string `json:"type,omitempty" yaml:"type,omitempty"`
	Title    string `json:"title,omitempty" yaml:"title,omitempty"`
	Status   int64  `json:"status,omitempty" yaml:"status,omitempty"`
	Detail   string `json:"detail,omitempty" yaml:"detail,omitempty"`
	Instance string `json:"instance,omitempty" yaml:"instance,omitempty"`
}

type Time struct {
	CurrentTime time.Time
	Name        string
	ThrowError  bool
	Message     string
}
