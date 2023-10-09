package entity

import "time"

type Log struct {
	Activity   string    `json:"activity" db:"activity"`
	Request    string    `json:"request" db:"request"`
	Response   string    `json:"response" db:"response"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}
