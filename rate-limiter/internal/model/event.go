package model

type RateEvent struct {
	UserID  string `json:"user_id"`
	Allowed bool   `json:"allowed"`
}
