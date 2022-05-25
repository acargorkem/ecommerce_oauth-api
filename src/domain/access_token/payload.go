package access_token

import (
	"time"

	"github.com/acargorkem/ecommerce_utils-go/rest_errors"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = rest_errors.NewError("token is invalid")
	ErrExpiredToken = rest_errors.NewError("token has expired")
)

type Payload struct {
	Id        uuid.UUID `json:"id"`
	UserId    int64     `json:"user_id"`
	ExpiredAt int64     `json:"expired_at"`
}

func NewPayload(userId int64, expiredAt int64) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		Id:        tokenId,
		UserId:    userId,
		ExpiredAt: expiredAt,
	}

	return payload, nil
}

func (paylod *Payload) Valid() error {
	now := time.Now().UTC()
	expirationTime := time.Unix(paylod.ExpiredAt, 0)
	if expirationTime.Before(now) {
		return ErrExpiredToken
	}
	return nil
}
