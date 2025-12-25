package utils

import (
	"bisecur/cli"
)

func Retry(retryCount int, f func() error) error {
	var err error

	for i := 0; i < retryCount; i++ {
		err = f()
		if err == nil {
			break
		}
		cli.Log.Warnf("Retriable error: %v", err)
	}

	return err
}
