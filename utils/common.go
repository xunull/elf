package utils

import "syscall"

// 获取可以打开的最大文件数
func GetMaxNumFiles() int {
	var rLimit syscall.Rlimit

	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		return 0
	}
	return int(rLimit.Max)
}
