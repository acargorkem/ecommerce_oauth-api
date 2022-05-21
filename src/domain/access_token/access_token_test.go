package access_token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {

	assert.Equal(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAcessToken(t *testing.T) {
	at := GetNewAccessToken()

	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
	assert.Equal(t, "", at.AccessToken, "new access token should not have defined access token id")
	assert.Equal(t, int64(0), at.UserId, "new access token should not have an associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = getExpirationTimestamp(3)
	assert.False(t, at.IsExpired(), "access token expiring three hours from now should not be expired")
}
