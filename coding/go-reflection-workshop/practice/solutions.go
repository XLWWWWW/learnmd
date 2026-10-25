//go:build solution

package practice

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Inspect(x any) (reflect.Type, reflect.Kind) {
	t := reflect.TypeOf(x)
	if t == nil {
		return nil, reflect.Invalid
	}
	return t, t.Kind()
}

func IsNilLike(x any) bool {
	v := reflect.ValueOf(x)
	if !v.IsValid() {
		return true
	}
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}

// structValue 只解一层指针；这是练习约定，并非通用反射库的唯一选择。
func structValue(x any) (reflect.Value, error) {
	v := reflect.ValueOf(x)
	if !v.IsValid() {
		return reflect.Value{}, fmt.Errorf("expected struct or *struct, got nil")
	}
	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return reflect.Value{}, fmt.Errorf("nil struct pointer")
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return reflect.Value{}, fmt.Errorf("expected struct, got %s", v.Kind())
	}
	return v, nil
}

func writableStruct(dst any) (reflect.Value, error) {
	v := reflect.ValueOf(dst)
	if !v.IsValid() || v.Kind() != reflect.Pointer || v.IsNil() {
		return reflect.Value{}, fmt.Errorf("destination must be a non-nil *struct")
	}
	v = v.Elem()
	if v.Kind() != reflect.Struct || !v.CanSet() {
		return reflect.Value{}, fmt.Errorf("destination must point to a settable struct")
	}
	return v, nil
}

func metadata(t reflect.Type) []FieldInfo {
	fields := make([]FieldInfo, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}
		tag, ok := f.Tag.Lookup("json")
		fields = append(fields, FieldInfo{
			Name: f.Name, Index: i, Type: f.Type, JSONTag: tag, HasJSONTag: ok,
		})
	}
	return fields
}

func Fields(x any) ([]FieldInfo, error) {
	v, err := structValue(x)
	if err != nil {
		return nil, err
	}
	return metadata(v.Type()), nil
}

// assigned 将 nil 转成目标类型的零值。非 nil 值只检查 AssignableTo，不做隐式转换。
func assigned(src any, target reflect.Type) (reflect.Value, error) {
	if src == nil {
		switch target.Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
			return reflect.Zero(target), nil
		default:
			return reflect.Value{}, fmt.Errorf("nil is not assignable to %s", target)
		}
	}
	v := reflect.ValueOf(src)
	if !v.Type().AssignableTo(target) {
		return reflect.Value{}, fmt.Errorf("%s is not assignable to %s", v.Type(), target)
	}
	return v, nil
}

func SetField(dst any, name string, src any) error {
	v, err := writableStruct(dst)
	if err != nil {
		return err
	}
	for i := 0; i < v.NumField(); i++ {
		f := v.Type().Field(i)
		if f.Name != name {
			continue
		}
		fv := v.Field(i)
		if !f.IsExported() || !fv.CanSet() {
			return fmt.Errorf("field %s is not writable", name)
		}
		sv, err := assigned(src, fv.Type())
		if err != nil {
			return fmt.Errorf("field %s: %w", name, err)
		}
		fv.Set(sv)
		return nil
	}
	return fmt.Errorf("direct field %q not found", name)
}

func ToMap(x any) (map[string]any, error) {
	v, err := structValue(x)
	if err != nil {
		return nil, err
	}
	out := make(map[string]any)
	seen := make(map[string]bool)
	for i := 0; i < v.NumField(); i++ {
		f := v.Type().Field(i)
		if !f.IsExported() {
			continue
		}
		parts := strings.Split(f.Tag.Get("map"), ",")
		omit := false
		for _, option := range parts[1:] {
			if option != "omitempty" {
				return nil, fmt.Errorf("field %s: unsupported map option %q", f.Name, option)
			}
			omit = true
		}
		if parts[0] == "-" {
			continue
		}
		name := parts[0]
		if name == "" {
			name = f.Name
		}
		// 先检查 schema 冲突，再决定是否省略当前值。
		if seen[name] {
			return nil, fmt.Errorf("duplicate map key %q", name)
		}
		seen[name] = true
		fv := v.Field(i)
		if omit && fv.IsZero() {
			continue
		}
		if !fv.CanInterface() {
			return nil, fmt.Errorf("field %s cannot be exposed", f.Name)
		}
		out[name] = fv.Interface()
	}
	return out, nil
}

type change struct {
	index int
	value reflect.Value
}

func ApplyDefaults(dst any) error {
	v, err := writableStruct(dst)
	if err != nil {
		return err
	}
	var pending []change
	for i := 0; i < v.NumField(); i++ {
		f := v.Type().Field(i)
		fv := v.Field(i)
		raw, ok := f.Tag.Lookup("default")
		if !ok || !f.IsExported() || !fv.IsZero() {
			continue
		}
		if !fv.CanSet() {
			return fmt.Errorf("field %s cannot be set", f.Name)
		}
		parsed := reflect.New(f.Type).Elem()
		switch fv.Kind() {
		case reflect.String:
			parsed.SetString(raw)
		case reflect.Bool:
			n, err := strconv.ParseBool(raw)
			if err != nil {
				return fmt.Errorf("field %s: %w", f.Name, err)
			}
			parsed.SetBool(n)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			n, err := strconv.ParseInt(raw, 10, f.Type.Bits())
			if err != nil {
				return fmt.Errorf("field %s: %w", f.Name, err)
			}
			parsed.SetInt(n)
		default:
			return fmt.Errorf("field %s: unsupported default type %s", f.Name, f.Type)
		}
		pending = append(pending, change{i, parsed})
	}
	for _, ch := range pending {
		v.Field(ch.index).Set(ch.value)
	}
	return nil
}

