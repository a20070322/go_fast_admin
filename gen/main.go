package main

import (
	"fmt"
	"github.com/a20070322/go_fast_admin/ent/schema"
	"strings"
)

//字符首字母大写转换
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

func PipeCapitalize(str string) string {
	arr := strings.Split(str, "_")
	for k, v := range arr {
		arr[k] = Capitalize(v)
	}
	return strings.Join(arr, "")
}
func main() {
	fields := schema.AdminUser{}.Fields()
	strR := ""
	strT := ""
	le := len(fields)
	for k, field := range fields {
		name := PipeCapitalize(field.Descriptor().Name)
		strR += fmt.Sprintf("Set%s(form.%s)", name, name)
		if k+1 < le {
			strR += ". \n"
		}
		strT += fmt.Sprintf("%s %s `json:\"%s\"", name, field.Descriptor().Info, field.Descriptor().Name)
		if field.Descriptor().Optional != true {
			strT += " binding:\"required\""
		}
		strT += fmt.Sprintf("` //%s\n", field.Descriptor().Comment)
	}
	fmt.Println(strR)
	fmt.Println("\n")
	fmt.Println(strT)

}
