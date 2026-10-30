package s2s

import "reflect"

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
			df.Set(ptrOrNew(sv))
		default:
			df.Set(reflect.Indirect(sv))
		}
	}
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
