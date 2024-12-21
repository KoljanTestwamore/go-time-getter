package timegetter

import (
	"time"

	"github.com/beevik/ntp"
)

func GetTime() (*time.Time, error) {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		return nil, err
	}

	return &time, nil
}