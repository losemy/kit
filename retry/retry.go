package retry

import (
	"time"
)

// WithRetry retry util job succeeded or retry count limit exceed
func WithRetry(retry int, job func() error) error {
	var retries int
	for {
		retries++
		if err := job(); err != nil {
			if retry == -1 || (retry > 0 && retries <= retry) {
				continue
			}
			return err
		}
		return nil
	}
}

// WithRetryDelay retry util job succeeded or retry count limit exceed
func WithRetryDelay(retry int, delay time.Duration, job func() error) error {
	var retries int
	for {
		retries++
		if err := job(); err != nil {
			if retry == -1 || (retry > 0 && retries <= retry) {
				time.Sleep(delay)
				continue
			}
			return err
		}
		return nil
	}
}
