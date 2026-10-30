package s2s_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/nilproj/s2s"
)

type MyTime struct {
	time.Time
}

type User struct {
	Name      string
	Age       int
	Phone     string
	Email     *string
	Address   *string
	Birthdate MyTime
}

type Employee struct {
	Name      string
	Age       *int
	Email     string
	Address   *string
	Role      string
	Birthdate time.Time
}

func TestCopy(t *testing.T) {
	name, age, email, addr, birth := "john", 30, "foo@bar.com", "somewhere", time.Now()
	u := User{Name: name, Age: age, Phone: "0123456789", Email: &email, Address: &addr, Birthdate: MyTime{birth}}
	e := Employee{Role: "admin"}
	want := Employee{Name: name, Age: &age, Email: email, Address: &addr, Role: "admin", Birthdate: birth}

	s2s.Copy(&e, &u)

	if !reflect.DeepEqual(want, e) {
		t.Errorf("FAIL\nexp: %#v\ngot: %#v", want, e)
	}
}
