package errs

import (
	"errors"
	"net/http"
)

type ErrResp struct {
	Message string `json:"message" example:"INTERNAL_SERVER_ERROR"`
	Code    int64  `json:"code" example:"500"`
} //@name ErrResp

func (e ErrResp) ToError() error {
	return errors.New(e.Message)
}

var (
	InternalServerErr        = ErrResp{"INTERNAL_SERVER_ERROR", http.StatusInternalServerError}
	UnauthorizedErr          = ErrResp{"UNAUTHORIZED", http.StatusUnauthorized}
	BadParamInBodyErr        = ErrResp{"BAD_PARAM_IN_BODY", http.StatusBadRequest}
	NotValidBodyParamErr     = ErrResp{"NOT_VALID_BODY_PARAM", http.StatusBadRequest}
	EmailAlreadyExistErr     = ErrResp{"EMAIL_ALREADY_EXISTS", http.StatusForbidden}
	UsernameAlreadyExistsErr = ErrResp{"USERNAME_ALREADY_EXISTS", http.StatusForbidden}
	UserDoesntExistErr       = ErrResp{"USER_DOESNT_EXIST", http.StatusBadRequest}
	WrongCredentialsErr      = ErrResp{"WRONG_EMAIL_OR_PASS", http.StatusBadRequest}
	NoDataInFormErr          = ErrResp{"NO_DATA_IN_FORM", http.StatusBadRequest}
	IncorrectAccountTypeErr  = ErrResp{"INCORRECT_ACCOUNT_TYPE", http.StatusForbidden}
	EmptyQueryParamErr       = ErrResp{"QUERY_PARAM_IS_EMPTY", http.StatusBadRequest}
)
