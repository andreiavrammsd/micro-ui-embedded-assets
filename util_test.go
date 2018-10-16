package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLimit(t *testing.T) {
	assert.Equal(t, 2, getLimit("2"))
	assert.Equal(t, defaultLimit, getLimit(strconv.Itoa(maxLimit+1)))
	assert.Equal(t, defaultLimit, getLimit("-1"))
}

func TestGetPaginationUrl(t *testing.T) {
	assert.Equal(t, "?limit=100&offset=100", getPaginationURL("0", "100", "next"))
	assert.Equal(t, "?limit=100&offset=100", getPaginationURL("200", "100", "prev"))
	assert.Equal(t, "", getPaginationURL("0", "100", "prev"))
}

func TestGetIntValue(t *testing.T) {
	assert.Equal(t, 1, getIntValue("1"))
	assert.Equal(t, 0, getIntValue("0"))
	assert.Equal(t, 0, getIntValue(" "))
	assert.Equal(t, 0, getIntValue(""))
	assert.Equal(t, 0, getIntValue("any string"))
}
