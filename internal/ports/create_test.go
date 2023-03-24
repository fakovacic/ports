package ports_test

import (
	"context"
	"errors"
	"testing"

	"github.com/fakovacic/ports/internal/ports"
	"github.com/fakovacic/ports/internal/ports/mocks"
	"github.com/matryer/is"
)

func TestCreate(t *testing.T) {
	cases := []struct {
		it string

		model *ports.Port

		// Store
		portCreateInput *ports.Port
		portCreateError error

		expectedError  string
		expectedResult *ports.Port
	}{
		{
			it: "it create port",

			model: &ports.Port{
				Name: "mock-name",
			},

			portCreateInput: &ports.Port{
				Name: "mock-name",
			},

			expectedResult: &ports.Port{
				Name: "mock-name",
			},
		},
		{
			it: "it return error on port create",

			model: &ports.Port{
				Name: "mock-name",
			},

			portCreateInput: &ports.Port{
				Name: "mock-name",
			},
			portCreateError: errors.New("mock-error"),

			expectedError: "create port: mock-error",
		},
		{
			it: "it return error on validation, name empty",

			model: &ports.Port{
				Name: "",
			},

			expectedError: "name is required",
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			store := &mocks.StoreMock{
				CreateFunc: func(ctx context.Context, model *ports.Port) error {
					checkIs.Equal(model, tc.portCreateInput)

					return tc.portCreateError
				},
			}

			service := ports.New(
				ports.NewConfig(""),
				store,
			)

			res, err := service.Create(context.Background(), tc.model)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
			checkIs.Equal(res, tc.expectedResult)
		})
	}
}
