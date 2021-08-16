package customTime

import (
	"errors"
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

var ErrWrongHost = errors.New("wrong host")

const defaultHost = "0.beevik-ntp.pool.ntp.org"

type CustomTime interface {
	SetHost(string) error
	Time() (*time.Time, error)
}

type customTime struct {
	host string
}

func New() CustomTime {
	return &customTime{
		host: defaultHost,
	}
}

func (ct *customTime) SetHost(host string) error {
	if host == "" {
		return fmt.Errorf("error #001: %s", ErrWrongHost)
	}
	ct.host = host
	return nil
}

func (ct *customTime) Time() (*time.Time, error) {
	time, err := ntp.Time(ct.host)
	if err != nil {
		return nil, fmt.Errorf("error #002: %s", err)
	}
	return &time, nil
}
