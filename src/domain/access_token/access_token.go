package access_token

import (
	"strings"
	"time"

	"github.com/acargorkem/ecommerce_oauth-api/src/utils/errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grandTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`

	// Used for password grant_type
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for client credentials grant_type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`

	Scope string `json:"scope"`
}

func (request *AccessTokenRequest) Validate() *errors.RestErr {
	switch request.GrantType {
	case grantTypePassword:
		break
	case grandTypeClientCredentials:
	default:
		return errors.NewBadRequestError("invalid grant_type parameter")
	}

	// TODO: Validate parameters for each grant_type
	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func getExpirationTimestamp(expiredIn time.Duration) int64 {
	return time.Now().UTC().Add(expiredIn * time.Hour).Unix()
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: getExpirationTimestamp(expirationTime),
	}
}

func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	return expirationTime.Before(now)
}

func (at *AccessToken) Generate() {
	//TODO: Implement real access token
	at.AccessToken = "test"
}
