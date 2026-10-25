//go:build !solution

package practice

import "reflect"

// E01：返回动态类型及 Kind；nil 返回 nil、reflect.Invalid。
func Inspect(x any) (reflect.Type, reflect.Kind) {
	panic("TODO E01: Inspect")
}

// E02：识别 nil 接口与常用六种可为 nil 的类型值；其余返回 false。
func IsNilLike(x any) bool {
	panic("TODO E02: IsNilLike")
}

// E03：接收 struct 或非 nil 的 *struct，返回直接声明的导出字段元数据。
func Fields(x any) ([]FieldInfo, error) {
	panic("TODO E03: Fields")
}

// E04：通过非 nil 的 *struct 修改直接声明的导出字段，只允许可赋值的类型。
func SetField(dst any, name string, src any) error {
	panic("TODO E04: SetField")
}

// E05：按自定义 map 标签生成浅层 map；omitempty 使用 Go 的 IsZero 语义。
func ToMap(x any) (map[string]any, error) {
	panic("TODO E05: ToMap")
}

// E06：按 default 标签填充零值；支持 string、bool、有符号整数，失败不写入。
func ApplyDefaults(dst any) error {
	panic("TODO E06: ApplyDefaults")
}

// E07：按 validate 标签执行 required / min=N，收集错误；非法规则返回 error。
func Validate(x any) ([]Violation, error) {
	panic("TODO E07: Validate")
}

// E08：调用固定参数个数的非 nil 函数。参数检查失败返回 error，业务 panic 继续传播。
func Invoke(fn any, args ...any) ([]any, error) {
	panic("TODO E08: Invoke")
}

// E09：给字符串键、有符号整数值的非 nil map 增加计数；检查加法和目标类型溢出。
func IncrementMap(m any, key string, delta int64) error {
	panic("TODO E09: IncrementMap")
}

// E10：并发安全地缓存 struct / *struct 的字段元数据；返回独立切片。
func (c *FieldCache) Metadata(t reflect.Type) ([]FieldInfo, error) {
	panic("TODO E10: FieldCache.Metadata")
}

// E11：按 db 标签把模拟行写入 *struct；严格赋值，任意失败都不写入。
func ScanRow(dst any, columns []string, values []any) error {
	panic("TODO E11: ScanRow")
}
