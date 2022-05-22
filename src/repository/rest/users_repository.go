package rest

import (
	"time"

	"github.com/acargorkem/ecommerce_oauth-api/src/domain/users"
	"github.com/acargorkem/ecommerce_oauth-api/src/utils/errors"
	resty "github.com/go-resty/resty/v2"
)

const (
	baseUrl = "http://fake-users-api-url:8080"
)

var (
	usersRestClient = resty.New().
		SetBaseURL(baseUrl).SetTimeout(30 * time.Second)
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct {
}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	var user users.User
	var restErr errors.RestErr
	response, err := usersRestClient.R().
		SetBody(request).
		SetResult(&user).
		SetError(&restErr).
		Post("/users/login")
	if err != nil {
		return nil, errors.NewInternalServerError("error when sending request to users api")
	}
	if response.IsError() {
		return nil, &restErr
	}

	return &user, nil
}
