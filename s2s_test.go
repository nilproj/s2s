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

func TestCopyStruct(t *testing.T) {
	name, age, email, addr, birth := "john", 30, "foo@bar.com", "somewhere", time.Now()
	u := User{Name: name, Age: age, Phone: "0123456789", Email: &email, Address: &addr, Birthdate: MyTime{birth}}
	e := Employee{Role: "admin"}
	want := Employee{Name: name, Age: &age, Email: email, Address: &addr, Role: "admin", Birthdate: birth}

	s2s.Copy(&e, &u)

	if !reflect.DeepEqual(want, e) {
		t.Errorf("FAIL\nexp: %#v\ngot: %#v", want, e)
	}
}

func TestCopyStructToSlice(t *testing.T) {
	now := time.Now()
	u := User{Name: "john", Birthdate: MyTime{now}}
	var employees []Employee
	want := []Employee{{Name: "john", Birthdate: now}}

	s2s.Copy(&employees, &u)

	if !reflect.DeepEqual(want, employees) {
		t.Errorf("FAIL\nexp: %#v\ngot: %#v", want, employees)
	}
}

func TestCopySliceToSlice(t *testing.T) {
	now := time.Now()
	users := []User{{Name: "john", Birthdate: MyTime{now}}, {Name: "will"}}
	var employees []Employee
	want := []Employee{{Name: "john", Birthdate: now}, {Name: "will"}}

	s2s.Copy(&employees, &users)

	if !reflect.DeepEqual(want, employees) {
		t.Errorf("FAIL\nexp: %#v\ngot: %#v", want, employees)
	}
}
