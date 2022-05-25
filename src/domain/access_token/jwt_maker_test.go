package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	secretKey = "itsupposetobeasecret"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(secretKey)
	assert.NoError(t, err)

	var userId int64 = 5
	expiredAt := getExpirationTimestamp(expirationTime)

	token, err := maker.CreateToken(userId, expiredAt)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	assert.NoError(t, err)
	assert.NotEmpty(t, payload)

	assert.NotZero(t, payload.Id)
	assert.Equal(t, userId, payload.UserId)
	assert.WithinDuration(t, time.Unix(expiredAt, 0), time.Unix(payload.ExpiredAt, 0), time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(secretKey)
	assert.NoError(t, err)

	var userId int64 = 5
	expiredAt := getExpirationTimestamp(-expirationTime)

	token, err := maker.CreateToken(userId, expiredAt)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrExpiredToken.Error())
	assert.Nil(t, payload)
}
