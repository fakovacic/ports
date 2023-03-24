package http_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fakovacic/ports/internal/ports"
	"github.com/fakovacic/ports/internal/ports/errors"
	handlers "github.com/fakovacic/ports/internal/ports/handlers/http"
	"github.com/fakovacic/ports/internal/ports/mocks"
	"github.com/julienschmidt/httprouter"
	"github.com/matryer/is"
)

func TestCreate(t *testing.T) {
	cases := []struct {
		it string

		requestBody string

		createInput    *ports.Port
		createResponse *ports.Port
		createError    error

		expectedError  string
		expectedResult string
		expectedStatus int
	}{
		{
			it:          "it create port",
			requestBody: `{"port":{"name":"mock-name"}}`,

			createInput: &ports.Port{
				Name: "mock-name",
			},

			createResponse: &ports.Port{
				Name: "mock-name",
			},

			expectedResult: `{"port":{"name":"mock-name","city":"","country":"","province":"","timezone":"","code":"","alias":null,"regions":null,"coordinates":null,"unlocs":null}}`,
			expectedStatus: http.StatusOK,
		},
		{
			it:          "it return error on service Create",
			requestBody: `{"port":{"name":"mock-name"}}`,

			createInput: &ports.Port{
				Name: "mock-name",
			},

			createError:    errors.Wrap("mock-error"),
			expectedResult: `{"message":"mock-error","status":500}`,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			service := &mocks.ServiceMock{
				CreateFunc: func(contextMoqParam context.Context, port *ports.Port) (*ports.Port, error) {
					checkIs.Equal(port, tc.createInput)

					return tc.createResponse, tc.createError
				},
			}

			req, err := http.NewRequest(
				http.MethodPost,
				"/",
				bytes.NewReader([]byte(tc.requestBody)),
			)

			req = req.WithContext(context.Background())
			if err != nil {
				t.Fatal(err)
			}

			h := handlers.New(
				ports.NewConfig(""),
				service,
			)

			router := httprouter.New()
			rr := httptest.NewRecorder()

			router.POST("/", h.Create)
			router.ServeHTTP(rr, req)

			checkIs.Equal(rr.Body.String(), tc.expectedResult)
			checkIs.Equal(rr.Code, tc.expectedStatus)
		})
	}
}
