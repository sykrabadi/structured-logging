package service

import (
	"context"
	"math/rand/v2"

	"github.com/pkg/errors"
)

func (s *Service) Home(_ context.Context) error {
	errRand := rand.IntN(5)

	// generate dummy error
	if errRand < 10 {
		return errors.Wrapf(errors.New("encountered error"), "error with rand error: %v", errRand)
	}

	return nil
}
