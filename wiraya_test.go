package gowiraya

import (
	"testing"
)

func TestGetBearerToken(t *testing.T) {
	c, err := NewWirayaClient("insert token here", nil)
	if err != nil {
		t.Error(err)
	}

	b, err := c.getBearerToken()
	if err != nil {
		t.Error(err)
	}
	t.Log(b)
}
