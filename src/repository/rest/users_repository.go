package rest

import (
	"time"

	"github.com/acargorkem/ecommerce_oauth-api/src/domain/users"
	"github.com/acargorkem/ecommerce_utils-go/rest_errors"
	resty "github.com/go-resty/resty/v2"
)

const (
	baseUrl = "http://localhost:8081"
)

var (
	usersRestClient = resty.New().
		SetBaseURL(baseUrl).SetTimeout(30 * time.Second)
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *rest_errors.RestErr)
}

type usersRepository struct {
}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *rest_errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	var user users.User
	var restErr rest_errors.RestErr
	response, err := usersRestClient.R().
		SetBody(request).
		SetResult(&user).
		SetError(&restErr).
		Post("/users/login")
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when sending request to users api", rest_errors.NewError("rest_client_error"))
	}
	if response.IsError() {
		return nil, &restErr
	}

	return &user, nil
}
