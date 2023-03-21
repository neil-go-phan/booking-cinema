package helper

import "net/http"

type ErrorReponse struct {
	ErrLog             error
	ErrorConvention
}

// ErrorConvention is used to declare ErrorName with 
type ErrorConvention struct {
	ErrorName string
	ErrorMessage	string	`json:"message"`
	HttpStatusCode int
}

func (e ErrorReponse)Error() string {
	return e.ErrorMessage
}

// just return errorName, error handler middleware will handle return
func NewErrorReponseAndLog(err error) *ErrorReponse {
	errorConvention := generateErrorProperty(err.Error())

	e := &ErrorReponse{
		ErrLog:         err,
		ErrorConvention: errorConvention,
	}
	// log

	return e
}

func generateErrorProperty(errorName string) (ErrorConvention) {
	switch errorName {
	case ERROR_VALIDATE_TOKEN_FAIL.ErrorName:
		return ERROR_VALIDATE_TOKEN_FAIL

	case ERROR_NO_REFESH_TOKEN.ErrorName:
		return ERROR_NO_REFESH_TOKEN

	case ERROR_GENERATE_TOKEN_FAIL.ErrorName:
		return ERROR_GENERATE_TOKEN_FAIL

	case ERROR_NO_PERMISSION.ErrorName:
		return ERROR_NO_PERMISSION

	case ERROR_BAD_REQUEST.ErrorName:
		return ERROR_BAD_REQUEST

	case ERROR_INPUT_INVALID.ErrorName:
		return ERROR_INPUT_INVALID

	case ERROR_SIGNIN_INCORRECT.ErrorName:
		return ERROR_SIGNIN_INCORRECT
		
	case ERROR_USERNAME_TAKEN.ErrorName:
		return ERROR_USERNAME_TAKEN

	case ERROR_UPDATE_FAIL.ErrorName:
		return ERROR_UPDATE_FAIL

	case ERROR_DELETE_FAIL.ErrorName:
		return ERROR_DELETE_FAIL

	case ERROR_SEARCH_QUERY.ErrorName:
		return ERROR_SEARCH_QUERY

	case ERROR_REQUEST_TO_ELASTRIC_SEARCH.ErrorName:
		return ERROR_REQUEST_TO_ELASTRIC_SEARCH

	case ERROR_WHEN_PARSE_RESPONSE_BODY.ErrorName:
		return ERROR_WHEN_PARSE_RESPONSE_BODY

	}
	return ERROR_SERVER;
}

var ERROR_VALIDATE_TOKEN_FAIL = ErrorConvention{
	ErrorName: "validate tolen fail",
	ErrorMessage: "Unauthorized access",
	HttpStatusCode: http.StatusUnauthorized,
}

var ERROR_NO_REFESH_TOKEN = ErrorConvention{
	ErrorName: "no refresh token",
	ErrorMessage: "No refresh token provided",
	HttpStatusCode: http.StatusUnauthorized,
}

var ERROR_GENERATE_TOKEN_FAIL = ErrorConvention{
	ErrorName: "cant generate token",
	ErrorMessage: "Fail to generate token",
	HttpStatusCode: http.StatusInternalServerError,
}

var ERROR_NO_PERMISSION = ErrorConvention{
	ErrorName: "no permisstion",
	ErrorMessage: "No permission granted",
	HttpStatusCode: http.StatusBadRequest,
}

var ERROR_BAD_REQUEST = ErrorConvention{
	ErrorName: "bad request",
	ErrorMessage: "Bad request",
	HttpStatusCode: http.StatusBadRequest,
}

var ERROR_INPUT_INVALID = ErrorConvention{
	ErrorName: "input invalid",
	ErrorMessage: "Input invalid",
	HttpStatusCode: http.StatusBadRequest,
}

var ERROR_SIGNIN_INCORRECT = ErrorConvention{
	ErrorName: "sign in incorrect",
	ErrorMessage: "Username or password is incorrect",
	HttpStatusCode: http.StatusBadRequest,
}

var ERROR_USERNAME_TAKEN = ErrorConvention{
	ErrorName: "username is already taken",
	ErrorMessage: "Username is already taken",
	HttpStatusCode: http.StatusInternalServerError,
}

var ERROR_UPDATE_FAIL = ErrorConvention{
	ErrorName: "update fail",
	ErrorMessage: "Fail to update",
	HttpStatusCode: http.StatusInternalServerError,
}

var ERROR_DELETE_FAIL = ErrorConvention{
	ErrorName: "delete fail",
	ErrorMessage: "Fail to delete",
	HttpStatusCode: http.StatusInternalServerError,
}

var ERROR_SEARCH_QUERY = ErrorConvention{
	ErrorName: "no search",
	ErrorMessage: "No search query present",
	HttpStatusCode: http.StatusInternalServerError,
}

var ERROR_REQUEST_TO_ELASTRIC_SEARCH = ErrorConvention{
	ErrorName: "cant get data from es",
	ErrorMessage: "Server error",
	HttpStatusCode: http.StatusInternalServerError,
}

var ERROR_WHEN_PARSE_RESPONSE_BODY = ErrorConvention{
	ErrorName: "fail to parse es response body",
	ErrorMessage: "Server error",
	HttpStatusCode: http.StatusInternalServerError,
}

var ERROR_SERVER = ErrorConvention{
	ErrorName: "server error",
	ErrorMessage: "Server error",
	HttpStatusCode: http.StatusInternalServerError,
}