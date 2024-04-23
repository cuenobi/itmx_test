package domain

import (
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("given Param is not valid")

	// 400 BadRequest
	ErrUploadLimit          = errors.New("upload limit reached")
	ErrVoteLimit            = errors.New("vote limit reached")
	ErrInvalidUserID        = errors.New("invalid userID")
	ErrIsNotWalkIn          = errors.New("this user is not walk-in register")
	ErrIsNotPass            = errors.New("this user this user criterion is not pass")
	ErrInvalidVideoID       = errors.New("invalid videoID")
	ErrVdoAndUserIDNotMatch = errors.New("videoID and userID is not match")
	ErrInValidVdoUrl        = errors.New("video url is invalid")
	ErrInvalidRegistType    = errors.New("invalid register type")

	// 401 StatusInvalidCredentials
	ErrStatusInvalidCredentials = errors.New("invalid credentials")

	// 403 StatusForbidden
	ErrPermissionDenied = errors.New("permission denied")
	ErrInvalidRecaptcha = errors.New("invalid recaptcha")

	// 404 StatusNotFound
	ErrProvinceNotFound = errors.New("province is not found")
	ErrUsernameNotFound = errors.New("username not found in the system")
	ErrScoutNotFound    = errors.New("scout not found in the system")

	// 409 StatusConflict
	ErrUsernameExist = errors.New("username already exists")
	ErrEmailExist    = errors.New("email already exists")
	ErrPhoneExist    = errors.New("phone number already exists")
	ErrDupPhoneExist = errors.New("phone number and phone number backup are duplicate")
	ErrIdNumberExist = errors.New("id card number already exists")
	ErrScoreExist    = errors.New("score is already exists")
	ErrCriState      = errors.New("criterion state is already updated")
	ErrVdoUrlExist   = errors.New("video url is already exists")
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError

	// 400 StatusBadRequest
	case ErrUploadLimit:
		return http.StatusBadRequest
	case ErrInvalidUserID:
		return http.StatusBadRequest
	case ErrIsNotWalkIn:
		return http.StatusBadRequest
	case ErrIsNotPass:
		return http.StatusBadRequest
	case ErrInvalidVideoID:
		return http.StatusBadRequest
	case ErrVdoAndUserIDNotMatch:
		return http.StatusBadRequest
	case ErrVoteLimit:
		return http.StatusBadRequest
	case ErrInValidVdoUrl:
		return http.StatusBadRequest
	case ErrInvalidRegistType:
		return http.StatusBadRequest

	// 401 StatusUnauthorized
	case ErrStatusInvalidCredentials:
		return http.StatusUnauthorized

	// 403 StatusForbidden
	case ErrPermissionDenied:
		return http.StatusForbidden
	case ErrInvalidRecaptcha:
		return http.StatusForbidden

	// 404 StatusNotFound
	case ErrNotFound:
		return http.StatusNotFound
	case ErrProvinceNotFound:
		return http.StatusNotFound
	case ErrUsernameNotFound:
		return http.StatusNotFound

	// 409 StatusConflict
	case ErrEmailExist:
		return http.StatusConflict
	case ErrPhoneExist:
		return http.StatusConflict
	case ErrIdNumberExist:
		return http.StatusConflict
	case ErrScoreExist:
		return http.StatusConflict
	case ErrCriState:
		return http.StatusConflict
	case ErrVdoUrlExist:
		return http.StatusConflict

	default:
		return http.StatusInternalServerError
	}
}
