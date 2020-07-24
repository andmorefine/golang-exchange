package models

import (
	"testing"
)

func TestCreate(t *testing.T) {
	user := User{
		UUID:       1,
		Name:       "test",
		DeleteFlag: false,
	}
	user.Create()
}

func TestUpdate(t *testing.T) {
	user := User{
		UUID:       1,
		Name:       "test_update",
		DeleteFlag: true,
	}
	user.Update()
}
