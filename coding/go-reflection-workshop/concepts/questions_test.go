package concepts

import (
	"reflect"
	"testing"
)

func TestQ10ValueDependentConversion(t *testing.T) {
	target := reflect.TypeOf([2]int{})
	short := reflect.ValueOf([]int{1})
	if !short.Type().ConvertibleTo(target) || short.CanConvert(target) {
		t.Fatal("type convertibility and current-value convertibility should differ")
	}
	full := reflect.ValueOf([]int{1, 2})
	if !full.CanConvert(target) || full.Convert(target).Interface() != [2]int{1, 2} {
		t.Fatal("a slice with enough elements should convert")
	}
}

func TestQ15IntegerTruncationAndBounds(t *testing.T) {
	var n int8
	v := reflect.ValueOf(&n).Elem()
	v.SetInt(128)
	if n != -128 {
		t.Fatalf("truncated value = %d", n)
	}
	for _, tc := range []struct {
		n        int64
		overflow bool
	}{{-129, true}, {-128, false}, {127, false}, {128, true}} {
		if v.OverflowInt(tc.n) != tc.overflow {
			t.Fatalf("OverflowInt(%d)", tc.n)
		}
	}
}

func TestQ07AppendWritesBack(t *testing.T) {
	s := []int{1}
	v := reflect.ValueOf(&s).Elem()
	v.Set(reflect.Append(v, reflect.ValueOf(2)))
	if !reflect.DeepEqual(s, []int{1, 2}) {
		t.Fatalf("got %v", s)
	}
}

func TestQ16MapStructCopy(t *testing.T) {
	type User struct{ Name string }
	m := map[string]User{"u1": {Name: "before"}}
	mv := reflect.ValueOf(m)
	k := reflect.ValueOf("u1")
	old := mv.MapIndex(k)
	if old.CanSet() || old.FieldByName("Name").CanSet() {
		t.Fatal("map element must not be settable")
	}
	copyValue := reflect.New(old.Type()).Elem()
	copyValue.Set(old)
	copyValue.FieldByName("Name").SetString("after")
	mv.SetMapIndex(k, copyValue)
	if m["u1"].Name != "after" {
		t.Fatal("modified copy was not written back")
	}
}
