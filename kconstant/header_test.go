package kconstant

import (
	"testing"
)

func TestHeader(t *testing.T) {
	Prefix = "name"
	t.Log(Timestamp)
	t.Log(Sign)
}
