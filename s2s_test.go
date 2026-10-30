package s2s_test

import (
	"reflect"
	"testing"

	"github.com/nilproj/s2s"
)

type User struct {
	Name    string
	Age     int
	Phone   string
	Email   *string
	Address *string
}

type Employee struct {
	Name    string
	Age     *int
	Email   string
	Address *string
	Role    string
}

func TestCopy(t *testing.T) {
	name, age, email, addr := "john", 30, "foo@bar.com", "somewhere"
	u := User{Name: name, Age: age, Phone: "0123456789", Email: &email, Address: &addr}
	e := Employee{Role: "admin"}
	want := Employee{Name: name, Age: &age, Email: email, Address: &addr, Role: "admin"}

	s2s.Copy(&e, &u)

	if !reflect.DeepEqual(want, e) {
		t.Errorf("FAIL\nexp: %#v\ngot: %#v", want, e)
	}
}
