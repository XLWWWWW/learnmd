package concepts

import (
	"fmt"
	"io"
	"reflect"
)

type userID int64
type counter struct{ N int }

func (c counter) Read() int  { return c.N }
func (c *counter) Add(n int) { c.N += n }

func Example_typeAndKind() {
	v := reflect.ValueOf(userID(7))
	fmt.Println(v.Type() == reflect.TypeOf(int64(0)), v.Kind() == reflect.Int64)
	fmt.Println(v.Int())
	// Output:
	// false true
	// 7
}

func Example_dynamicType() {
	var x any = int64(7)
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.ValueOf(&x).Elem().Kind())
	fmt.Println(reflect.ValueOf(&x).Elem().Elem().Kind())
	fmt.Println(reflect.TypeFor[io.Writer]() == reflect.TypeOf((*io.Writer)(nil)).Elem())
	// Output:
	// int64
	// interface
	// int64
	// true
}

func Example_settable() {
	x := 10
	value := reflect.ValueOf(x)
	pointer := reflect.ValueOf(&x)
	target := pointer.Elem()
	fmt.Println(value.CanAddr(), value.CanSet())
	fmt.Println(pointer.CanAddr(), pointer.CanSet())
	fmt.Println(target.CanAddr(), target.CanSet())
	target.SetInt(20)
	fmt.Println(x)
	u := struct{ secret int }{secret: 7}
	f := reflect.ValueOf(&u).Elem().Field(0)
	fmt.Println(f.CanAddr(), f.CanSet(), f.CanInterface(), f.Int())
	// Output:
	// false false
	// false false
	// true true
	// 20
	// true false false 7
}

func Example_nilStates() {
	var p *int
	var x any = p
	fmt.Println(x == nil, reflect.ValueOf(x).IsValid(), reflect.ValueOf(x).IsNil())
	fmt.Println(reflect.ValueOf(nil).IsValid(), reflect.ValueOf(nil).Kind())
	fmt.Println(reflect.Zero(reflect.TypeOf(0)).IsValid(), reflect.Zero(reflect.TypeOf(0)).IsZero())
	fmt.Println(reflect.ValueOf([]int{}).IsZero(), reflect.ValueOf([]int(nil)).IsZero())
	// Output:
	// false true true
	// false invalid
	// true true
	// false true
}

func Example_collectionMutation() {
	s := []int{1}
	a := [1]int{1}
	fmt.Println(reflect.ValueOf(s).CanSet(), reflect.ValueOf(s).Index(0).CanSet())
	fmt.Println(reflect.ValueOf(a).Index(0).CanSet())
	reflect.ValueOf(s).Index(0).SetInt(9)
	reflect.ValueOf(&a).Elem().Index(0).SetInt(8)
	fmt.Println(s, a)
	// Output:
	// false true
	// false
	// [9] [8]
}

func Example_mapElement() {
	m := map[string]int{"hits": 1}
	v := reflect.ValueOf(m)
	key := reflect.ValueOf("hits")
	fmt.Println(v.CanSet(), v.MapIndex(key).CanSet())
	v.SetMapIndex(key, reflect.ValueOf(2))
	fmt.Println(m["hits"])
	v.SetMapIndex(key, reflect.Value{})
	fmt.Println(len(m), v.MapIndex(key).IsValid())
	// Output:
	// false false
	// 2
	// 0 false
}

func Example_methodSets() {
	c := counter{N: 1}
	fmt.Println(reflect.ValueOf(c).MethodByName("Add").IsValid())
	method := reflect.ValueOf(&c).MethodByName("Add")
	fmt.Println(method.IsValid(), method.Type().NumIn())
	method.Call([]reflect.Value{reflect.ValueOf(2)})
	fmt.Println(c.N)
	// Output:
	// false
	// true 1
	// 3
}

func Example_tags() {
	t := reflect.TypeOf(struct {
		A string `json:"a,omitempty"`
		B string `json:""`
		C string
	}{})
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		raw, ok := f.Tag.Lookup("json")
		fmt.Printf("%s %q %v\n", f.Name, raw, ok)
	}
	// Output:
	// A "a,omitempty" true
	// B "" true
	// C "" false
}

func Example_newAndZero() {
	t := reflect.TypeOf(int(0))
	p := reflect.New(t)
	fmt.Println(p.Kind(), p.Elem().CanSet())
	p.Elem().SetInt(42)
	fmt.Println(p.Elem().Interface())
	z := reflect.Zero(t)
	fmt.Println(z.IsValid(), z.CanSet(), z.Int())
	// Output:
	// ptr true
	// 42
	// true false 0
}

func Example_assignmentAndConversion() {
	source := reflect.ValueOf(int64(257))
	target := reflect.TypeOf(int8(0))
	fmt.Println(source.Type().AssignableTo(target), source.CanConvert(target))
	fmt.Println(source.Convert(target).Int())
	dst := reflect.New(target).Elem()
	fmt.Println(dst.OverflowInt(257))
	// Output:
	// false true
	// 1
	// true
}

func Example_variadicCall() {
	add := func(base int, rest ...int) int {
		for _, n := range rest {
			base += n
		}
		return base
	}
	v := reflect.ValueOf(add)
	a := v.Call([]reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2), reflect.ValueOf(3)})
	b := v.CallSlice([]reflect.Value{reflect.ValueOf(1), reflect.ValueOf([]int{2, 3})})
	fmt.Println(a[0].Int(), b[0].Int())
	// Output:
	// 6 6
}

func Example_deepEqual() {
	fmt.Println(reflect.DeepEqual([]int(nil), []int{}))
	fmt.Println(reflect.DeepEqual(int(1), int64(1)))
	fn := func() {}
	fmt.Println(reflect.DeepEqual(fn, fn))
	// Output:
	// false
	// false
	// false
}
