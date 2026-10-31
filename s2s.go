package s2s

import (
	"reflect"
)

func Copy(dst, src any) {
	dstEl := reflect.ValueOf(dst).Elem()
	srcEl := reflect.ValueOf(src).Elem()
	dk, sk := dstEl.Kind(), srcEl.Kind()

	if dk == reflect.Struct && sk == reflect.Struct {
		copyStruct(dstEl, srcEl)
	} else if dk == reflect.Slice && sk == reflect.Struct {
		copyStructToSlice(dstEl, srcEl)
	} else if dk == reflect.Slice && sk == reflect.Slice {
		for i := range srcEl.Len() {
			copyStructToSlice(dstEl, srcEl.Index(i))
		}
	}
}

func copyStructToSlice(dstEl, srcEl reflect.Value) {
	iv := reflect.New(dstEl.Type().Elem())
	ie := iv.Elem()
	copyStruct(ie, srcEl)
	dstEl.Set(reflect.Append(dstEl, ie))
}

func copyStruct(dstEl, srcEl reflect.Value) {
	for sf, sv := range srcEl.Fields() {
		df := dstEl.FieldByName(sf.Name)
		if !df.CanSet() || !sv.IsValid() || sv.IsZero() {
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
