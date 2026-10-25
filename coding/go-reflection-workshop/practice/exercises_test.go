package practice

import (
	"bytes"
	"errors"
	"io"
	"math"
	"reflect"
	"sync"
	"testing"
)

func mustError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error")
	}
}

func mustOK(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func TestE01Inspect(t *testing.T) {
	type UserID int64
	for _, tc := range []struct {
		name string
		x    any
		kind reflect.Kind
	}{
		{"nil", nil, reflect.Invalid},
		{"named", UserID(9), reflect.Int64},
		{"pointer", (*UserID)(nil), reflect.Pointer},
		{"slice", []string{}, reflect.Slice},
	} {
		t.Run(tc.name, func(t *testing.T) {
			typ, kind := Inspect(tc.x)
			if typ != reflect.TypeOf(tc.x) || kind != tc.kind {
				t.Fatalf("got (%v, %v)", typ, kind)
			}
		})
	}
}

func TestE02IsNilLike(t *testing.T) {
	for _, tc := range []struct {
		name string
		x    any
		want bool
	}{
		{"nil", nil, true}, {"typed pointer", (*int)(nil), true},
		{"nil slice", []int(nil), true}, {"nil map", map[string]int(nil), true},
		{"nil channel", (chan int)(nil), true}, {"nil func", (func())(nil), true},
		{"empty slice", []int{}, false}, {"empty map", map[string]int{}, false},
		{"zero int", 0, false}, {"zero string", "", false},
		{"zero bool", false, false}, {"pointer", new(int), false},
		{"struct", struct{}{}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsNilLike(tc.x); got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestE03Fields(t *testing.T) {
	type Inner struct{ Promoted int }
	type User struct {
		ID     int    `json:"id,omitempty"`
		Name   string `json:""`
		secret string
		Inner
	}
	u := User{}
	for _, x := range []any{u, &u} {
		got, err := Fields(x)
		mustOK(t, err)
		want := []FieldInfo{
			{"ID", 0, reflect.TypeOf(0), "id,omitempty", true},
			{"Name", 1, reflect.TypeOf(""), "", true},
			{"Inner", 3, reflect.TypeOf(Inner{}), "", false},
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %#v, want %#v", got, want)
		}
	}
}

func TestE03FieldsInvalid(t *testing.T) {
	type User struct{ ID int }
	u := &User{}
	for _, x := range []any{nil, (*User)(nil), 42, &[]int{1}, &u} {
		_, err := Fields(x)
		mustError(t, err)
	}
}

func TestE04SetField(t *testing.T) {
	type Count int
	type Inner struct{ Promoted int }
	type User struct {
		Name   string
		Count  Count
		Writer io.Writer
		Ptr    *int
		secret string
		Inner
	}
	u := User{Name: "before", Ptr: new(int)}
	mustOK(t, SetField(&u, "Name", "after"))
	mustOK(t, SetField(&u, "Count", Count(2)))
	buf := &bytes.Buffer{}
	mustOK(t, SetField(&u, "Writer", buf))
	mustOK(t, SetField(&u, "Ptr", nil))
	if u.Name != "after" || u.Count != 2 || u.Writer != buf || u.Ptr != nil {
		t.Fatalf("bad assignment: %#v", u)
	}
	mustOK(t, SetField(&u, "Writer", nil))
	if u.Writer != nil {
		t.Fatal("nil should clear the interface")
	}
	before := u
	for _, tc := range []struct {
		dst  any
		name string
		src  any
	}{
		{u, "Name", "x"}, {(*User)(nil), "Name", "x"}, {nil, "Name", "x"},
		{&u, "Missing", 1}, {&u, "secret", "x"}, {&u, "Name", 7},
		{&u, "Name", nil}, {&u, "Count", 3}, {&u, "Promoted", 8},
	} {
		mustError(t, SetField(tc.dst, tc.name, tc.src))
		if !reflect.DeepEqual(u, before) {
			t.Fatal("failed assignment modified destination")
		}
	}
}

func TestE05ToMap(t *testing.T) {
	type User struct {
		ID       int    `map:"user_id"`
		Name     string `map:",omitempty"`
		Password string `map:"-"`
		Empty    []int  `map:"empty,omitempty"`
		Nil      []int  `map:"nil,omitempty"`
		private  string
	}
	u := User{ID: 7, Password: "secret", Empty: []int{}}
	got, err := ToMap(&u)
	mustOK(t, err)
	want := map[string]any{"user_id": 7, "empty": []int{}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v", got)
	}
	for _, x := range []any{
		nil, (*User)(nil), 42,
		struct {
			A int `map:"x,omitempty"`
			B int `map:"x,omitempty"`
		}{},
		struct {
			A int `map:"a,unknown"`
		}{},
		struct {
			A int `map:"-,unknown"`
		}{},
	} {
		_, err := ToMap(x)
		mustError(t, err)
	}
}

func TestE06ApplyDefaults(t *testing.T) {
	type Port int16
	type Config struct {
		Host    string `default:"localhost"`
		Port    Port   `default:"8080"`
		Debug   bool   `default:"true"`
		Empty   string `default:""`
		Retries int    `default:"3"`
		secret  int    `default:"bad"`
	}
	c := Config{Host: "custom", Retries: 9}
	mustOK(t, ApplyDefaults(&c))
	if c.Host != "custom" || c.Port != 8080 || !c.Debug || c.Retries != 9 || c.secret != 0 {
		t.Fatalf("bad defaults: %#v", c)
	}
	for _, dst := range []any{nil, c, (*Config)(nil), new(int)} {
		mustError(t, ApplyDefaults(dst))
	}
	t.Run("atomic overflow", func(t *testing.T) {
		c := struct {
			Name string `default:"changed"`
			Tiny int8   `default:"128"`
		}{}
		mustError(t, ApplyDefaults(&c))
		if c.Name != "" || c.Tiny != 0 {
			t.Fatal("parse failure must leave all fields unchanged")
		}
	})
	for _, dst := range []any{
		&struct {
			Flag bool `default:"maybe"`
		}{},
		&struct {
			N int `default:"abc"`
		}{},
		&struct {
			N []int `default:"1"`
		}{},
	} {
		mustError(t, ApplyDefaults(dst))
	}
}

func TestE07Validate(t *testing.T) {
	type Form struct {
		Name   string `validate:"required,min=2"`
		Age    int    `validate:"min=18"`
		secret string `validate:"unknown"`
	}
	got, err := Validate(Form{})
	mustOK(t, err)
	want := []Violation{{"Name", "required"}, {"Name", "min=2"}, {"Age", "min=18"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v", got)
	}
	got, err = Validate(&Form{Name: "小明", Age: 18})
	mustOK(t, err)
	if len(got) != 0 {
		t.Fatalf("Unicode length or valid form failed: %#v", got)
	}
	got, err = Validate(struct {
		Items []int `validate:"required"`
	}{Items: []int{}})
	mustOK(t, err)
	if len(got) != 0 {
		t.Fatal("required uses Go IsZero semantics in this exercise")
	}
	for _, x := range []any{
		nil, 1, (*Form)(nil),
		struct {
			A int `validate:"max=3"`
		}{},
		struct {
			A int `validate:"min=oops"`
		}{},
		struct {
			A bool `validate:"min=1"`
		}{},
		struct {
			A string `validate:"min=-1"`
		}{},
	} {
		_, err := Validate(x)
		mustError(t, err)
	}
}

func TestE08Invoke(t *testing.T) {
	wantErr := errors.New("business error")
	got, err := Invoke(func(a, b int) (int, error) { return a + b, wantErr }, 2, 3)
	mustOK(t, err)
	if len(got) != 2 || got[0] != 5 || got[1] != wantErr {
		t.Fatalf("got %#v", got)
	}
	got, err = Invoke(func(x *int, y any) bool { return x == nil && y == nil }, nil, nil)
	mustOK(t, err)
	if len(got) != 1 || got[0] != true {
		t.Fatalf("nil argument failed: %#v", got)
	}
	buf := &bytes.Buffer{}
	_, err = Invoke(func(w io.Writer) { _, _ = w.Write([]byte("ok")) }, buf)
	mustOK(t, err)
	if buf.String() != "ok" {
		t.Fatal("concrete value should be assignable to interface parameter")
	}
	var nilFn func(int)
	for _, tc := range []struct {
		fn   any
		args []any
	}{
		{nil, nil}, {1, nil}, {nilFn, []any{1}},
		{func(int) {}, nil}, {func(int) {}, []any{int64(1)}},
		{func(int) {}, []any{nil}}, {func(...int) {}, []any{1}},
	} {
		_, err := Invoke(tc.fn, tc.args...)
		mustError(t, err)
	}
	t.Run("business panic propagates", func(t *testing.T) {
		defer func() {
			if got := recover(); got != "boom" {
				t.Fatalf("panic = %#v, want boom", got)
			}
		}()
		_, _ = Invoke(func() { panic("boom") })
	})
}

func TestE09IncrementMap(t *testing.T) {
	type Key string
	type Count int8
	m := map[Key]Count{"a": 4}
	mustOK(t, IncrementMap(m, "a", 2))
	mustOK(t, IncrementMap(m, "new", -2))
	if m["a"] != 6 || m["new"] != -2 {
		t.Fatalf("got %#v", m)
	}
	mustError(t, IncrementMap(m, "a", 122))
	if m["a"] != 6 {
		t.Fatal("overflow modified map")
	}
	for _, x := range []any{nil, map[string]int(nil), map[int]int{}, map[string]uint{}, &m} {
		mustError(t, IncrementMap(x, "x", 1))
	}
	for _, tc := range []struct{ old, delta int64 }{
		{math.MaxInt64, 1}, {math.MinInt64, -1}, {1, math.MaxInt64}, {-1, math.MinInt64},
	} {
		m := map[string]int64{"x": tc.old}
		mustError(t, IncrementMap(m, "x", tc.delta))
		if m["x"] != tc.old {
			t.Fatal("int64 overflow modified map")
		}
	}
	m64 := map[string]int64{}
	mustOK(t, IncrementMap(m64, "x", math.MinInt64))
	if m64["x"] != math.MinInt64 {
		t.Fatal("valid boundary rejected")
	}
}

func TestE10FieldCache(t *testing.T) {
	type User struct {
		Name string `json:"name"`
		age  int
	}
	var c FieldCache
	typ := reflect.TypeOf(User{})
	fields, err := c.Metadata(typ)
	mustOK(t, err)
	if len(fields) != 1 || fields[0].Name != "Name" {
		t.Fatalf("got %#v", fields)
	}
	fields[0].Name = "corrupted"
	again, err := c.Metadata(reflect.PointerTo(typ))
	mustOK(t, err)
	if again[0].Name != "Name" {
		t.Fatal("caller modified shared cache storage")
	}
	var wg sync.WaitGroup
	for i := 0; i < 32; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := 0; n < 20; n++ {
				got, err := c.Metadata(typ)
				if err != nil || len(got) != 1 || got[0].Name != "Name" {
					t.Errorf("concurrent metadata: %#v, %v", got, err)
					return
				}
				got[0].Name = "local change"
			}
		}()
	}
	wg.Wait()
	for _, typ := range []reflect.Type{nil, reflect.TypeOf(42), reflect.PointerTo(reflect.PointerTo(typ))} {
		_, err := c.Metadata(typ)
		mustError(t, err)
	}
}

func TestE11ScanRow(t *testing.T) {
	type Row struct {
		ID     int64   `db:"user_id"`
		Name   string  `db:"name"`
		Note   *string `db:"note"`
		Skip   int     `db:"-"`
		secret int
	}
	row := Row{Skip: 8}
	mustOK(t, ScanRow(&row, []string{"name", "user_id", "note"}, []any{"小明", int64(7), nil}))
	if row.ID != 7 || row.Name != "小明" || row.Note != nil || row.Skip != 8 {
		t.Fatalf("got %#v", row)
	}
	before := row
	for _, tc := range []struct {
		columns []string
		values  []any
	}{
		{[]string{"name", "user_id"}, []any{"changed", 3}},
		{[]string{"name", "missing"}, []any{"changed", 3}},
		{[]string{"name", "name"}, []any{"a", "b"}},
		{[]string{"name"}, []any{nil}},
		{[]string{"Skip"}, []any{1}},
		{[]string{"secret"}, []any{1}},
		{[]string{"name"}, nil},
	} {
		mustError(t, ScanRow(&row, tc.columns, tc.values))
		if !reflect.DeepEqual(row, before) {
			t.Fatal("failed scan must not partially write")
		}
	}
	for _, dst := range []any{nil, row, (*Row)(nil), new(int)} {
		mustError(t, ScanRow(dst, nil, nil))
	}
	duplicate := struct {
		A int `db:"x"`
		B int `db:"x"`
	}{}
	mustError(t, ScanRow(&duplicate, nil, nil))
}
