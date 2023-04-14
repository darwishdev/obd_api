package auth

import "time"

type Maker interface {
	// CreateToken(username string, duration time.Duration) (string, error)
	CreateToken(username string, userId int64, carId int64, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
