// conv 各种数据转化的方法
package conv

import (
	"encoding/json"
	"fmt"
	"github.com/kenSevLeb/go-framework/util/strings"
	"log"
	"reflect"
	"strconv"
	"unsafe"
)

// int转string
func Int2String(i int) string {
	return strconv.Itoa(i)
}

// string转int
func String2Int(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// string转byte
func String2Byte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// byte转string
func Byte2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// struct转map,默认使用属性名称为字段名称，可以用json标签重命名
func Struct2Map(obj interface{}) map[string]interface{} {
	var node map[string]interface{}
	objT := reflect.TypeOf(obj)
	if objT.Kind() != reflect.Struct {
		log.Println("argument is not of the expected type")
		return node
	}
	objV := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < objT.NumField(); i++ {
		filedName := strings.Default(objT.Field(i).Tag.Get("json"), objT.Field(i).Name)
		switch objV.Field(i).Type().Kind() {
		case reflect.Struct:
			node = Struct2Map(objV.Field(i).Interface())
			data[filedName] = node
		case reflect.Slice:
			target := objV.Field(i).Interface()
			tmp := make([]map[string]interface{}, reflect.ValueOf(target).Len())
			for j := 0; j < reflect.ValueOf(target).Len(); j++ {
				if reflect.ValueOf(target).Index(j).Kind() == reflect.Struct {
					node = Struct2Map(reflect.ValueOf(target).Index(j).Interface())
					tmp[j] = node
				}
			}
			data[filedName] = tmp
		default:
			data[filedName] = objV.Field(i).Interface()
		}
	}
	return data
}

// map转struct,用json转换，要求obj为指针
func Map2Struct(mapInstance map[string]interface{}, obj interface{}) error {
	buf, err := json.Marshal(mapInstance)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(buf, obj); err != nil {
		return err
	}

	return nil
}

// 将整形或字符串的interface转回成字符串
func InterToStr(data interface{}) string {
	var ret string
	if data == nil {
		return ret
	}
	dataType := fmt.Sprintf("%T", data)
	if dataType == "string" {
		ret = data.(string)
	} else if dataType == "float64" {
		//float转字符串
		ret = fmt.Sprintf("%.f", data.(float64))
	} else if dataType == "int64" {
		ret = Int2String(int(data.(int64)))
	} else { //int
		ret = Int2String(data.(int))
	}
	return ret
}

// 将整形或字符串的interface转回成int
func InterToInt(data interface{}) int {
	var ret int
	if data == nil {
		return ret
	}
	dataType := fmt.Sprintf("%T", data)
	if dataType == "string" {
		ret = String2Int(data.(string))
	} else if dataType == "float64" {
		ret, _ = strconv.Atoi(fmt.Sprintf("%1.0f", data.(float64)))
	} else if dataType == "int64" {
		ret = int(data.(int64))
	} else {
		ret = data.(int)
	}
	return ret

}
