package responses

import "time"

type UserLogin struct {
	Id        string    `json:"id"`
	Ttl       int       `json:"ttl"`
	CreatedAt time.Time `json:"created"`
	UserId    string    `json:"userId"`
}
