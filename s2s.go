package s2s

import "reflect"

func Copy(dst, src any) {
	dstEl := reflect.ValueOf(dst).Elem()
	srcEl := reflect.ValueOf(src).Elem()

	for f, v := range srcEl.Fields() {
		dstEl.FieldByName(f.Name).Set(v)
	}
}
