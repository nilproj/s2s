package s2s

import (
	"reflect"
)

func Copy(dst, src any) {
	dstEl := reflect.ValueOf(dst).Elem()
	srcEl := reflect.ValueOf(src).Elem()

	for sf, sv := range srcEl.Fields() {
		df := dstEl.FieldByName(sf.Name)
		if !df.CanSet() {
			continue
		}

		switch df.Kind() {
		case reflect.Pointer:
			sv = ptrOrNew(sv)
		default:
			sv = reflect.Indirect(sv)
		}

		setField(df, sv)
	}
}

func setField(f reflect.Value, v reflect.Value) bool {
	ft, vt := f.Type(), v.Type()

	if ft == vt {
		f.Set(v)
		return true
	}

	for ef, ev := range v.Fields() {
		if ef.Anonymous && setField(f, ev) {
			return true
		}
	}

	return false
}

func ptrOrNew(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Pointer {
		return v
	}
	return newPtr(v)
}

func newPtr(v reflect.Value) reflect.Value {
	p := reflect.New(v.Type())
	p.Elem().Set(v)
	return p
}
