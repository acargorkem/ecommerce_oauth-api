package access_token

import (
	"strings"
	"time"

	"github.com/acargorkem/ecommerce_oauth-api/src/utils/config"
	"github.com/acargorkem/ecommerce_utils-go/rest_errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grandTypeClientCredentials = "client_credentials"
)

var (
	jwtSecretKey = config.JWT_SECRET_KEY
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

func (request *AccessTokenRequest) Validate() *rest_errors.RestErr {
	switch request.GrantType {
	case grantTypePassword:
		break
	case grandTypeClientCredentials:
	default:
		return rest_errors.NewBadRequestError("invalid grant_type parameter")
	}

	// TODO: Validate parameters for each grant_type
	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	ExpiredAt   int64  `json:"expired_at"`
}

func (at *AccessToken) Validate() *rest_errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return rest_errors.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return rest_errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return rest_errors.NewBadRequestError("invalid client id")
	}
	if at.ExpiredAt <= 0 {
		return rest_errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func getExpirationTimestamp(expiredIn time.Duration) int64 {
	return time.Now().UTC().Add(expiredIn * time.Hour).Unix()
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:    userId,
		ExpiredAt: getExpirationTimestamp(expirationTime),
	}
}

func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.ExpiredAt, 0)
	return expirationTime.Before(now)
}

func (at *AccessToken) Generate() *rest_errors.RestErr {
	maker, err := NewJWTMaker(jwtSecretKey)
	if err != nil {
		return rest_errors.NewInternalServerError("sorry can't create access token", rest_errors.NewError("An error during creating jwt maker"))
	}

	token, err := maker.CreateToken(at.UserId, at.ExpiredAt)
	if err != nil {
		return rest_errors.NewInternalServerError("sorry can't create access token", rest_errors.NewError("An error during create token"))
	}

	at.AccessToken = token
	return nil
}
