package token

import "time"

type Maker interface {
	CreateToekn(username string, duration time.Duration) (string, error)

	VerifyToekn(token string) (*Payload, error)
}
