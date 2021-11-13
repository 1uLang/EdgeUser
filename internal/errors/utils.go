package errors

import "strings"

// 判断是否为资源无法查看错误
func IsResourceNotFound(err error) bool {
	if err == nil {
		return false
	}
	if strings.Index(err.Error(), "resource not found") >= 0 {
		return true
	}
	return false
}
