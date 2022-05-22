package rest

import (
	"os"
	"testing"

	"github.com/acargorkem/ecommerce_oauth-api/src/domain/users"
	"github.com/acargorkem/ecommerce_oauth-api/src/utils/errors"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	mockClient := usersRestClient.GetClient()
	mockClient.Transport = httpmock.DefaultTransport
	httpmock.ActivateNonDefault(mockClient)
	os.Exit(m.Run())
}

func TestLoginUserNotFound(t *testing.T) {
	httpmock.Reset()
	defer httpmock.DeactivateAndReset()

	repository := usersRepository{}
	errorMessage := "user not found"

	responseBody := errors.NewNotFoundError(errorMessage)
	responder := httpmock.NewJsonResponderOrPanic(404, responseBody)
	fakeUrl := baseUrl + "/users/login"
	httpmock.RegisterResponder("POST", fakeUrl, responder)

	user, err := repository.LoginUser("", "")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, 404, err.Status)
	assert.Equal(t, errorMessage, err.Message)
	assert.Equal(t, "not_found", err.Error)
}

func TestLoginUserInvalidCredentials(t *testing.T) {
	httpmock.Reset()
	defer httpmock.DeactivateAndReset()

	repository := usersRepository{}
	errorMessage := "invalid credentials"

	responseBody := errors.NewUnauthorizedError(errorMessage)
	responder := httpmock.NewJsonResponderOrPanic(401, responseBody)
	fakeUrl := baseUrl + "/users/login"
	httpmock.RegisterResponder("POST", fakeUrl, responder)

	user, err := repository.LoginUser("", "")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, 401, err.Status)
	assert.Equal(t, errorMessage, err.Message)
	assert.Equal(t, "unauthorized", err.Error)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	httpmock.Reset()
	defer httpmock.DeactivateAndReset()

	repository := usersRepository{}
	responseBody := `{}`
	responder := httpmock.NewJsonResponderOrPanic(500, responseBody)
	fakeUrl := baseUrl + "/users/login"
	httpmock.RegisterResponder("POST", fakeUrl, responder)

	user, err := repository.LoginUser("", "")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, 500, err.Status)
	assert.Equal(t, "internal_server_error", err.Error)
}

func TestLoginUserSuccessful(t *testing.T) {
	httpmock.Reset()
	defer httpmock.DeactivateAndReset()

	repository := usersRepository{}
	responseBody := users.User{
		Id:        123,
		FirstName: "gorkem",
		LastName:  "acar",
		Email:     "acargorkem@outlook.com",
	}
	responder := httpmock.NewJsonResponderOrPanic(200, responseBody)
	fakeUrl := baseUrl + "/users/login"
	httpmock.RegisterResponder("POST", fakeUrl, responder)

	user, err := repository.LoginUser("", "")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, responseBody.Id, user.Id)
	assert.Equal(t, responseBody.FirstName, user.FirstName)
	assert.Equal(t, responseBody.LastName, user.LastName)
	assert.Equal(t, responseBody.Email, user.Email)
}
