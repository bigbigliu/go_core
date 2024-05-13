package pkgs

import "strings"

// IsNoRowFoundError gorm 'record not found' 错误处理
func IsNoRowFoundError(err error) bool {
	if err == nil {
		return false
	}

	if strings.Contains(err.Error(), "record not found") {
		return true
	}
	return false
}

// IsNoRowFoundError redis 'record not found' 错误处理
func IsRedisNilError(err error) bool {
	if err == nil {
		return false
	}

	if strings.Contains(err.Error(), "redis: nil") {
		return true
	}
	return false
}
