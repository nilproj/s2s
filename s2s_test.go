package s2s_test

import (
	"testing"

	"github.com/nilproj/s2s"
)

type User struct {
	Name string
	Age  int
}

type Employee struct {
	Name string
	Age  int
	Role string
}

func TestCopy(t *testing.T) {
	u := User{Name: "john", Age: 30}
	e := Employee{Role: "admin"}
	want := Employee{Name: "john", Age: 30, Role: "admin"}

	s2s.Copy(&e, &u)

	if want != e {
		t.Errorf("FAIL\nexp: %#v\ngot: %#v", want, e)
	}
}
