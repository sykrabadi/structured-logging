package service

import (
	"context"
	"math/rand/v2"

	"github.com/pkg/errors"
)

func (s *Service) Home(_ context.Context) error {
	errRand := rand.IntN(2)

	// generate dummy error
	if errRand < 2 {
		return errors.Wrap(errors.New("sengaja error"), "kena error mang")
	}

	return nil
}
