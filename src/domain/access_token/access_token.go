package access_token

import (
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func getExpirationTimestamp(expiredIn time.Duration) int64 {
	return time.Now().UTC().Add(expiredIn * time.Hour).Unix()
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: getExpirationTimestamp(expirationTime),
	}
}

func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	return expirationTime.Before(now)
}
