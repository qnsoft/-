package openurl

import "testing"

func TestOpen(t *testing.T) {
	if err := Open("http://192.168.1.132:8199"); err != nil {
		t.Error(err)
	}
}
