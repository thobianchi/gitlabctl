package sdk

import (
	"reflect"
	"testing"
)

func TestGetHome(t *testing.T) {
	home, _ := getHome()
	if reflect.TypeOf(home).Kind() != reflect.String {
		t.Errorf("GetHome does not return string")
	}
}
