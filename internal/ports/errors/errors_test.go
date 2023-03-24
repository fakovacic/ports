package errors_test

import (
	goErrors "errors"
	"net/http"
	"testing"

	"github.com/fakovacic/ports/internal/ports/errors"
	"github.com/matryer/is"
)

func TestStatus(t *testing.T) {
	err := errors.New("wrap-error")
	cases := []struct {
		it             string
		errorInput     error
		expectedStatus int
		expectedError  string
	}{
		{
			it:             "it returns wrap error with InternalServer status",
			errorInput:     errors.Wrapf(err, "mock-error"),
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "mock-error: wrap-error",
		},
		{
			it:             "it returns error with InternalServer status",
			errorInput:     errors.Internal("mock-error"),
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "mock-error",
		},
		{
			it:             "it returns wrap error with InternalServer status",
			errorInput:     errors.InternalWrap(err, "mock-error"),
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "mock-error: wrap-error",
		},
		{
			it:             "it returns error with BadRequest status",
			errorInput:     errors.BadRequest("mock-error"),
			expectedStatus: http.StatusBadRequest,
			expectedError:  "mock-error",
		},
		{
			it:             "it returns wrap error with BadRequest status",
			errorInput:     errors.BadRequestWrap(err, "mock-error"),
			expectedStatus: http.StatusBadRequest,
			expectedError:  "mock-error: wrap-error",
		},
		{
			it:             "it returns error with NotFound status",
			errorInput:     errors.NotFound("mock-error"),
			expectedStatus: http.StatusNotFound,
			expectedError:  "mock-error",
		},
		{
			it:             "it returns wrap error with NotFound status",
			errorInput:     errors.NotFoundWrap(err, "mock-error"),
			expectedStatus: http.StatusNotFound,
			expectedError:  "mock-error: wrap-error",
		},
		{
			it:             "it returns error with Unauthorized status",
			errorInput:     errors.Unauthorized("mock-error"),
			expectedStatus: http.StatusUnauthorized,
			expectedError:  "mock-error",
		},
		{
			it:             "it returns wrap error with Unauthorized status",
			errorInput:     errors.UnauthorizedWrap(err, "mock-error"),
			expectedStatus: http.StatusUnauthorized,
			expectedError:  "mock-error: wrap-error",
		},
		{
			it:             "it returns error with MethodNotAllowed status",
			errorInput:     errors.MethodNotAllowed("mock-error"),
			expectedStatus: http.StatusMethodNotAllowed,
			expectedError:  "mock-error",
		},
		{
			it:             "it returns error with MethodNotAllowed status",
			errorInput:     errors.MethodNotAllowedWrap(err, "mock-error"),
			expectedStatus: http.StatusMethodNotAllowed,
			expectedError:  "mock-error: wrap-error",
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			check := is.New(t)

			var internalError errors.Error

			ok := goErrors.As(tc.errorInput, &internalError)
			if !ok {
				t.Error("errorInput is not errors.Error")
			}

			status := internalError.GetStatus()
			er := internalError.Error()
			check.Equal(status, tc.expectedStatus)
			check.Equal(er, tc.expectedError)
		})
	}
}
