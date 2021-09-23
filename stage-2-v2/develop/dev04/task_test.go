package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_makeSet(t *testing.T) {
	data := []string{
		"листок",
		"тяпка",
		"пятка",
		"пятак",
		"Пятак",
		"пятак",
		"слиток",
		"столик",
	}

	Out := map[string][]string{
		"листок": {"слиток", "столик"},
		"тяпка":  {"пятак", "пятка"},
	}

	rez := find(data)
	assert.Equal(t, Out, rez)
}
