package utils

import (
	"bytes"
	"github.com/a20070322/go_fast_admin/pkg/auto_curd/config_type"
	"log"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"unicode"
)

func IsHasUUidFn(fields []*config_type.FieldType) bool {
	bool := false
	for _, v := range fields {
		if v.Type == "UUID" {
			bool = true
			break
		}
	}
	return bool
}

//获取 Controller Module Name
func GetControllerModule(str string) string {
	arr := strings.Split(str, "/")
	return arr[len(arr)-1]
}

//ToUpper
func HumpToLowercase(str string) string {
	return strings.ToLower(str)
}

func IdIsUUIDFn(fields []*config_type.FieldType) (b bool) {
	b = false
	for _, v := range fields {
		if v.Name == "id" && v.Type == "UUID" {
			b = true
			break
		}
	}
	return b
}


func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}
// 内嵌bytes.Buffer，支持连写
type Buffer struct {
	*bytes.Buffer
}
func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*****内存不够了！******")
		}
	}()
	b.WriteString(s)
	return b
}
func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}
// 驼峰式写法转为下划线写法
func Camel2Case(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}



// 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// 首字母大写
func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// 首字母小写
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func CheckEqLine(line, key int) bool {
	return line-1 == key
}
