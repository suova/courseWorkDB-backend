package models

import "time"

type Thread struct {
	Thread_created *time.Time `json:"created"`
	Thread_title   string     `json:"title"`
	Id_thread      string      `json:"id" db:"d_thread"`
}
