package practice

import (
	"reflect"
	"sync"
)

// FieldInfo 描述一个直接声明的导出字段，不包含运行时的字段值。
type FieldInfo struct {
	Name       string
	Index      int
	Type       reflect.Type
	JSONTag    string
	HasJSONTag bool
}

type Violation struct {
	Field string
	Rule  string
}

// FieldCache 使用后不能复制；缓存中只保存类型元数据。
type FieldCache struct {
	fields sync.Map
}
