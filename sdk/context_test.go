package sdk

import (
	"reflect"
	"testing"
)

func TestGetHome(t *testing.T) {
	home := getHome()
	if reflect.TypeOf(home).Kind() != reflect.String {
		t.Errorf("GetHome does not return string")
	}
}
