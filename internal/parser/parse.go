package parser

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/fakovacic/ports/internal/errors"
	"github.com/fakovacic/ports/internal/ports"
)

func (s *service) Parse(ctx context.Context, input io.Reader) error {
	decoder := json.NewDecoder(input)

	_, err := decoder.Token()
	if err != nil {
		return errors.Wrapf(err, "failed to get token")
	}

	for decoder.More() {
		key, err := decoder.Token()
		if err != nil {
			return errors.Wrapf(err, "failed read key")
		}

		var port ports.Port

		err = decoder.Decode(&port)
		if err != nil {
			return errors.Wrapf(err, "decode")
		}

		port.ID = fmt.Sprintf("%s", key)

		_, err = s.service.Create(ctx, &port)
		if err != nil {
			return errors.Wrapf(err, "failed to create port")
		}
	}

	return nil
}
