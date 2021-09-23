package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	rTime := time.Now()

	ntpTime := NtpTime{
		Host: "0.beevik-ntp.pool.ntp.org",
	}
	testTime, _ := ntpTime.Get()

	assert.Equal(t, rTime.Local().Hour(), testTime.Local().Hour())
	assert.Equal(t, rTime.Local().Minute(), testTime.Local().Minute())
	assert.Equal(t, rTime.Local().Second(), testTime.Local().Second())
}
