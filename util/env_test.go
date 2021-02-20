package util

import "testing"

func TestGetEnv(t *testing.T) {
	env := GetEnv()
	if env != "test" {
		t.Errorf("Get Env Failed")
	}
}