func Validate(x any) ([]Violation, error) {
	v, err := structValue(x)
	if err != nil {
		return nil, err
	}
	var violations []Violation
	for i := 0; i < v.NumField(); i++ {
		f := v.Type().Field(i)
		raw := f.Tag.Get("validate")
		if !f.IsExported() || raw == "" {
			continue
		}
		fv := v.Field(i)
		for _, rule := range strings.Split(raw, ",") {
			failed := false
			switch {
			case rule == "required":
				failed = fv.IsZero()
			case strings.HasPrefix(rule, "min="):
				minimum, err := strconv.ParseInt(strings.TrimPrefix(rule, "min="), 10, 64)
				if err != nil {
					return nil, fmt.Errorf("field %s: bad rule %q", f.Name, rule)
				}
				switch fv.Kind() {
				case reflect.String:
					if minimum < 0 {
						return nil, fmt.Errorf("field %s: string min must be non-negative", f.Name)
					}
					failed = int64(utf8.RuneCountInString(fv.String())) < minimum
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					failed = fv.Int() < minimum
				default:
					return nil, fmt.Errorf("field %s: min does not support %s", f.Name, f.Type)
				}
			default:
				return nil, fmt.Errorf("field %s: unknown rule %q", f.Name, rule)
			}
			if failed {
				violations = append(violations, Violation{f.Name, rule})
			}
		}
	}
	return violations, nil
}

func Invoke(fn any, args ...any) ([]any, error) {
	v := reflect.ValueOf(fn)
	if !v.IsValid() || v.Kind() != reflect.Func || v.IsNil() {
		return nil, fmt.Errorf("expected a non-nil function")
	}
	t := v.Type()
	if t.IsVariadic() {
		return nil, fmt.Errorf("variadic functions are outside this exercise")
	}
	if len(args) != t.NumIn() {
		return nil, fmt.Errorf("expected %d arguments, got %d", t.NumIn(), len(args))
	}
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		var err error
		in[i], err = assigned(arg, t.In(i))
		if err != nil {
			return nil, fmt.Errorf("argument %d: %w", i, err)
		}
	}
	// 业务函数的 panic 原样传播；不要用无差别 recover 混淆输入错误与业务缺陷。
	values := v.Call(in)
	out := make([]any, len(values))
	for i, value := range values {
		out[i] = value.Interface()
	}
	return out, nil
}

func IncrementMap(m any, key string, delta int64) error {
	v := reflect.ValueOf(m)
	if !v.IsValid() || v.Kind() != reflect.Map || v.IsNil() {
		return fmt.Errorf("expected a non-nil map")
	}
	t := v.Type()
	if t.Key().Kind() != reflect.String {
		return fmt.Errorf("map keys must have string kind")
	}
	switch t.Elem().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	default:
		return fmt.Errorf("map values must have signed integer kind")
	}
	k := reflect.ValueOf(key).Convert(t.Key())
	old := int64(0)
	if current := v.MapIndex(k); current.IsValid() {
		old = current.Int()
	}
	if (delta > 0 && old > math.MaxInt64-delta) || (delta < 0 && old < math.MinInt64-delta) {
		return fmt.Errorf("int64 addition overflow")
	}
	next := old + delta
	value := reflect.New(t.Elem()).Elem()
	if value.OverflowInt(next) {
		return fmt.Errorf("%d overflows %s", next, t.Elem())
	}
	value.SetInt(next)
	v.SetMapIndex(k, value)
	return nil
}

func (c *FieldCache) Metadata(t reflect.Type) ([]FieldInfo, error) {
	if t == nil {
		return nil, fmt.Errorf("nil type")
	}
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct or *struct type, got %s", t)
	}
	if cached, ok := c.fields.Load(t); ok {
		return append([]FieldInfo(nil), cached.([]FieldInfo)...), nil
	}
	fields, _ := c.fields.LoadOrStore(t, metadata(t))
	return append([]FieldInfo(nil), fields.([]FieldInfo)...), nil
}

func ScanRow(dst any, columns []string, values []any) error {
	v, err := writableStruct(dst)
	if err != nil {
		return err
	}
	if len(columns) != len(values) {
		return fmt.Errorf("column/value counts differ")
	}
	indices := make(map[string]int)
	for i := 0; i < v.NumField(); i++ {
		f := v.Type().Field(i)
		if !f.IsExported() {
			continue
		}
		name := f.Tag.Get("db")
		if name == "-" {
			continue
		}
		if name == "" {
			name = f.Name
		}
		if _, exists := indices[name]; exists {
			return fmt.Errorf("duplicate db field %q", name)
		}
		indices[name] = i
	}
	seen := make(map[string]bool)
	var pending []change
	for i, column := range columns {
		index, ok := indices[column]
		if !ok {
			return fmt.Errorf("unknown column %q", column)
		}
		if seen[column] {
			return fmt.Errorf("duplicate column %q", column)
		}
		seen[column] = true
		fv := v.Field(index)
		if !fv.CanSet() {
			return fmt.Errorf("field %s cannot be set", v.Type().Field(index).Name)
		}
		sv, err := assigned(values[i], fv.Type())
		if err != nil {
			return fmt.Errorf("column %s: %w", column, err)
		}
		pending = append(pending, change{index, sv})
	}
	for _, ch := range pending {
		v.Field(ch.index).Set(ch.value)
	}
	return nil
}
