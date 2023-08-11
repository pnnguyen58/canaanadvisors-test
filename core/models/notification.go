package models

import (
	"time"
)

type Notification struct {
	Id         int64                  `json:"id,omitempty"`
	Message    string                 `json:"message,omitempty"`
	Recipients []int64                `json:"recipients,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
}
