package practice

import (
	"reflect"
	"testing"
)

var benchmarkIntSink int64
var benchmarkFieldsSink []FieldInfo

// E12：先预测，再实测。这里测的是同一字段的读取，不是完整业务请求。
func BenchmarkE12FieldAccess(b *testing.B) {
	u := struct{ Age int64 }{Age: 42}
	b.Run("Direct", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			benchmarkIntSink = u.Age
		}
	})
	b.Run("ReflectByName", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			benchmarkIntSink = reflect.ValueOf(u).FieldByName("Age").Int()
		}
	})
	b.Run("ReflectByIndex", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			benchmarkIntSink = reflect.ValueOf(u).Field(0).Int()
		}
	})
}

func BenchmarkE12Metadata(b *testing.B) {
	type User struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Enabled bool   `json:"enabled"`
	}
	u := User{}
	b.Run("Uncached", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			fields, err := Fields(u)
			if err != nil {
				b.Fatal(err)
			}
			benchmarkFieldsSink = fields
		}
	})
	b.Run("CachedWithDefensiveCopy", func(b *testing.B) {
		var c FieldCache
		typ := reflect.TypeOf(u)
		if _, err := c.Metadata(typ); err != nil {
			b.Fatal(err)
		}
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			fields, err := c.Metadata(typ)
			if err != nil {
				b.Fatal(err)
			}
			benchmarkFieldsSink = fields
		}
	})
}
