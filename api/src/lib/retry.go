package lib

import (
	"context"
	"errors"
	"time"
)

type RetryOptions struct {
	Retries int
	Delay   time.Duration
	Prober  func(err error) bool
}

func Retry(ctx context.Context, callback func() error, options RetryOptions) error {
	var err error

	retries, prober := options.Retries, options.Prober

	for r := 0; ; r++ {
		delay := options.Delay
		if r == 0 {
			delay = 0
		}
		select {
		case <-ctx.Done():
			if err != nil {
				return err
			}
			return errors.New("ctx done")
		case <-time.After(delay):
			err = callback()
			if err == nil || (retries != -1 && r >= retries) || (prober != nil && !prober(err)) {
				return err
			}

			continue
		}
	}
}
