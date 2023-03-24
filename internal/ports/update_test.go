package ports_test

import (
	"context"
	"errors"
	"testing"

	"github.com/fakovacic/ports/internal/ports"
	"github.com/fakovacic/ports/internal/ports/mocks"
	"github.com/matryer/is"
)

func TestUpdate(t *testing.T) {
	cases := []struct {
		it string

		id    string
		model *ports.Port

		// Store
		portUpdateInputID    string
		portUpdateInputModel *ports.Port
		portUpdateError      error

		expectedError  string
		expectedResult *ports.Port
	}{
		{
			it: "it update port",

			id: "mock-id",
			model: &ports.Port{
				Name: "mock-name",
			},

			portUpdateInputID: "mock-id",
			portUpdateInputModel: &ports.Port{
				Name: "mock-name",
			},

			expectedResult: &ports.Port{
				Name: "mock-name",
			},
		},
		{
			it: "it return error on port update",

			id: "mock-id",
			model: &ports.Port{
				Name: "mock-name",
			},

			portUpdateInputID: "mock-id",
			portUpdateInputModel: &ports.Port{
				Name: "mock-name",
			},
			portUpdateError: errors.New("mock-error"),

			expectedError: "update port: mock-error",
		},
		{
			it: "it return error on validation, name empty",

			id: "mock-id",
			model: &ports.Port{
				Name: "",
			},

			expectedError: "name is required",
		},
		{
			it:            "it return error on empty id",
			expectedError: "id is empty",
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			store := &mocks.StoreMock{
				UpdateFunc: func(ctx context.Context, id string, model *ports.Port) error {
					checkIs.Equal(id, tc.portUpdateInputID)
					checkIs.Equal(model, tc.portUpdateInputModel)

					return tc.portUpdateError
				},
			}

			service := ports.New(
				ports.NewConfig(""),
				store,
			)

			res, err := service.Update(context.Background(), tc.id, tc.model)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
			checkIs.Equal(res, tc.expectedResult)
		})
	}
}
