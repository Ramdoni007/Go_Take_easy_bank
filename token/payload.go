package token

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

// Different Types Of Error returned by the VerifyToken func
var (
	errExpiredToken = errors.New("Token has expired")
	errInvalidToken = errors.New("Token has invalid")
)

// NewPayload creates a new token payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return errExpiredToken
	}
	return nil
}
