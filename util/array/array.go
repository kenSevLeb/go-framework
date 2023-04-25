package array

import (
	"log"
	"reflect"
	"strconv"
)

// UniqueInt int数组去重
func UniqueInt(items []int) []int {
	result := make([]int, 0, len(items))
	temp := map[int]struct{}{}
	for _, item := range items {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// UniqueString string数组去重
func UniqueString(items []string) []string {
	result := make([]string, 0, len(items))
	temp := map[string]struct{}{}
	for _, item := range items {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// 字符串转int
func String2Int(items []string) []int {
	result := make([]int, 0, len(items))
	for _, item := range items {
		i, err := strconv.Atoi(item)
		if err != nil {
			continue
		}
		result = append(result, i)
	}
	return result
}

// int转字符串
func Int2String(items []int) []string {
	result := make([]string, 0, len(items))
	for _, item := range items {
		result = append(result, strconv.Itoa(item))
	}
	return result
}

// 检查数组是否存在某元素
// haystack supported types: slice, array or map
func InArray(needle interface{} /* 元素 */, haystack interface{} /* 数组 */) bool {
	if haystack == nil {
		return false
	}
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(needle, val.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				return true
			}
		}
	default:
		log.Println("haystack: haystack type must be slice, array or map")
	}

	return false
}

// 删除int数组里的某个元素
func RemoveInt(items []int, index int) []int {
	if index < len(items)-1 {
		items = append(items[:index], items[index+1:]...)
	} else {
		items = items[:index]
	}
	return items
}

// 删除字符串数组里的某个元素
func RemoveString(items []string, index int) []string {
	if index < len(items)-1 {
		items = append(items[:index], items[index+1:]...)
	} else {
		items = items[:index]
	}
	return items
}
