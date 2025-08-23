package token

import "time"

// Maker is and interface for managing tokens
type Maker interface {
	// CreateToken creates a new token for specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	//VerifyToken it will check if token valid or not
	VerifyToken(token string) (*Payload, error)
}
