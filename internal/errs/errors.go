package errs

import (
	"errors"
	"net/http"
)

type ErrResp struct {
	Message string `json:"error" example:"INTERNAL_SERVER_ERROR"`
	Code    int64  `json:"code" example:"500"`
} //@name ErrResp

func (e ErrResp) ToError() error {
	return errors.New(e.Message)
}

var (
	InternalServerErr          = ErrResp{"INTERNAL_SERVER_ERROR", http.StatusInternalServerError}
	UnauthorizedErr            = ErrResp{"UNAUTHORIZED", http.StatusUnauthorized}
	BadParamInBodyErr          = ErrResp{"BAD_PARAM_IN_BODY", http.StatusConflict}
	NotValidBodyParamErr       = ErrResp{"NOT_VALID_BODY_PARAM", http.StatusBadRequest}
	EmailAlreadyExistErr       = ErrResp{"EMAIL_ALREADY_EXISTS", http.StatusConflict}
	UserAlreadyExistsErr       = ErrResp{"User already exists", http.StatusConflict}
	UserNotFoundErr            = ErrResp{"USER_NOT_FOUND", http.StatusBadRequest}
	WrongCredentialsErr        = ErrResp{"Invalid password or login", http.StatusBadRequest}
	NoDataInFormErr            = ErrResp{"NO_DATA_IN_FORM", http.StatusBadRequest}
	IncorrectAccountTypeErr    = ErrResp{"INCORRECT_ACCOUNT_TYPE", http.StatusForbidden}
	EmptyQueryParamErr         = ErrResp{"QUERY_PARAM_IS_EMPTY", http.StatusBadRequest}
	NotVerifiedAccountErr      = ErrResp{"NOT_VERIFIED_ACCOUNT", http.StatusForbidden}
	UserSocialAlreadyExistsErr = ErrResp{"USER_SOCIAL_ALREADY_EXISTS", http.StatusConflict}
	NovelNotFoundErr           = ErrResp{"NOVEL_NOT_FOUND", http.StatusBadRequest}
	CompetitonNotFoundErr      = ErrResp{"COMPETITION_NOT_FOUND", http.StatusBadRequest}
	QueryParamIsNotValidErr    = ErrResp{"NOT_VALID_QUERY_PARAM", http.StatusBadRequest}
	CompetitionIsNotActiveErr  = ErrResp{"COMPETITION_IS_NOT_ACTIVE", http.StatusConflict}
	IncorrectUserForVotingErr  = ErrResp{"INCORRECT_USER_FOR_VOTING", http.StatusConflict}
	UserAlreadyVotedErr        = ErrResp{"USER_ALREADY_VOTED", http.StatusConflict}
)
